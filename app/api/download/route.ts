import { createArchive } from "@/lib/archive";
import curriculum from "@/lib/curriculum.json";

export const runtime = "nodejs";

export function GET() {
  const archive = createArchive([
    { name: "go-koans/go.mod", content: "module go.koans.local/exercises\n\ngo 1.23\n" },
    { name: "go-koans/README.txt", content: `Go Koans\n\n${curriculum.length} exercises. The tests initially fail.\nInstall Go 1.23 or later. Run the first test:\n\n  go test -v ./${curriculum[0].id}\n\nEdit the TODO value or missing implementation, then run the test again.\nEach numbered directory is an independent exercise.\nRun all exercises with: go test ./...\n\nReferences:\nhttps://github.com/edgecase/ruby_koans\nhttps://github.com/cdarwin/go-koans\nhttps://github.com/VictorDenisov/go_koans\nhttps://go.dev/wiki/\n` },
    ...curriculum.map((exercise) => ({ name: `go-koans/${exercise.id}/koan_test.go`, content: exercise.code })),
  ]);
  return new Response(new Uint8Array(archive), { headers: { "Content-Type": "application/zip", "Content-Disposition": 'attachment; filename="go-koans.zip"' } });
}
