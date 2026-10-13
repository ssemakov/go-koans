# Go Koans

120 runnable Go exercises in 30 topics. The tests initially fail. Replace the TODO values or implement the missing code, then run the tests again.

The interface has a topic list, code editor, test output, and previous/next controls. Browser edits and progress are saved locally. There are no accounts or services to configure.

## Exercises

The sequence starts with assertions, variables, constants, numbers, types, strings, and arrays before slices. It continues through maps, control flow, loops, functions, closures, pointers, allocation, structs, methods, embedding, interfaces, errors, defer, panic, channels, concurrency, context, IO, files, JSON, regular expressions, generics, standard-library functions, testing, HTTP, and time. Four projects cover triangles, dice scoring, a dice set, and a channel pipeline.

The structure follows [Ruby Koans](https://github.com/edgecase/ruby_koans): run a failing test, change the code, and run the test again. Topic coverage and project types also use [cdarwin/go-koans](https://github.com/cdarwin/go-koans) and [VictorDenisov/go_koans](https://github.com/VictorDenisov/go_koans) as references. These are newly written exercises, with current Go APIs, deterministic concurrency tests, and no assertions about unspecified slice growth or map order. The [Go wiki](https://go.dev/wiki/), language specification, and standard-library documentation are linked as further references. This curriculum covers the language and practical standard-library topics, not every operations or platform page in the wiki.

## Run in the browser

Use **Run tests** or Ctrl/Cmd + Enter. The Next.js route forwards code to the official Go Playground over HTTPS. It reports actual test failures and compiler errors. Skipped tests and runs with no completed tests do not mark an exercise passed. The runtime has the Go Playground's execution and network limits.

Enable **Vim mode** above the editor for Vim keybindings. Use `i` to enter insert mode and `Esc` to return to normal mode. Normal-mode motions, visual selections, search, and undo are available. The setting is saved in this browser; toggling it preserves the code. Edits save automatically, and Ctrl/Cmd + Enter still runs tests.

**Download test** saves your current edits. **Download all tests** creates a ZIP containing the full starter set and a Go module.

## Run locally

Install Go 1.23 or later. In this repository:

```sh
cd koans
go test -v ./001-assert-truth
```

Edit `001-assert-truth/koan_test.go` and run that command again. Each numbered directory is an independent exercise. `go test ./...` runs the full set and initially fails as expected. An individual test downloaded from the browser also runs with `go test -v filename_test.go`.

## Develop the website

Use Node.js 24 and npm.

```sh
npm ci
npm run dev
```

```sh
npm run lint
npm run typecheck
npm run build
npm start
```

The site uses Next.js, TypeScript, and Biome. No secret or application environment variable is required. The runner needs outbound HTTPS access to `go.dev`.

## Validate curriculum changes

With Go 1.23+ on PATH, run `npm run test:curriculum`. This checks that all 120 starters execute and fail, all 120 reference solutions pass, and the web and local starter files match. It works in a temporary directory and leaves learners' files untouched. The reference solutions in `scripts/solutions.json` are validation fixtures and are not imported into the website or download archive.

Canonical web exercises are in `lib/curriculum.json`; matching local exercises are in `koans/`. UI code is in `app/koans.tsx`, and runner/download routes are in `app/api/`.

## Deploy to Vercel

Import `ssemakov/go-koans` from `main`. Use the Next.js framework, repository root, Node.js 24.x, `npm ci`, and `npm run build`. Leave the output directory at its default. Go is not installed or executed in the Vercel function; Go Playground provides execution. Local build and runtime checks do not publish a deployment.
