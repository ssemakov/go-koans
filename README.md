# Go Koans

128 runnable Go exercises in 31 topics. The tests initially fail. Follow each exercise's instructions, edit the marked code, and run the tests again.

The interface has a topic list, code editor, test output, and previous/next controls. Browser edits and progress are saved locally. There are no accounts or services to configure.

## Exercises

The sequence starts with assertions, variables, constants, numbers, types, strings, and arrays before slices. It continues through maps, control flow, loops, functions, closures, pointers, allocation, structs, methods, embedding, interfaces, errors, defer, panic, channels, concurrency, context, IO, files, JSON, regular expressions, generics, standard-library functions, testing, HTTP, and time. Four projects cover triangles, dice scoring, a dice set, and a channel pipeline.

Exercises 121–128 form **Bit flags and operations**: binary masks, flag constants with `iota`, setting flags with OR, checking any/all flags with AND, clearing flags with AND NOT, toggling with XOR, complementing a `uint8`, and packing/unpacking four-bit fields with shifts and masks. The tests include empty masks, overlapping flags, the highest bit, and all 256 byte values for field round trips.

There are 33 exercises that ask for expected values, 19 that ask for missing code or inputs, 16 repairs, and 60 implementations. Expected-value exercises compare related states or inputs. Implementations have fixed checks for several cases, including empty inputs, boundaries, ownership, errors, or termination where relevant. Later exercises reuse earlier concepts.

All checks use ordinary Go tests: `testing.T`, `t.Run`, direct comparisons, and standard-library functions such as `slices.Equal`, `maps.Equal`, `errors.Is`, and `errors.As`. There are no assertion helpers or third-party Go packages.

