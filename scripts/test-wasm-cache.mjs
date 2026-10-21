import assert from "node:assert/strict";
import { webcrypto } from "node:crypto";
import { readFileSync } from "node:fs";
import { test } from "node:test";
import { runInNewContext } from "node:vm";
import { gzipSync } from "node:zlib";
import { splitAsset } from "./wasm-chunks.mjs";

const fixture = Buffer.from("a long compiler fixture with many distinct chunk payloads for retry and parallel-download checks");
const split = splitAsset(fixture, 8);
const asset = { ...split, contentType: "application/wasm", chunks: split.chunks.map(({ data: _, ...chunk }) => chunk) };
const origin = "https://koans.test";
const logical = `${origin}/wasm/1234567890abcdef/toolchains/compiler.wasm?sha256=${asset.sha256}`;
const source = readFileSync(new URL("../lib/wasm/cache-worker.js", import.meta.url), "utf8");

function worker({ store = new Map(), corrupt = false, encoded = false } = {}) {
  const transfer = encoded ? splitAsset(gzipSync(fixture), 8) : split;
  const metadata = encoded ? { ...asset, encoding: "gzip", chunks: transfer.chunks.map(({ data: _, ...chunk }) => chunk) } : asset;
  const handlers = {};
  const calls = [];
  let active = 0;
  let peak = 0;
  const cache = {
    async match(url) { return store.get(url)?.clone(); },
    async put(url, response) { store.set(url, response.clone()); },
    async delete(url) { store.delete(url); },
  };
  runInNewContext(source.replace("const assets = __WASM_ASSETS__;", `const assets = ${JSON.stringify({ "compiler.wasm": metadata })};`), {
    self: { location: { origin }, addEventListener(type, fn) { handlers[type] = fn; } },
    caches: { async open() { return cache; } },
    crypto: webcrypto, Response, URL, ReadableStream, DecompressionStream, Uint8Array, AbortSignal,
    async fetch(url) {
      calls.push(url); active++; peak = Math.max(peak, active);
      try {
        await new Promise((resolve) => setTimeout(resolve, 2));
        const chunk = transfer.chunks.find((c) => url.endsWith(`${c.sha256}.bin`));
        assert.ok(chunk, url);
        return new Response(corrupt ? "corrupted download" : chunk.data);
      } finally { active--; }
    },
  });
  function request(url = logical, method = "GET", mode = "cors") {
    let response;
    const waits = [];
    handlers.fetch({ request: { url, method, mode }, respondWith(value) { response = value; }, waitUntil(value) { waits.push(value); } });
    return { response, settled: () => Promise.all(waits) };
  }
  return { request, calls, store, peak: () => peak };
}

test("binary chunks reconstruct source bytes and stay below the configured size", () => {
  assert.deepEqual(Buffer.concat(split.chunks.map((c) => c.data)), fixture);
  assert.ok(split.chunks.every((c) => c.bytes <= 8));
  assert.throws(() => splitAsset(fixture, 0), /Invalid/);
});

test("compressed runtime transfers deliver the exact uncompressed WebAssembly bytes", async () => {
  const current = worker({ encoded: true });
  const result = current.request().response;
  assert.equal(result.headers.get("content-length"), String(fixture.length));
  assert.deepEqual(Buffer.from(await result.arrayBuffer()), fixture);
});

test("deduplicate concurrent SDK reads, bound parallel transfers, and persist across reloads/deployments", async () => {
  const current = worker();
  const first = current.request();
  const duplicate = current.request();
  for (const response of [first.response, duplicate.response]) {
    assert.equal(response.headers.get("content-type"), "application/wasm");
    assert.deepEqual(Buffer.from(await response.arrayBuffer()), fixture);
  }
  await first.settled(); await duplicate.settled();
  assert.equal(current.calls.length, split.chunks.length);
  assert.ok(current.peak() <= 4);
  const reloaded = worker({ store: current.store });
  const next = reloaded.request(logical.replace("1234567890abcdef", "fedcba0987654321"));
  assert.deepEqual(Buffer.from(await next.response.arrayBuffer()), fixture);
  assert.equal(reloaded.calls.length, 0);
});

test("cancelled downloads keep completed chunks and a retry fetches only missing parts", async () => {
  const first = worker();
  const request = first.request();
  const reader = request.response.body.getReader();
  await reader.read();
  await reader.cancel();
  await request.settled();
  const completed = new Set(first.store.keys());
  assert.ok(completed.size > 0 && completed.size < split.chunks.length);
  const retry = worker({ store: first.store });
  assert.deepEqual(Buffer.from(await retry.request().response.arrayBuffer()), fixture);
  assert.ok(retry.calls.every((url) => !completed.has(url)));
  assert.equal(retry.calls.length + completed.size, split.chunks.length);
});

test("corrupt cached data is refetched, corrupt downloads and mismatched pins are rejected", async () => {
  const current = worker();
  await current.request().response.arrayBuffer();
  const damaged = current.calls[0];
  current.store.set(damaged, new Response("bad cached bytes"));
  const next = worker({ store: current.store });
  assert.deepEqual(Buffer.from(await next.request().response.arrayBuffer()), fixture);
  assert.deepEqual(next.calls, [damaged]);
  const corrupt = worker({ corrupt: true });
  await assert.rejects(corrupt.request().response.arrayBuffer(), /integrity/);
  assert.equal(corrupt.store.size, 0);
  assert.equal(current.request(logical.replace(asset.sha256, "bad-pin")).response.status, 409);
});

test("the download worker never intercepts pages, app code, APIs, or unknown binaries", () => {
  const current = worker();
  for (const path of ["/", "/api/run", "/_next/static/app.js", "/wasm/1234567890abcdef/runtime/client.js", "/wasm/1234567890abcdef/toolchains/unknown.wasm"]) {
    assert.equal(current.request(`${origin}${path}`).response, undefined);
  }
  assert.equal(current.request(logical, "POST").response, undefined);
  assert.equal(current.request(logical, "GET", "navigate").response, undefined);
  assert.equal(current.request(logical.replace(origin, "https://other.test")).response, undefined);
});
