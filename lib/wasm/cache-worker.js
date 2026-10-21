/* This worker handles only the pinned compiler binaries listed at build time.
 * Navigations, application scripts, API requests, and curriculum are untouched. */
// __WASM_ASSETS__ is injected by prepare-wasm.mjs.
const assets = __WASM_ASSETS__;
const cacheName = "go-koans:wasm-chunks:v1";
const pending = new Map();
let activeDownloads = 0;
const waiting = [];

self.addEventListener("install", (event) => event.waitUntil(self.skipWaiting()));
self.addEventListener("activate", (event) => event.waitUntil(self.clients.claim()));

self.addEventListener("fetch", (event) => {
  const request = event.request;
  const url = new URL(request.url);
  if (request.method !== "GET" || request.mode === "navigate" || url.origin !== self.location.origin) return;
  const match = /^\/wasm\/[a-f0-9]{16}\/(?:toolchains|runtime\/assets)\/([^/]+)$/.exec(url.pathname);
  const asset = match && assets[match[1]];
  if (!asset) return;
  const pin = url.searchParams.get("sha256");
  if (pin && pin !== asset.sha256) {
    event.respondWith(new Response("Compiler asset version changed. Reload the page.", { status: 409 }));
    return;
  }
  event.respondWith(assemble(asset, event));
});

async function sha256(bytes) {
  const hash = new Uint8Array(await crypto.subtle.digest("SHA-256", bytes));
  return Array.from(hash, (byte) => byte.toString(16).padStart(2, "0")).join("");
}

async function valid(bytes, chunk) {
  return bytes.byteLength === chunk.bytes && await sha256(bytes) === chunk.sha256;
}

async function download(chunk) {
  const url = new URL(`/wasm/chunks/${chunk.sha256}.bin`, self.location.origin).href;
  let cache;
  try {
    cache = await caches.open(cacheName);
    const stored = await cache.match(url);
    if (stored) {
      const bytes = new Uint8Array(await stored.arrayBuffer());
      if (await valid(bytes, chunk)) return bytes;
      await cache.delete(url);
    }
  } catch { /* Private browsing or a full disk can disable persistent caching. */ }

  // Reserve slots across all assets, including concurrent compiler workers.
  if (activeDownloads < 4) activeDownloads++;
  else await new Promise((resolve) => waiting.push(resolve));
  try {
    const response = await fetch(url, { signal: AbortSignal.timeout(90_000) });
    if (!response.ok) throw new Error(`Compiler chunk download failed (${response.status}).`);
    const bytes = new Uint8Array(await response.arrayBuffer());
    if (!await valid(bytes, chunk)) throw new Error("Compiler chunk failed integrity verification.");
    try { await cache?.put(url, new Response(bytes)); }
    catch { /* HTTP caching and the current run remain available without disk space. */ }
    return bytes;
  } finally {
    const next = waiting.shift();
    if (next) next();
    else activeDownloads--;
  }
}

function getChunk(chunk) {
  let result = pending.get(chunk.sha256);
  if (!result) {
    result = download(chunk).finally(() => pending.delete(chunk.sha256));
    pending.set(chunk.sha256, result);
  }
  return result;
}

function assemble(asset, event) {
  let index = 0;
  let cancelled = false;
  let finish;
  const completed = new Promise((resolve) => { finish = resolve; });
  event.waitUntil(completed);
  const reads = [];
  function queue(index) {
    if (index >= asset.chunks.length) return;
    // Observe rejections immediately, even if a later chunk fails first.
    reads[index] = getChunk(asset.chunks[index]).then((bytes) => ({ bytes }), (error) => ({ error }));
  }
  const body = new ReadableStream({
    start() { for (let i = 0; i < 4; i++) queue(i); },
    async pull(controller) {
      try {
        if (index === asset.chunks.length) { controller.close(); finish(); return; }
        const result = await reads[index];
        if (cancelled) return;
        if (result.error) throw result.error;
        controller.enqueue(result.bytes);
        index++;
        queue(index + 3);
      } catch (error) { controller.error(error); finish(); }
    },
    cancel() {
      cancelled = true;
      // Keep completed chunks on disk and finish only the bounded in-flight set.
      Promise.allSettled(reads).then(finish);
    },
  });
  return new Response(asset.encoding === "gzip" ? body.pipeThrough(new DecompressionStream("gzip")) : body, { headers: {
    "Content-Type": asset.contentType,
    "Content-Length": String(asset.bytes),
    "Cache-Control": "public, max-age=31536000, immutable",
    "Cross-Origin-Resource-Policy": "same-origin",
  } });
}
