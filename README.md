# Go Koans

128 runnable Go exercises in 31 topics. The tests initially fail. Follow each exercise's instructions, edit the marked code, and run the tests again.

The interface has a topic list, code editor, test output, and previous/next controls. Browser edits and progress are saved locally. There are no accounts or services to configure.

## Exercises

The sequence starts with assertions, variables, constants, numbers, types, strings, and arrays before slices. It continues through maps, control flow, loops, functions, closures, pointers, allocation, structs, methods, embedding, interfaces, errors, defer, panic, channels, concurrency, context, IO, files, JSON, regular expressions, generics, standard-library functions, testing, HTTP, and time. Four projects cover triangles, dice scoring, a dice set, and a channel pipeline.

Exercises 121–128 form **Bit flags and operations**: binary masks, flag constants with `iota`, setting flags with OR, checking any/all flags with AND, clearing flags with AND NOT, toggling with XOR, complementing a `uint8`, and packing/unpacking four-bit fields with shifts and masks. The tests include empty masks, overlapping flags, the highest bit, and all 256 byte values for field round trips.

There are 32 exercises that ask for expected values, 19 that ask for missing code or inputs, 17 repairs, and 60 implementations. Expected-value exercises compare related states or inputs. Implementations have fixed checks for several cases, including empty inputs, boundaries, ownership, errors, or termination where relevant. Later exercises reuse earlier concepts.

All checks use ordinary Go tests: `testing.T`, `t.Run`, direct comparisons, and standard-library functions such as `slices.Equal`, `maps.Equal`, `errors.Is`, and `errors.As`. There are no assertion helpers or third-party Go packages.

