"""Regression checks for brace text objects against a running server."""
import os
from playwright.sync_api import sync_playwright, expect

CODE = '''func TestKoan(t *testing.T) {
\toriginal := map[string]int{"a": 1, "b": 2}
\talias := original
\talias["a"] = 9
\tdelete(alias, "b")
\toriginal["c"] = 3
\twantOriginal := map[string]int{} // TODO: all entries.
\twantAlias := map[string]int{}    // TODO: all entries.
\tif !maps.Equal(original, wantOriginal) {
\t\tt.Errorf("original=%v; want %v", original, wantOriginal)
\t}
\tif !maps.Equal(alias, wantAlias) {
\t\tt.Errorf("alias=%v; want %v", alias, wantAlias)
\t}
}'''

with sync_playwright() as p:
    browser = p.chromium.launch(executable_path="/usr/bin/chromium", headless=True, args=["--no-sandbox"])
    page = browser.new_page()
    page.goto(os.environ.get("KOANS_BASE_URL", "http://localhost:3100"))
    toggle = page.get_by_role("checkbox", name="Vim mode")
    editor = page.locator(".cm-content")

    def reset(code=CODE):
        toggle.uncheck()
        page.get_by_label("Go test code").fill(code)
        toggle.check()
        editor.click()
        page.keyboard.press("Escape")

    def code():
        toggle.uncheck()
        return page.get_by_label("Go test code").input_value()

    # Empty objects must enter insert mode without deleting the function.
    for command in ("ci{", "ci}", "ciB"):
        reset()
        page.keyboard.type("7G0" + command)
        expect(page.locator(".cm-vim-panel")).to_contain_text("INSERT")
        page.keyboard.type('"a": 9, "c": 3')
        page.keyboard.press("Escape")
        assert code() == CODE.replace("int{} //", 'int{"a": 9, "c": 3} //')

    reset()
    page.keyboard.type("2G0di{")
    assert code() == CODE.replace('int{"a": 1, "b": 2}', "int{}")

    reset()
    page.keyboard.type("7G0ca{replacement")
    page.keyboard.press("Escape")
    assert code() == CODE.replace("int{} //", "intreplacement //")

    # Repeating the change resolves the next text object at the new cursor.
    reset()
    page.keyboard.type('7G0ci{"a": 9')
    page.keyboard.press("Escape")
    page.keyboard.type("8G0.")
    assert code() == CODE.replace("int{}", 'int{"a": 9}')

    reset()
    page.keyboard.type('7G0ci{"a": 9')
    page.keyboard.press("Escape")
    page.keyboard.type("u")
    assert code() == CODE

    # A cursor already inside a same-line object keeps that object.
    reset('\n'.join(["func f() {", "  x := {first; {second}}", "}"]))
    page.keyboard.type("2G0/first")
    page.keyboard.press("Enter")
    page.keyboard.type("ci{replacement")
    page.keyboard.press("Escape")
    assert code() == "func f() {\n  x := {replacement}\n}"

    # With no following brace on the line, retain the enclosing block.
    reset("func f() {\n  value := 1\n}")
    page.keyboard.type("2G0ci{replacement")
    page.keyboard.press("Escape")
    assert "value" not in code()
    browser.close()
print("Vim brace text-object checks passed.")
