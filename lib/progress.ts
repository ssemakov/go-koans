export type Entry = { code: string; passed: boolean; version: number };

type ExerciseVersion = { id: string; version: number };

export function restoreEntries(saved: unknown, exercises: readonly ExerciseVersion[], unversioned = false) {
  const entries: Record<string, Entry> = {};
  let updated = 0;
  if (!saved || typeof saved !== "object" || Array.isArray(saved)) return { entries, updated };

  for (const exercise of exercises) {
    const entry: unknown = (saved as Record<string, unknown>)[exercise.id];
    if (!entry || typeof entry !== "object" || !("code" in entry) || typeof entry.code !== "string" || entry.code.length > 40_000) continue;
    // The v2 storage format predates exercise versions and contains revision 1.
    const version = "version" in entry ? entry.version : unversioned ? 1 : undefined;
    if (version !== exercise.version) { updated++; continue; }
    entries[exercise.id] = { code: entry.code, passed: "passed" in entry && entry.passed === true, version: exercise.version };
  }
  return { entries, updated };
}
