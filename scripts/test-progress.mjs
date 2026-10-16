import assert from "node:assert/strict";
import { test } from "node:test";
import { restoreEntries } from "../lib/progress.ts";

const exercises = [{ id: "unchanged", version: 1 }, { id: "updated", version: 2 }];

test("an updated exercise loses its edits and completion while unchanged work survives", () => {
  const saved = {
    unchanged: { code: "// my solution", passed: true, version: 1 },
    updated: { code: "// outdated solution", passed: true, version: 1 },
  };
  const restored = restoreEntries(saved, exercises);
  assert.deepEqual(restored, { entries: { unchanged: saved.unchanged }, updated: 1 });
  assert.equal(saved.updated.passed, true);
});

test("edits and completion saved against the current version survive subsequent loads", () => {
  const saved = {
    unchanged: { code: "// in progress", passed: false, version: 1 },
    updated: { code: "// new solution", passed: true, version: 2 },
  };
  assert.deepEqual(restoreEntries(JSON.parse(JSON.stringify(saved)), exercises), { entries: saved, updated: 0 });
});

test("unversioned v2 saves migrate selectively even if the browser skipped the versioning release", () => {
  const saved = {
    unchanged: { code: "// existing solution", passed: true },
    updated: { code: "// existing outdated solution", passed: true },
  };
  const migrated = restoreEntries(saved, exercises, true);
  assert.deepEqual(migrated, {
    entries: { unchanged: { ...saved.unchanged, version: 1 } }, updated: 1,
  });
  assert.deepEqual(restoreEntries(migrated.entries, exercises), { entries: migrated.entries, updated: 0 });
});

test("versioned storage requires versions and also rejects saves from a different newer version", () => {
  assert.deepEqual(restoreEntries({ unchanged: { code: "// no version", passed: true } }, exercises), { entries: {}, updated: 1 });
  assert.deepEqual(restoreEntries({ updated: { code: "// newer solution", passed: true, version: 3 } }, exercises), { entries: {}, updated: 1 });
});

test("malformed, oversized, and removed exercise entries cannot restore progress", () => {
  for (const saved of [null, [], true, "invalid", 1]) {
    assert.deepEqual(restoreEntries(saved, exercises), { entries: {}, updated: 0 });
  }
  for (const invalid of [null, [], { code: 1 }, { code: "x".repeat(40_001), version: 1 }]) {
    assert.deepEqual(restoreEntries({ unchanged: invalid }, exercises), { entries: {}, updated: 0 });
  }
  const saved = { unchanged: { code: "", passed: "true", version: 1 }, removed: { code: "// deleted exercise", passed: true, version: 1 } };
  assert.deepEqual(restoreEntries(saved, exercises), {
    entries: { unchanged: { code: "", passed: false, version: 1 } }, updated: 0,
  });
});
