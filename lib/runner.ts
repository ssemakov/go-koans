export type RunResult = { status: "passed" | "failed" | "error"; output: string; passed: number; failed: number; skipped: number };

export function readWasmResult(result: { code: number; stdout: string; stderr: string; termination: string; trapMessage?: string }): RunResult {
  const output = [result.stdout, result.stderr].filter(Boolean).join("\n");
  const parsed = readPlaygroundResult({ Errors: "", Events: [{ Message: output }] });
  if (result.termination !== "exited") {
    return { ...parsed, status: "error", output: `${output}\nBrowser runner stopped: ${result.termination}.${result.trapMessage ? ` ${result.trapMessage}` : ""}` };
  }
  return { ...parsed, status: result.code !== 0 ? "failed" : parsed.status };
}

export function readPlaygroundResult(value: unknown): RunResult {
  if (!value || typeof value !== "object") throw new Error("Invalid Go Playground response.");
  const data = value as { Errors?: unknown; VetErrors?: unknown; Events?: unknown };
  if (typeof data.Errors !== "string") throw new Error("Invalid Go Playground response.");
  if (/timeout running go build/i.test(data.Errors)) return { status: "error", output: "Go Playground timed out building this test. Try again, or download the test and run it locally.", passed: 0, failed: 0, skipped: 0 };
  if (data.Errors) return { status: "failed", output: data.Errors, passed: 0, failed: 0, skipped: 0 };
  if (data.Events != null && !Array.isArray(data.Events)) throw new Error("Invalid Go Playground response.");
  const events: unknown[] = Array.isArray(data.Events) ? data.Events : [];
  const output = events.map((event: unknown) => {
    if (!event || typeof event !== "object" || typeof (event as { Message?: unknown }).Message !== "string") throw new Error("Invalid Go Playground output.");
    return (event as { Message: string }).Message;
  }).join("");
  const passed = (output.match(/^\s*--- PASS:/gm) || []).length;
  const failed = (output.match(/^\s*--- FAIL:/gm) || []).length;
  const skipped = (output.match(/^\s*--- SKIP:/gm) || []).length;
  const vet = typeof data.VetErrors === "string" ? data.VetErrors : "";
  if (data.Errors || vet) return { status: "failed", output: [data.Errors, vet, output].filter(Boolean).join("\n"), passed, failed, skipped };
  if (!passed && !failed && !skipped) return { status: "error", output: `${output}\nNo tests completed. Keep the TestKoan function and run it again.`, passed, failed, skipped };
  return { status: passed > 0 && failed === 0 && skipped === 0 && /^PASS\s*$/m.test(output) ? "passed" : "failed", output, passed, failed, skipped };
}
