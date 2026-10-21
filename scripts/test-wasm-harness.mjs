import assert from "node:assert/strict";
import { readFileSync } from "node:fs";
import { test } from "node:test";
import { testFiles, testNames } from "../lib/wasm/harness.ts";

test("discover top-level Go tests, excluding comments, strings, methods, and lowercase helpers", () => {
  const code = `package main
// func TestComment(t *testing.T) {}
/* func TestBlockComment(t *testing.T) {} */
var text = "func TestString(t *testing.T) {}"
func TestKoan(t *testing.T) {}
func TestAdded(t *testing.T) {}
func TestΩ(t *testing.T) {}
func Testé(t *testing.T) {}
func Testhelper(t *testing.T) {}
func (s *State) TestMethod(t *testing.T) {}
`;
  assert.deepEqual(testNames(code), ["TestKoan", "TestAdded", "TestΩ"]);
  const files = testFiles(code);
  assert.equal(files["koan.go"], code);
  assert.match(files["main.go"], /testing.Main/);
  assert.match(files["main.go"], /F: TestAdded/);
  assert.doesNotMatch(files["main.go"], /TestComment|TestString|TestMethod/);
});

test("missing tests, oversized code, and unsupported TestMain cannot silently complete an exercise", () => {
  assert.throws(() => testFiles("package main\n// func TestKoan() {}"), /No tests found/);
  assert.throws(() => testFiles(" ".repeat(40_001)), /40,000/);
  assert.throws(() => testFiles("package main\nfunc TestMain(m *testing.M) {}"), /TestMain/);
});

test("all curriculum tests and reference solutions get a real Go testing harness without source rewriting", () => {
  const curriculum = JSON.parse(readFileSync(new URL("../lib/curriculum.json", import.meta.url)));
  const solutions = JSON.parse(readFileSync(new URL("./solutions.json", import.meta.url)));
  for (const exercise of curriculum) {
    for (const code of [exercise.code, solutions[exercise.id]]) {
      assert.deepEqual(testNames(code), ["TestKoan"], exercise.id);
      assert.equal(testFiles(code)["koan.go"], code);
    }
  }
});
