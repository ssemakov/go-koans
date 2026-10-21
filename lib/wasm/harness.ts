import { parser } from "@lezer/go";

/** Discover declarations, so comments, strings, and methods never become tests. */
export function testNames(code: string): string[] {
  const names: string[] = [];
  const root = parser.parse(code).topNode;
  for (let node = root.firstChild; node; node = node.nextSibling) {
    if (node.name !== "FunctionDecl") continue;
    const nameNode = node.getChild("DefName");
    if (!nameNode) continue;
    const name = code.slice(nameNode.from, nameNode.to);
    if (name === "TestMain") throw new Error("TestMain is not supported by the browser runner. Select Go Playground or download the test.");
    if (/^Test($|[^\p{Ll}])/u.test(name)) names.push(name);
  }
  return names;
}

export function testFiles(code: string): Record<string, string> {
  if (code.length > 40_000) throw new Error("Keep the test code under 40,000 characters.");
  const names = testNames(code);
  if (!names.length) throw new Error("No tests found. Keep the TestKoan function and run it again.");
  // Let Go check the signatures. Use the real testing package, including subtests,
  // Cleanup, Skip, and process exit status, rather than emulating testing.T.
  return {
    "koan.go": code,
    "main.go": `package main
import ("regexp"; "testing")
func main() {
  testing.Main(regexp.MatchString, []testing.InternalTest{
    ${names.map((name) => `{Name: ${JSON.stringify(name)}, F: ${name}},`).join("\n    ")}
  }, nil, nil)
}
`,
  };
}
