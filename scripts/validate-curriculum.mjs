import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { mkdtemp, mkdir, readFile, rm, writeFile } from "node:fs/promises";
import os from "node:os";
import path from "node:path";

const root = path.resolve(import.meta.dirname, "..");
const curriculum = JSON.parse(await readFile(path.join(root, "lib/curriculum.json"), "utf8"));
const solutions = JSON.parse(await readFile(path.join(root, "scripts/solutions.json"), "utf8"));
assert.equal(new Set(curriculum.map((e) => e.id)).size, curriculum.length);
assert.equal(Object.keys(solutions).length, curriculum.length);
const go = process.env.GO_BINARY || "go";
const scratch = await mkdtemp(path.join(os.tmpdir(), "go-koans-validation-"));

function checkTests(directory, expected, label) {
  const run = spawnSync(go, ["test", "-json", "-count=1", "-timeout=10s", "./..."], {
    cwd: directory, encoding: "utf8", timeout: 180_000, maxBuffer: 10_000_000,
  });
  if (run.error) throw run.error;
  const events = run.stdout.trim().split("\n").filter(Boolean).map((line) => JSON.parse(line));
  const tests = events.filter((e) => e.Test === "TestKoan" && ["pass", "fail", "skip"].includes(e.Action));
  const unexpected = tests.filter((e) => e.Action !== expected);
  assert.equal(tests.length, curriculum.length, `${label}: expected ${curriculum.length} executed tests. ${run.stderr}`);
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
  console.log(`Checking ${curriculum.length} starter exercises...`);
  // Starters run from a temporary copy, preserving learners' working files.
  for (const exercise of curriculum) await writeFile(path.join(scratch, exercise.id, "koan_test.go"), exercise.code);
  checkTests(scratch, "fail", "Starters");
} finally {
  await rm(scratch, { recursive: true, force: true });
}
