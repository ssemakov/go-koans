/** Activate binary delivery before loading the SDK or any of its workers. */
export async function prepareWasmCache(signal: AbortSignal): Promise<void> {
  if (!("serviceWorker" in navigator)) throw new Error("Browser compilation needs service workers for compiler downloads. Select Go Playground or download the test.");
  signal.throwIfAborted();
  const registration = await navigator.serviceWorker.register("/wasm-cache-sw.js", { scope: "/", updateViaCache: "none" });
  signal.throwIfAborted();
  const worker = registration.installing ?? registration.waiting ?? registration.active;
  if (!worker) throw new Error("The compiler download cache could not start. Reload and try again.");
  await new Promise<void>((resolve, reject) => {
    const timer = setTimeout(() => settle(new Error("The compiler download cache did not start. Reload and try again.")), 15_000);
    function settle(error?: unknown) {
      clearTimeout(timer);
      worker?.removeEventListener("statechange", check);
      navigator.serviceWorker.removeEventListener("controllerchange", check);
      signal.removeEventListener("abort", abort);
      if (error) reject(error); else resolve();
    }
    function check() {
      if (worker?.state === "redundant") settle(new Error("The compiler download cache could not activate. Reload and try again."));
      else if (worker?.state === "activated" && navigator.serviceWorker.controller) settle();
    }
    function abort() { settle(signal.reason); }
    worker.addEventListener("statechange", check);
    navigator.serviceWorker.addEventListener("controllerchange", check);
    signal.addEventListener("abort", abort, { once: true });
    if (signal.aborted) abort(); else check();
  });
}