The structure follows [Ruby Koans](https://github.com/edgecase/ruby_koans): run a failing test, change the code, and run the test again. Topic coverage and project types also use [cdarwin/go-koans](https://github.com/cdarwin/go-koans) and [VictorDenisov/go_koans](https://github.com/VictorDenisov/go_koans) as references. These are newly written exercises, with current Go APIs, deterministic concurrency tests, and no assertions about unspecified slice growth or map order. The [Go wiki](https://go.dev/wiki/), language specification, and standard-library documentation are linked as further references. This curriculum covers the language and practical standard-library topics, not every operations or platform page in the wiki.

## Run in the browser

Use **Run tests** or Ctrl/Cmd + Enter. The Next.js route forwards code to the official Go Playground over HTTPS. It reports actual test failures and compiler errors. Skipped tests and runs with no completed tests do not mark an exercise passed. The runtime has the Go Playground's execution and network limits.

Each exercise's documentation links open the exact specification sections, API references, or tutorials for its topic in new tabs. Exercise 008 has separate links for [constant expressions and exact arithmetic](https://go.dev/ref/spec#Constant_expressions) and [explicit type conversions](https://go.dev/ref/spec#Conversions).

Use Ctrl/Cmd + Shift + Enter to open the next exercise in either editor mode. The shortcut preserves your edits and keeps focus in the editor. It stops at the last exercise.

The rewritten curriculum uses separate browser progress storage. Earlier saved edits are preserved and can be exported with **Download previous edits**. Old completion markers do not carry over to the new exercises.

Each exercise displays its version. When its starter code or instructions change, the app loads the updated starter and clears saved edits and completion for that exercise. Edits and progress for unchanged exercises are preserved.

Enable **Vim mode** above the editor for Vim keybindings. Use `i` to enter insert mode and `Esc` to return to normal mode. Normal-mode motions, visual selections, search, and undo are available. The setting is saved in this browser; toggling it preserves the code. Edits save automatically, and Ctrl/Cmd + Enter still runs tests.

Vim mode uses relative line numbers, with the current line's absolute number highlighted. Use counts such as `5j`, `3k`, or `4dd` with the gutter. `:set nornu` switches to absolute numbers; `:set rnu` restores relative numbers for this editor session.

| Commands | Behavior |
| --- | --- |
| `h j k l`, `w b e`, `0 ^ $`, `gg G`, `%` | Move by character, word, line, document, or matching bracket. |
| Ctrl-D / Ctrl-U, Ctrl-F / Ctrl-B | Scroll half a screen or a full screen; counts work with half-screen scrolling. |
| `i a I A`, `o O` | Insert text or open a line with Go indentation. |
| Enter | In Normal mode, move to the next line's first nonblank character. In Insert mode, insert a newline with Go indentation, including nested blocks, braces, imports, and `case` clauses. |
| Tab / Shift-Tab | In Insert mode, insert a tab or reduce indentation. These keys keep focus in the editor in Normal and Visual modes too. |
| `v`, `V`, Ctrl-V | Select characters, lines, or a rectangular block. |
| `d c y`, `p P`, `ciw`, `ci"` | Delete, change, copy, paste, or edit inside a word or quotes. |
| `> <`, `=` | Shift indentation or reindent a selection/motion; `gg=G` reindents the document with tabs. |
| `/`, `?`, `n N`, `:%s/old/new/g` | Search forward/backward, repeat a search, or replace text. |
| `u`, Ctrl-R, Ctrl/Cmd-Z, Ctrl/Cmd-Shift-Z, `.` | Undo, redo, or repeat the last edit. Desktop undo/redo shortcuts keep a normal caret in Normal mode and preserve Insert mode while typing. |
| `"ayy`, `"ap`, `qa…q`, `@a`, `ma`, `'a` | Use named registers, record/replay macros, or set/jump to marks. |
| Ctrl-A / Ctrl-X | Increment or decrement a number. |

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

With Python Playwright and Chromium at `/usr/bin/chromium`, run `KOANS_BASE_URL=http://127.0.0.1:3000 python scripts/vim-commands-smoke.py` against the running production server to check Vim commands with Linux and Mac key bindings. `scripts/vim-smoke.py` checks editing, an actual Go test run, progress, downloads, persistence, and mobile layout against the same URL.

## Validate curriculum changes

With Go 1.23+ and a C compiler on PATH, run `npm run test:curriculum`. This checks that all 128 starters execute and fail, all 128 reference solutions pass, and the web and local starter files match. It also runs concurrency solutions with the race detector and verifies that the tests reject selected incorrect implementations: shared clone storage, lost error identity, dropped IO errors, incorrect table-driven checks, overflow-prone triangle validation, reused dice storage, a pipeline that waits for input closure, incorrect bitwise operators, and swapped or unmasked bit fields. It works in a temporary directory and leaves learners' files untouched.

The reference solutions in `scripts/solutions.json` are validation fixtures and are not imported into the website or download archive. A passing check establishes the tested behavior; it does not measure a learner's retention.

For maintainers, `scripts/author-curriculum.py` defines the exercises and reference edits. After editing it, run `python scripts/author-curriculum.py` with `gofmt` on PATH, then validate. This overwrites the generated starter files in `koans/`. Exercise IDs stay stable, and the `revisits` metadata records earlier exercises used again later.

The authoring script automatically increments an exercise's `version` when its formatted starter code or instructions change. Regenerating unchanged exercises keeps their versions stable. Unversioned browser saves migrate as version 1; the updated Boolean zero value exercise starts at version 2. Versioned saves use separate storage so tabs using the previous storage format cannot overwrite them.

Canonical web exercises are in `lib/curriculum.json`; matching local exercises are in `koans/`. UI code is in `app/koans.tsx`, and runner/download routes are in `app/api/`.

## Deploy to Vercel

Import `ssemakov/go-koans` from `main`. Use the Next.js framework, repository root, Node.js 24.x, `npm ci`, and `npm run build`. Leave the output directory at its default. Go is not installed or executed in the Vercel function; Go Playground provides execution. Local build and runtime checks do not publish a deployment.