The structure follows [Ruby Koans](https://github.com/edgecase/ruby_koans): run a failing test, change the code, and run the test again. Topic coverage and project types also use [cdarwin/go-koans](https://github.com/cdarwin/go-koans) and [VictorDenisov/go_koans](https://github.com/VictorDenisov/go_koans) as references. These are newly written exercises, with current Go APIs, deterministic concurrency tests, and no assertions about unspecified slice growth or map order. The [Go wiki](https://go.dev/wiki/), language specification, and standard-library documentation are linked as further references. This curriculum covers the language and practical standard-library topics, not every operations or platform page in the wiki.

## Run in the browser

Use the **Theme** menu in the header to choose **System**, **Light**, or **Dark**. System follows your device's appearance; explicit choices are saved in this browser. Both code editors use the selected theme.

Use **Run tests** or Ctrl/Cmd + Enter. **Browser (WebAssembly)** is the default runner: Go 1.26.5 compiles your edited source and runs the real `testing` package in browser workers, using [WASM-OJ Forge](https://github.com/wasm-oj/forge). Your code stays in the browser. The first run downloads roughly 50 MB of compiler, standard library, and compressed runtime assets from this site's origin. Transfers use SHA-256-verified pieces of at most 4 MiB, with at most four network downloads at a time. Completed pieces are saved in browser Cache Storage, shared between SDK workers, and reused after cancellation, a failed request, reload, or an app update. Subsequent runs also reuse cached builds. A first run can still take longer than Go Playground, especially on a slow connection or device.

The browser needs WebAssembly, workers, service workers, SharedArrayBuffer, and cross-origin isolation over HTTPS (localhost works for development). Persistent storage is best effort: browsers can evict it or deny writes, in which case pieces are downloaded again. Unsupported browsers report an error. Select **Go Playground** explicitly to use the existing remote runner, which sends your code through the Next.js API to the official Go Playground. There is no automatic remote fallback.

Both runners report actual test failures and compiler errors. Skipped tests and runs with no completed tests do not mark an exercise passed. The browser runner also checks the process exit status. Cancel, editing, navigation, and switching runners stop the current run and prevent stale results from changing progress.

The WebAssembly target is WASI, with a temporary virtual filesystem and no network sockets or subprocesses. It supports this curriculum's standard-library imports, goroutines, and HTTP handler tests (which use in-memory requests and recorders). Arbitrary native Go programs and third-party modules are outside its scope. Top-level `Test…` functions run through `testing.Main`; custom `TestMain`, example tests, benchmarks, the race detector, and `go vet` are not provided. Download a test to use the full native Go toolchain.

Guest execution is bounded by a 10-second wall and virtual-time limit, a billion weighted instructions, 256 MiB of memory, 256 KiB of output, and an 8 MiB/1,024-entry filesystem write allowance. The Go test timeout is 5 seconds. Browser runs stop after three minutes **without progress**, so an advancing first download can take longer. Each chunk request has a 90-second deadline; failed or corrupt chunks are not saved. Retrying reuses completed chunks. The SDK also bounds compilation independently. These are execution limits, not a guarantee that WASM is faster than a remote runner.

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
| `u`, Ctrl-R, `.` | Undo, redo, or repeat the last edit. |
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
npm test
npm run build
npm start
```

The site uses Next.js, TypeScript, and Biome. No secret or application environment variable is required. `npm run dev` and `npm run build` first bundle the browser adapter, compress the runtime Wasm files, split pinned binaries into content-addressed chunks, and copy license notices into the ignored `public/wasm/` directory. Adapter URLs include a hash of the dependency lockfile and runner sources. Chunk URLs depend on their bytes, allowing unchanged chunks to survive app deployments; only versioned assets receive immutable browser and CDN caching. Restart `npm run dev` after changing the adapter so the assets and URL regenerate together.

Build preparation also generates `public/wasm-cache-sw.js`. This service worker is activated lazily before the first browser run. It reconstructs only the pinned binary/manifest responses required by the SDK, verifies chunk digests, and streams decompressed runtime bytes. It never intercepts navigations, application scripts, exercise content, or API requests. Its script uses `no-store` so deployments update it promptly. Large logical binary paths are reconstructed by the worker; the published files on disk are their smaller chunks. Clearing browser site data removes downloaded chunks and saved edits/progress.

`next.config.ts` supplies COOP/COEP/CORP headers required for SharedArrayBuffer. Preserve them when adding a reverse proxy or changing hosts. The optional Go Playground route needs outbound HTTPS access to `go.dev`; browser execution needs no compilation backend.

With Python Playwright and Chromium at `/usr/bin/chromium`, run `KOANS_BASE_URL=http://127.0.0.1:3000 python scripts/vim-commands-smoke.py` against the running production server to check Vim commands with Linux and Mac key bindings. `scripts/vim-smoke.py` checks editing, an actual Go test run, progress, downloads, persistence, and mobile layout against the same URL.

Run `KOANS_BASE_URL=http://127.0.0.1:3000 python scripts/wasm-smoke.py` to check actual WASM compilation and execution, failures, skips, added tests, panic, temporary files, execution limits, cancellation/retry, navigation, persistence, isolation/cache headers, mobile dark mode, and explicit remote selection. It also verifies that default runs never request `/api/run` or an external origin. Add `--curriculum` to compile and execute all 128 starters and reference solutions in Chromium; every starter must fail and every solution must pass. This takes several minutes. Set `CHROMIUM_PATH` for a different Chromium executable. The smoke checks exercise Chromium; other browser engines still need compatibility testing.

`KOANS_BASE_URL=http://127.0.0.1:3000 python scripts/wasm-download-smoke.py` runs through a local proxy that deliberately fails the last standard-library chunk. It verifies recovery by fetching only missing parts, reuse after clearing the HTTP cache and reloading, and fresh application HTML on reload. It also advances the browser clock beyond three minutes with intervening download progress to check the inactivity deadline. Both browser scripts use actual compiled Go tests.

`npm ci` applies `scripts/patch-vim.mjs` to the Vim dependency so brace text objects prefer a pair later on the current line over a multiline enclosing block. Review this guarded patch when upgrading `@replit/codemirror-vim-core`. With a server on port 3100, Python Playwright and `/usr/bin/chromium` installed, run `python scripts/vim-braces-smoke.py` for the brace regression checks (or set `KOANS_BASE_URL` for another address).

## Validate curriculum changes

With Go 1.23+ and a C compiler on PATH, run `npm run test:curriculum`. This checks that all 128 starters execute and fail, all 128 reference solutions pass, and the web and local starter files match. It also runs concurrency solutions with the race detector and verifies that the tests reject selected incorrect implementations: shared clone storage, lost error identity, dropped IO errors, incorrect table-driven checks, overflow-prone triangle validation, reused dice storage, a pipeline that waits for input closure, incorrect bitwise operators, and swapped or unmasked bit fields. It works in a temporary directory and leaves learners' files untouched.

The reference solutions in `scripts/solutions.json` are validation fixtures and are not imported into the website or download archive. A passing check establishes the tested behavior; it does not measure a learner's retention.

For maintainers, `scripts/author-curriculum.py` defines the exercises and reference edits. After editing it, run `python scripts/author-curriculum.py` with `gofmt` on PATH, then validate. This overwrites the generated starter files in `koans/`. Exercise IDs stay stable, and the `revisits` metadata records earlier exercises used again later.

The authoring script automatically increments an exercise's `version` when its formatted starter code or instructions change. Regenerating unchanged exercises keeps their versions stable. Unversioned browser saves migrate as version 1; the updated Boolean zero value exercise starts at version 2. Versioned saves use separate storage so tabs using the previous storage format cannot overwrite them.

Canonical web exercises are in `lib/curriculum.json`; matching local exercises are in `koans/`. UI code is in `app/koans.tsx`, browser execution is in `lib/wasm/`, and remote runner/download routes are in `app/api/`.

## Deploy to Vercel

Use the Next.js framework, repository root, Node.js 24.x, `npm ci`, and `npm run build`. Leave the output directory at its default. Build preparation emits about 48 MB of static assets, with every binary transfer at most 4 MiB. These are served as static files with immutable CDN caching and are excluded from the runner function. No native Go installation is required on Vercel. Feature branches can be tested with Vercel preview deployments without merging into `main`. Local build and runtime checks do not publish a deployment.
