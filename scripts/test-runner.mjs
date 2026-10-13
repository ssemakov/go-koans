import assert from "node:assert/strict";
import { test } from "node:test";
import { readPlaygroundResult } from "../lib/runner.ts";

const output = (Message) => ({ Errors: "", Events: [{ Message, Kind: "stdout" }], VetErrors: "" });
test("passing tests require completed tests and the runner PASS result", () => {
  assert.equal(readPlaygroundResult(output("--- PASS: TestKoan (0.00s)\nPASS\n")).status, "passed");
  assert.equal(readPlaygroundResult(output("PASS\n")).status, "error");
  assert.equal(readPlaygroundResult(output("--- PASS: TestKoan (0.00s)\npanic: later failure\n")).status, "failed");
});
test("failed, skipped, and zero-test runs cannot complete an exercise", () => {
  assert.equal(readPlaygroundResult(output("--- FAIL: TestKoan (0.00s)\nFAIL\n")).status, "failed");
  assert.equal(readPlaygroundResult(output("--- SKIP: TestKoan (0.00s)\nPASS\n")).status, "failed");
  assert.equal(readPlaygroundResult(output("--- PASS: TestKoan (0.00s)\n    --- SKIP: TestKoan/part (0.00s)\nPASS\n")).status, "failed");
  assert.match(readPlaygroundResult({ Errors: "", Events: null }).output, /No tests completed/);
  assert.equal(readPlaygroundResult({ Errors: "", Events: [] }).status, "error");
});
test("compiler errors and backend failures have distinct outcomes", () => {
  assert.equal(readPlaygroundResult({ Errors: "prog_test.go: syntax error", Events: null }).status, "failed");
  assert.equal(readPlaygroundResult({ Errors: "timeout running go build\n" }).status, "error");
  assert.throws(() => readPlaygroundResult({}), /Invalid/);
  assert.throws(() => readPlaygroundResult({ Errors: "", Events: {} }), /Invalid/);
});
