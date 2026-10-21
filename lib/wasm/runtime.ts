import { createBrowserEngine, prefetchBrowserToolchain } from "@wasm-oj/browser";
import { browserSource } from "@wasm-oj/toolchain-go";
import { readWasmResult, type RunResult } from "../runner";
import { testFiles } from "./harness";

const toolchains = [browserSource(new URL("../toolchains/", import.meta.url).href)];
let engine: Awaited<ReturnType<typeof createBrowserEngine>> | undefined;
let prefetched = false;

export async function runWasmTests(code: string, signal: AbortSignal, progress: (label: string) => void): Promise<RunResult> {
  const files = testFiles(code);
  signal.throwIfAborted();
  // Download outside the compiler's deadline. Completed, integrity-checked
  // downloads are reused by the SDK; aborted downloads can be retried.
  if (!prefetched) {
    progress("Preparing Go 1.26.5… First use downloads about 50 MB. Completed parts are saved for retries.");
    await prefetchBrowserToolchain(toolchains, {
      language: "go", target: "wasip1", optimization: "release", signal,
      onProgress: ({ loadedBytes, totalBytes }) => progress(`Downloading Go compiler… ${(loadedBytes / 1_000_000).toFixed(1)} / ${(totalBytes / 1_000_000).toFixed(1)} MB (${Math.round(loadedBytes / totalBytes * 100)}%)`),
    });
    prefetched = true;
  }
  signal.throwIfAborted();

  let active = engine;
  const cancel = () => {
    active?.dispose();
    if (engine === active) engine = undefined;
  };
  signal.addEventListener("abort", cancel, { once: true });
  let unsubscribe: (() => void) | undefined;
  try {
    if (!active) {
      progress("Starting browser compiler…");
      active = await createBrowserEngine({ toolchains, artifactCache: true });
      if (signal.aborted) { active.dispose(); signal.throwIfAborted(); }
      engine = active;
    }
    unsubscribe = active.onProgress(({ phase, progress: fraction }) => {
      const label = phase === "running" ? "Running Go tests…" : phase === "loading-toolchain" ? "Preparing compiler and runtime…" : "Compiling Go tests…";
      progress(`${label}${phase === "loading-toolchain" && fraction !== undefined ? ` ${Math.round(fraction * 100)}%` : ""}`);
    });
    const build = await active.compile({ language: "go", target: "wasip1", optimization: "release", entry: "main.go", files });
    signal.throwIfAborted();
    if (!build.success || !build.artifact) {
      const output = [build.stderr, build.stdout, ...build.diagnostics.map((d) => `${d.file}:${d.line}:${d.column}: ${d.message}`)].filter(Boolean).join("\n");
      return { status: "failed", output: output || "Go compilation failed.", passed: 0, failed: 0, skipped: 0 };
    }
    const result = await active.run(build.artifact, {
      args: ["-test.v", "-test.timeout=5s"],
      env: { TMPDIR: "/tmp" },
      // A writable virtual /tmp supports testing.T.TempDir and os file exercises.
      files: { "/tmp/.keep": "" },
      resources: {
        instructionBudget: 1_000_000_000, logicalTimeLimitMs: 10_000,
        memoryLimitBytes: 256 * 1024 * 1024, outputLimitBytes: 256 * 1024,
        filesystemWriteLimitBytes: 8 * 1024 * 1024, filesystemEntryLimit: 1024,
        wallTimeLimitMs: 10_000,
      },
    });
    signal.throwIfAborted();
    return readWasmResult(result);
  } catch (error) {
    cancel(); // Discard unhealthy workers; a retry reuses downloaded toolchains.
    throw error;
  } finally {
    unsubscribe?.();
    signal.removeEventListener("abort", cancel);
  }
}
