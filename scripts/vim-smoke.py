"""Check Vim editing against a running production server."""
import json
import os
from pathlib import Path
from playwright.sync_api import sync_playwright, expect

ROOT = Path(__file__).resolve().parent.parent
EXERCISES = json.loads((ROOT / "lib/curriculum.json").read_text())
BASE = os.environ.get("KOANS_BASE_URL", "http://127.0.0.1:3100")

with sync_playwright() as p:
    browser = p.chromium.launch(executable_path="/usr/bin/chromium", headless=True, args=["--no-sandbox"])
    page = browser.new_page(accept_downloads=True, viewport={"width":1280,"height":900})
    errors = []
    page.on("pageerror", lambda error: errors.append(str(error)))
    page.goto(BASE)
    toggle = page.get_by_role("checkbox", name="Vim mode")
    expect(toggle).not_to_be_checked()
    expect(page.get_by_label("Go test code")).to_have_value(EXERCISES[0]["code"])
    toggle.check()
    editor = page.locator(".cm-content")
    expect(editor).to_be_visible()
    expect(page.locator(".cm-vim-panel")).to_contain_text("NORMAL")

    def downloaded_code():
        with page.expect_download() as download:
            page.get_by_role("button", name="Download test", exact=True).click()
        return Path(download.value.path()).read_text()

    editor.click()
    page.keyboard.press("Escape")
    page.keyboard.type("gg0i")
    expect(page.locator(".cm-vim-panel")).to_contain_text("INSERT")
    page.keyboard.type("// Vim edit")
    page.keyboard.press("Enter")
    page.keyboard.press("Escape")
    changed = "// Vim edit\n" + EXERCISES[0]["code"]
    assert downloaded_code() == changed
    editor.click(); page.keyboard.type("u")
    after_undo = downloaded_code()
    assert after_undo == "// Vim edit" + EXERCISES[0]["code"]
    editor.click(); page.keyboard.type("u")
    assert downloaded_code() == EXERCISES[0]["code"]
    editor.click(); page.keyboard.press("Control+r")
    assert downloaded_code() == after_undo
    editor.click(); page.keyboard.press("Control+r")
    assert downloaded_code() == changed
    editor.click(); page.keyboard.type("ggdd")
    assert downloaded_code() == EXERCISES[0]["code"]
    editor.click(); page.keyboard.type("ggP")
    assert downloaded_code() == changed
    editor.click(); page.keyboard.type("ggVj")
    expect(page.locator(".cm-vim-panel")).to_contain_text("VISUAL")
    page.keyboard.press("Escape")
    page.keyboard.type("gg")
    page.keyboard.press("/")
    page.keyboard.type("want")
    page.keyboard.press("Enter")
    page.keyboard.type("wwciwtrue")
    page.keyboard.press("Escape")
    solved = changed.replace("want := false", "want := true")
    assert downloaded_code() == solved
    editor.click(); page.keyboard.press("Control+Enter")
    expect(page.locator(".result.passed")).to_be_visible(timeout=30000)
    expect(page.locator(".progress")).to_have_text("1 / 120 passed")
    toggle.uncheck()
    expect(page.get_by_label("Go test code")).to_have_value(solved)
    expect(page.locator(".progress")).to_have_text("1 / 120 passed")
    toggle.check()
    expect(page.locator(".cm-content")).to_be_visible()
    assert downloaded_code() == solved
    page.get_by_role("button", name="Next exercise", exact=True).click()
    expect(page.get_by_role("heading", name="Assert equality", exact=True)).to_be_visible()
    assert downloaded_code() == EXERCISES[1]["code"]
    page.get_by_role("button", name="Previous", exact=True).click()
    assert downloaded_code() == solved
    page.reload()
    expect(toggle).to_be_checked()
    expect(page.locator(".cm-content")).to_be_visible()
    assert downloaded_code() == solved
    expect(page.locator(".progress")).to_have_text("1 / 120 passed")
    page.once("dialog", lambda dialog: dialog.accept())
    page.get_by_role("button", name="Reset", exact=True).click()
    assert downloaded_code() == EXERCISES[0]["code"]
    expect(page.locator(".progress")).to_have_text("0 / 120 passed")
    editor.click(); page.keyboard.press("Escape"); page.keyboard.type("gg0i")
    page.keyboard.type("// Insert mode shortcut")
    page.keyboard.press("Enter")
    page.keyboard.press("Control+Enter")
    expect(page.locator(".result.failed")).to_be_visible(timeout=30000)
    page.set_viewport_size({"width":375,"height":812})
    assert page.evaluate("document.documentElement.scrollWidth <= innerWidth")
    page.screenshot(path="/tmp/go-koans-vim-mobile.png", full_page=True)
    toggle.uncheck()
    page.reload()
    expect(toggle).not_to_be_checked()
    expect(page.get_by_label("Go test code")).to_have_value("// Insert mode shortcut\n"+EXERCISES[0]["code"])
    assert not errors, errors
    browser.close()
print("Vim checks passed: insert/normal/visual modes, motions, search, delete/paste, undo/redo, run shortcut, toggling, navigation, reset, persistence, and mobile layout.")
