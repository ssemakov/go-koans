import { createHash } from "node:crypto";
import { readFileSync } from "node:fs";

// Content-address the complete deployment so an older tab cannot mix a new
// adapter with cached compiler assets. No service worker intercepts app loads.
const hash = createHash("sha256");
for (const path of ["package-lock.json", "lib/runner.ts", "lib/wasm/runtime.ts", "lib/wasm/harness.ts", "scripts/prepare-wasm.mjs"]) {
  hash.update(readFileSync(new URL(`../${path}`, import.meta.url)));
}
export const wasmRoot = `/wasm/${hash.digest("hex").slice(0, 16)}`;
