import { createHash } from "node:crypto";

export const chunkSize = 4 * 1024 * 1024;
export const digest = (bytes) => createHash("sha256").update(bytes).digest("hex");

export function splitAsset(bytes, size = chunkSize) {
  if (!Number.isSafeInteger(size) || size < 1) throw new Error("Invalid chunk size");
  const chunks = [];
  for (let offset = 0; offset < bytes.length; offset += size) {
    const data = bytes.subarray(offset, offset + size);
    chunks.push({ data, sha256: digest(data), bytes: data.length });
  }
  return { bytes: bytes.length, sha256: digest(bytes), chunks };
}
