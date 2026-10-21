import type { NextConfig } from "next";
import { wasmRoot } from "./scripts/wasm-assets.mjs";

const config: NextConfig = {
  env: { NEXT_PUBLIC_WASM_ROOT: wasmRoot },
  async headers() {
    return [{
      source: "/:path*",
      headers: [
        { key: "Cross-Origin-Opener-Policy", value: "same-origin" },
        { key: "Cross-Origin-Embedder-Policy", value: "require-corp" },
        { key: "Cross-Origin-Resource-Policy", value: "same-origin" },
      ],
    }, {
      source: "/wasm/:path*",
      headers: [
        { key: "Cache-Control", value: "public, max-age=31536000, immutable" },
        { key: "Vercel-CDN-Cache-Control", value: "public, max-age=31536000, immutable" },
      ],
    }, {
      source: "/wasm-cache-sw.js",
      headers: [{ key: "Cache-Control", value: "no-store" }],
    }];
  },
};
export default config;
