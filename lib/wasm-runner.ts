import type { RunResult } from "./runner";
import type * as Runtime from "./wasm/runtime";
import { prepareWasmCache } from "./wasm-cache";

export async function runWasmTests(code: string, signal: AbortSignal, progress: (label: string) => void): Promise<RunResult> {
  if (!window.isSecureContext || !window.crossOriginIsolated || typeof SharedArrayBuffer === "undefined") {
    throw new Error("Browser WebAssembly needs cross-origin isolation and SharedArrayBuffer over HTTPS (or localhost). Select Go Playground or download the test to run it locally.");
  }
  progress("Preparing compiler download cache…");
  await prepareWasmCache(signal);
  progress("Loading browser runner…");
  const url = `${process.env.NEXT_PUBLIC_WASM_ROOT}/runtime/client.js`;
  const runtime: typeof Runtime = await import(/* webpackIgnore: true */ /* turbopackIgnore: true */ url);
  signal.throwIfAborted();
  return runtime.runWasmTests(code, signal, progress);
}
