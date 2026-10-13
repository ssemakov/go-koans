import { readPlaygroundResult } from "@/lib/runner";

export const runtime = "nodejs";
export const maxDuration = 30;

export async function POST(request: Request) {
  const error = (output: string, status: number) => Response.json({ status: "error", output, passed: 0, failed: 0, skipped: 0 }, { status });
  if (Number(request.headers.get("content-length")) > 65_536) return error("Code is too large. The limit is 40,000 characters.", 413);
  let code: unknown;
  try {
    const text = await request.text();
    if (text.length > 65_536) return error("Request is too large.", 413);
    code = JSON.parse(text).code;
  } catch { return error("Send a JSON object with a code string.", 400); }
  if (typeof code !== "string" || !code.trim()) return error("Code is required.", 400);
  if (code.length > 40_000) return error("Code is too large. The limit is 40,000 characters.", 413);
  try {
    const response = await fetch("https://go.dev/_/compile", {
      method: "POST", headers: { "Content-Type": "application/x-www-form-urlencoded" },
      body: new URLSearchParams({ version: "2", body: code, withVet: "true" }),
      signal: AbortSignal.timeout(20_000), cache: "no-store",
    });
    if (!response.ok) return error(`Go Playground returned HTTP ${response.status}. Try again, or download the test and run it locally.`, 503);
    return Response.json(readPlaygroundResult(await response.json()));
  } catch {
    return error("Could not run the tests on the Go Playground. Try again, or download the test and run it locally.", 503);
  }
}
