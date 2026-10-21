import { cp, mkdir, readFile, readdir, rm, writeFile } from "node:fs/promises";
import { dirname, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import { gzipSync } from "node:zlib";
import { build } from "esbuild";
import { wasmRoot } from "./wasm-assets.mjs";
import { digest, splitAsset } from "./wasm-chunks.mjs";

const root = fileURLToPath(new URL("../", import.meta.url));
const output = resolve(root, `public${wasmRoot}`);
await rm(resolve(root, "public/wasm"), { recursive: true, force: true });
await mkdir(`${output}/runtime`, { recursive: true });
await build({
  entryPoints: [resolve(root, "lib/wasm/runtime.ts")],
  bundle: true, format: "esm", platform: "browser", target: "es2022",
  minify: true, outfile: `${output}/runtime/client.js`,
});
const browser = dirname(fileURLToPath(import.meta.resolve("@wasm-oj/browser")));
const go = resolve(dirname(fileURLToPath(import.meta.resolve("@wasm-oj/toolchain-go"))), "..");
await cp(`${browser}/assets`, `${output}/runtime/assets`, { recursive: true });
await cp(`${go}/assets`, `${output}/toolchains`, { recursive: true });
// Keep each transfer small and cacheable. Chunks retain their URL across app
// deployments, so a cancelled download or an adapter update reuses existing data.
const chunkDirectory = resolve(root, "public/wasm/chunks");
await mkdir(chunkDirectory, { recursive: true });
const assets = {};
for (const directory of [`${output}/toolchains`, `${output}/runtime/assets`]) {
  for (const name of await readdir(directory)) {
    if (!/\.(?:wasm|bin|json)$/.test(name)) continue;
    const path = `${directory}/${name}`;
    const original = await readFile(path);
    const compressed = name.endsWith(".wasm");
    const asset = splitAsset(compressed ? gzipSync(original, { level: 9 }) : original);
    const chunks = [];
    for (const { data, ...chunk } of asset.chunks) {
      await writeFile(`${chunkDirectory}/${chunk.sha256}.bin`, data);
      chunks.push(chunk);
    }
    assets[name] = { bytes: original.length, sha256: digest(original), encoding: compressed ? "gzip" : "identity", contentType: name.endsWith(".wasm") ? "application/wasm" : name.endsWith(".json") ? "application/json" : "application/octet-stream", chunks };
    await rm(path);
  }
}
await build({
  entryPoints: [resolve(root, "lib/wasm/cache-worker.js")], bundle: true,
  format: "iife", platform: "browser", target: "es2022", minify: true,
  define: { __WASM_ASSETS__: JSON.stringify(assets) },
  outfile: resolve(root, "public/wasm-cache-sw.js"),
});
// Preserve upstream attribution alongside the redistributed binaries.
for (const [name, directory] of [["browser", resolve(browser, "..")], ["go", go]]) {
  await mkdir(`${output}/licenses/${name}`, { recursive: true });
  for (const file of ["LICENSE", "THIRD_PARTY_NOTICES.md", "licenses"]) {
    await cp(`${directory}/${file}`, `${output}/licenses/${name}/${file}`, { recursive: true });
  }
}
console.log(`Prepared browser Go compiler at ${wasmRoot} with 4 MiB cached downloads`);
