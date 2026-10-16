import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { mkdtemp, mkdir, readFile, rm, writeFile } from "node:fs/promises";
import os from "node:os";
import path from "node:path";
import { mutations } from "./curriculum-mutations.mjs";

const root = path.resolve(import.meta.dirname, "..");
const curriculum = JSON.parse(await readFile(path.join(root, "lib/curriculum.json"), "utf8"));
const solutions = JSON.parse(await readFile(path.join(root, "scripts/solutions.json"), "utf8"));
assert.equal(new Set(curriculum.map((e) => e.id)).size, curriculum.length);
assert.equal(Object.keys(solutions).length, curriculum.length);
for (const exercise of curriculum) {
  assert.ok(Number.isSafeInteger(exercise.version) && exercise.version > 0, `${exercise.id}: missing exercise version`);
  assert.ok(["observe", "complete", "implement", "repair"].includes(exercise.kind), `${exercise.id}: missing exercise kind`);
  for (const earlier of exercise.revisits) {
    assert.ok(curriculum.some((e) => e.id === earlier) && earlier < exercise.id, `${exercise.id}: invalid earlier exercise`);
  }
  assert.doesNotMatch(exercise.code, /\b(?:assert|Assert)(?:Equal|True|False|Nil)?\s*\(/, `${exercise.id}: use ordinary Go checks`);
}
const go = process.env.GO_BINARY || "go";
const scratch = await mkdtemp(path.join(os.tmpdir(), "go-koans-validation-"));

function checkTests(directory, expected, label, count = curriculum.length, flags = [], patterns = ["./..."]) {
  const run = spawnSync(go, ["test", "-json", "-count=1", "-timeout=10s", "-p=4", ...flags, ...patterns], {
    cwd: directory, encoding: "utf8", timeout: 180_000, maxBuffer: 10_000_000,
  });
  if (run.error) throw run.error;
  const events = run.stdout.trim().split("\n").filter(Boolean).map((line) => JSON.parse(line));
  const tests = events.filter((e) => e.Test === "TestKoan" && ["pass", "fail", "skip"].includes(e.Action));
  const unexpected = tests.filter((e) => e.Action !== expected);
  assert.equal(tests.length, count, `${label}: expected ${count} executed tests. ${run.stderr || run.stdout}`);
  if (unexpected.length) {
    for (const test of unexpected) console.error(events.filter((e) => e.Package === test.Package && e.Output).map((e) => e.Output).join(""));
  }
  assert.equal(unexpected.length, 0, `${label}: unexpected results`);
  assert.equal(run.status, expected === "pass" ? 0 : 1, `${label}: unexpected command exit status`);
  console.log(`${label}: ${tests.length} tests ${expected === "pass" ? "passed" : "failed as expected"}.`);
}

try {
  await writeFile(path.join(scratch, "go.mod"), "module koans.validation\n\ngo 1.23\n");
  for (const exercise of curriculum) {
    const offline = await readFile(path.join(root, "koans", exercise.id, "koan_test.go"), "utf8");
    assert.equal(offline, exercise.code, `${exercise.id}: offline exercise differs from web exercise`);
    const directory = path.join(scratch, exercise.id);
    await mkdir(directory);
    await writeFile(path.join(directory, "koan_test.go"), solutions[exercise.id]);
  }
  console.log(`Checking ${curriculum.length} reference solutions...`);
  checkTests(scratch, "pass", "Solutions");
  const concurrency = curriculum.filter((exercise) => [91, 92, 93, 94, 120].includes(Number(exercise.id.slice(0, 3))));
  console.log("Checking concurrent solutions with the race detector...");
  checkTests(scratch, "pass", "Concurrent solutions", concurrency.length, ["-race"], concurrency.map((e) => `./${e.id}`));
  const mutants = path.join(scratch, "incorrect-implementations");
  await mkdir(mutants);
  await writeFile(path.join(mutants, "go.mod"), "module koans.incorrect\n\ngo 1.23\n");
  for (const mutation of mutations) {
    const exercise = curriculum.find((e) => Number(e.id.slice(0, 3)) === mutation.number);
    assert.ok(exercise, `Unknown mutation exercise ${mutation.number}`);
    let code = solutions[exercise.id];
    for (const [before, after] of mutation.edits) {
      assert.equal(code.split(before).length, 2, `${mutation.name}: mutation must change exactly one occurrence`);
      code = code.replace(before, after);
    }
    const directory = path.join(mutants, `${exercise.id}-${mutation.name}`);
    await mkdir(directory);
    await writeFile(path.join(directory, "koan_test.go"), code);
  }
  console.log(`Checking rejection of ${mutations.length} incorrect implementations...`);
  checkTests(mutants, "fail", "Incorrect implementations", mutations.length, ["-race"]);
  console.log(`Checking ${curriculum.length} starter exercises...`);
  // Starters run from a temporary copy, preserving learners' working files.
  for (const exercise of curriculum) await writeFile(path.join(scratch, exercise.id, "koan_test.go"), exercise.code);
  checkTests(scratch, "fail", "Starters");
} finally {
  await rm(scratch, { recursive: true, force: true });
}
