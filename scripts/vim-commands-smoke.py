"""Audit Vim behavior against a production build, with Linux and Mac keymaps."""
import json
import os
from pathlib import Path
from playwright.sync_api import sync_playwright, expect

BASE = os.environ.get("KOANS_BASE_URL", "http://127.0.0.1:3100")
GO = "package main\n\nfunc demo() {\n\tif true {\n\t}\n}\n"
LONG = "package main\n\n" + "".join(f"// line {number:03d}\n" for number in range(1, 201))
SHORT_VARIABLE = (Path(__file__).resolve().parent.parent / "koans/006-short-variable-declaration/koan_test.go").read_text()

with sync_playwright() as p:
    browser = p.chromium.launch(executable_path="/usr/bin/chromium", headless=True, args=["--no-sandbox"])
    for platform in ["Linux x86_64", "MacIntel"]:
        page = browser.new_page(viewport={"width": 1280, "height": 900})
        page.add_init_script("Object.defineProperty(navigator, 'platform', {get: () => " + json.dumps(platform) + "})")
        errors = []
        page.on("pageerror", lambda error: errors.append(str(error)))
        page.goto(BASE)
        toggle = page.get_by_role("checkbox", name="Vim mode")
        content = page.locator(".cm-content")

        def load(text):
            toggle.uncheck()
            page.get_by_label("Go test code").fill(text)
            toggle.check()
            content.click()
            page.keyboard.press("Escape")
            page.keyboard.type("gg0")

        def state():
            # Read the complete virtualized document and caret without adding an
            # application debug API or changing the editor's selection in tests.
            return page.evaluate("""() => {
                const view = document.querySelector('.cm-content').cmTile.root.view;
                const cursor = view.cm.getCursor();
                return {doc: view.state.doc.toString(), line: cursor.line + 1, ch: cursor.ch,
                        top: view.scrollDOM.scrollTop, insert: view.cm.state.vim.insertMode,
                        visual: view.cm.state.vim.visualMode,
                        collapsed: view.state.selection.ranges.every(range => range.empty)};
            }""")

        def text_is(expected):
            actual = state()["doc"]
            assert actual == expected, (platform, actual, expected)

        def ex(command):
            page.keyboard.press("Escape")
            page.keyboard.type(":" + command)
            page.keyboard.press("Enter")

        def normal_caret(line):
            current = state()
            assert current["line"] == line, (platform, current, line)
            assert not current["insert"] and not current["visual"] and current["collapsed"], (platform, current)
            expect(page.locator(".cm-vim-panel")).to_have_text("--NORMAL--")
            text = current["doc"].split("\n")[line - 1]
            assert 0 <= current["ch"] <= max(0, len(text) - 1)
            # Check the painted caret too: a correct model position alone would
            # miss a cursor drawn on the wrong line after a blank/tab-only line.
            page.wait_for_function("""() => {
                const view = document.querySelector('.cm-content').cmTile.root.view;
                const position = view.coordsAtPos(view.state.selection.main.head);
                const caret = document.querySelector('.cm-fat-cursor')?.getBoundingClientRect();
                return position && caret && Math.abs(position.top - caret.top) < 1;
            }""")

        # Exercise 006: o below t.Errorf, Esc, kk back to if, then cross the
        # inserted blank line and closing braces from several starting columns.
        for start in ["0", "^", "012l", "$"]:
            load(SHORT_VARIABLE)
            page.keyboard.type("19Go"); page.keyboard.press("Escape")
            opened = state()["doc"]
            assert opened.split("\n")[19] == "\t\t\t"
            page.keyboard.type("kk" + start)
            column = state()["ch"]
            for line in [19, 20, 21, 22, 23]:
                page.keyboard.press("j")
                normal_caret(line)
            for line in [22, 21, 20, 19, 18]:
                page.keyboard.press("k")
                normal_caret(line)
            assert state()["ch"] == column
            page.keyboard.type("3j")
            normal_caret(21)
            page.keyboard.type("3k")
            normal_caret(18)
            page.keyboard.press("ArrowDown")
            normal_caret(19)
            page.keyboard.press("ArrowUp")
            normal_caret(18)
            text_is(opened)

        undo_key = "Meta+z" if platform == "MacIntel" else "Control+z"
        redo_keys = ["Meta+Shift+Z", "Meta+y"] if platform == "MacIntel" else ["Control+Shift+Z"]
        for edit, reverted in [
            ("~", "alpha beta\ngamma delta\nomega"),
            ("dd", "alpha beta\ngamma delta\nomega"),
            (">>", "alpha beta\ngamma delta\nomega"),
            ("ciwEDIT", " beta\ngamma delta\nomega"),
        ]:
            for redo_key in redo_keys:
                load("alpha beta\ngamma delta\nomega")
                page.keyboard.type(edit); page.keyboard.press("Escape")
                changed = state()["doc"]
                assert changed != "alpha beta\ngamma delta\nomega"
                page.keyboard.press(undo_key)
                text_is(reverted)
                normal_caret(1)
                page.keyboard.press("j")
                normal_caret(2)
                page.keyboard.press(redo_key)
                text_is(changed)
                normal_caret(state()["line"])

        load("alpha\nbeta")
        page.keyboard.type("vj")
        assert state()["visual"]
        page.keyboard.press(undo_key)
        text_is("alpha\nbeta")
        normal_caret(state()["line"])
        load("alpha\nbeta")
        page.keyboard.type("iEDIT")
        page.keyboard.press(undo_key)
        text_is("alpha\nbeta")
        assert state()["insert"]
        page.keyboard.press(redo_keys[0])
        text_is("EDITalpha\nbeta")
        assert state()["insert"]

        load(LONG)
        page.keyboard.type("50G")
        before = state()
        page.keyboard.press("Control+d")
        page.wait_for_function("top => document.querySelector('.cm-scroller').scrollTop > top + 100", arg=before["top"])
        down = state()
        assert down["line"] > before["line"] + 5
        page.keyboard.press("Control+u")
        page.wait_for_function("top => document.querySelector('.cm-scroller').scrollTop < top - 100", arg=down["top"])
        assert state()["line"] == before["line"]
        text_is(LONG)
        page.keyboard.type("5"); page.keyboard.press("Control+d")
        assert state()["line"] == 55
        page.keyboard.type("5"); page.keyboard.press("Control+u")
        assert state()["line"] == 50
        page.keyboard.press("Control+f")
        assert state()["line"] > 65
        page.keyboard.press("Control+b")
        assert state()["line"] < 65
        text_is(LONG)

        page.keyboard.type("10G")
        active = page.locator(".cm-lineNumbers .cm-activeLineGutter")
        expect(active).to_have_text("10")
        assert active.evaluate("el => el.previousElementSibling.textContent") == "1"
        assert active.evaluate("el => el.nextElementSibling.textContent") == "1"
        page.keyboard.type("j")
        expect(active).to_have_text("11")
        assert active.evaluate("el => el.previousElementSibling.textContent") == "1"
        ex("set nornu")
        assert active.evaluate("el => el.previousElementSibling.textContent") == "10"
        assert active.evaluate("el => el.nextElementSibling.textContent") == "12"
        ex("set rnu")
        assert active.evaluate("el => el.previousElementSibling.textContent") == "1"

        load("alpha\n\tbeta\ngamma")
        page.keyboard.press("Enter")
        assert (state()["line"], state()["ch"]) == (2, 1)
        text_is("alpha\n\tbeta\ngamma")

        load(GO)
        page.keyboard.type("4GA"); page.keyboard.press("Enter")
        page.keyboard.type("value := 1"); page.keyboard.press("Escape")
        text_is(GO.replace("\tif true {\n", "\tif true {\n\t\tvalue := 1\n"))
        load("package main\n\nfunc demo() {}")
        page.keyboard.type("3Gf{a"); page.keyboard.press("Enter")
        page.keyboard.type("value := 1"); page.keyboard.press("Escape")
        text_is("package main\n\nfunc demo() {\n\tvalue := 1\n}")
        for mode in ["normal", "visual"]:
            load("first\nsecond")
            if mode == "visual": page.keyboard.type("vj")
            for key in ["Tab", "Shift+Tab"]:
                page.keyboard.press(key)
                expect(content).to_be_focused()
                text_is("first\nsecond")
        for opening, inserted in [("import (", '"fmt"'), ("type Example struct {", "Value int"), ("func demo(", "value int,")]:
            load("package main\n\n" + opening)
            page.keyboard.type("GA"); page.keyboard.press("Enter")
            page.keyboard.type(inserted); page.keyboard.press("Escape")
            text_is("package main\n\n" + opening + "\n\t" + inserted)
        load("package main\n\nfunc demo() {\n\tif true {\n\t\t")
        page.keyboard.type("GA}"); page.keyboard.press("Enter"); page.keyboard.type("}"); page.keyboard.press("Escape")
        text_is("package main\n\nfunc demo() {\n\tif true {\n\t}\n}")
        load("package main\n\nfunc demo(n int) {\n\tswitch n {")
        page.keyboard.type("GA"); page.keyboard.press("Enter"); page.keyboard.type("case 1:")
        page.keyboard.press("Enter"); page.keyboard.type("return"); page.keyboard.press("Enter")
        page.keyboard.type("default:"); page.keyboard.press("Enter"); page.keyboard.type("return"); page.keyboard.press("Escape")
        text_is("package main\n\nfunc demo(n int) {\n\tswitch n {\n\tcase 1:\n\t\treturn\n\tdefault:\n\t\treturn")
        for command in ["3Go", "5GO"]:
            load("package main\n\nfunc demo() {\n\tvalue := 1\n}\n")
            page.keyboard.type(command); page.keyboard.type("other := 2"); page.keyboard.press("Escape")
            assert "\n\tother := 2\n" in state()["doc"]
        load("package main\n\nfunc demo() {\n\n}")
        page.keyboard.type("4Gi"); page.keyboard.press("Tab"); page.keyboard.press("Tab")
        page.keyboard.press("Shift+Tab"); page.keyboard.type("value := 1"); page.keyboard.press("Escape")
        text_is("package main\n\nfunc demo() {\n\tvalue := 1\n}")

        load("package main\n\nfunc demo() {\nvalue := 1\nother := 2\n}\n")
        page.keyboard.type("4GVj>")
        text_is("package main\n\nfunc demo() {\n\tvalue := 1\n\tother := 2\n}\n")
        page.keyboard.type("u")
        text_is("package main\n\nfunc demo() {\nvalue := 1\nother := 2\n}\n")
        load(GO.replace("\t", ""))
        page.keyboard.type("gg=G")
        text_is(GO)
        load("abc\ndef\nghi")
        page.keyboard.press("Control+v"); page.keyboard.type("jd")
        text_is("bc\nef\nghi")

        load("one two\nthree four")
        page.keyboard.type("ciwONE"); page.keyboard.press("Escape"); page.keyboard.type("j0.")
        text_is("ONE two\nONE four")
        load('say("old")')
        page.keyboard.type('f"ci"new'); page.keyboard.press("Escape")
        text_is('say("new")')
        load("{\n\tvalue\n}")
        page.keyboard.type("%")
        assert state()["line"] == 3
        load("first\nsecond")
        page.keyboard.type('"ayyG"ap')
        text_is("first\nsecond\nfirst")
        load("one\ntwo")
        page.keyboard.type("qaA!"); page.keyboard.press("Escape"); page.keyboard.type("qj@a")
        text_is("one!\ntwo!")
        load("first\nsecond\nthird")
        page.keyboard.type("jmaG'a")
        assert state()["line"] == 2
        load("value := 10")
        page.keyboard.type("f1"); page.keyboard.press("Control+a")
        text_is("value := 11")
        page.keyboard.press("Control+x")
        text_is("value := 10")
        load("one two\ntwo three")
        ex("%s/two/TWO/g")
        text_is("one TWO\nTWO three")

        runs = []
        def run_request(route):
            runs.append(json.loads(route.request.post_data)["code"])
            route.fulfill(content_type="application/json", body=json.dumps({"status": "failed", "output": "Shortcut reached the runner.", "passed": 0, "failed": 1, "skipped": 0}))
        page.route("**/api/run", run_request)
        load(GO)
        for shortcut in (["Control+Enter", "Meta+Enter"] if platform == "MacIntel" else ["Control+Enter"]):
            for mode in ["normal", "insert"]:
                content.click(); page.keyboard.press("Escape")
                if mode == "insert": page.keyboard.type("i")
                count = len(runs)
                with page.expect_response("**/api/run"):
                    page.keyboard.press(shortcut)
                expect(page.locator(".output pre")).to_have_text("Shortcut reached the runner.")
                assert len(runs) == count + 1 and runs[-1] == GO
                text_is(GO)
        content.click(); page.keyboard.press("Escape"); page.keyboard.press("Control+Shift+Enter")
        expect(page.locator("h2")).to_have_text("Assert equality")
        expect(content).to_be_focused()
        page.set_viewport_size({"width": 375, "height": 812})
        assert page.evaluate("document.documentElement.scrollWidth <= innerWidth")
        assert not errors, (platform, errors)
        page.close()
        print(f"{platform}: blank-line navigation/caret, native undo/redo, scrolling, counts, relative numbers, Go indentation, text objects, visual blocks, registers, macros, marks, replacement, and shortcuts passed.", flush=True)
    browser.close()
