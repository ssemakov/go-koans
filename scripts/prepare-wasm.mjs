import { cp, mkdir, rm } from "node:fs/promises";
import { dirname, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import { build } from "esbuild";
import { wasmRoot } from "./wasm-assets.mjs";

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
// Preserve upstream attribution alongside the redistributed binaries.
for (const [name, directory] of [["browser", resolve(browser, "..")], ["go", go]]) {
  await mkdir(`${output}/licenses/${name}`, { recursive: true });
  for (const file of ["LICENSE", "THIRD_PARTY_NOTICES.md", "licenses"]) {
    await cp(`${directory}/${file}`, `${output}/licenses/${name}/${file}`, { recursive: true });
  }
}
console.log(`Prepared browser Go compiler at ${wasmRoot}`);
