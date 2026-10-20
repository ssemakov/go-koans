"""Audit Vim behavior against a production build, with Linux and Mac keymaps."""
import json
import os
from playwright.sync_api import sync_playwright, expect

BASE = os.environ.get("KOANS_BASE_URL", "http://127.0.0.1:3100")
GO = "package main\n\nfunc demo() {\n\tif true {\n\t}\n}\n"
LONG = "package main\n\n" + "".join(f"// line {number:03d}\n" for number in range(1, 201))

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
                        top: view.scrollDOM.scrollTop, insert: view.cm.state.vim.insertMode};
            }""")

        def text_is(expected):
            actual = state()["doc"]
            assert actual == expected, (platform, actual, expected)

        def ex(command):
            page.keyboard.press("Escape")
            page.keyboard.type(":" + command)
            page.keyboard.press("Enter")

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
        print(f"{platform}: scrolling, counts, relative numbers, Go indentation, text objects, visual blocks, registers, macros, marks, replacement, and shortcuts passed.", flush=True)
    browser.close()
