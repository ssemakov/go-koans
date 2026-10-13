"""Run against a production server with Python Playwright and Chromium installed."""
import json
import os
from pathlib import Path
from playwright.sync_api import sync_playwright, expect

ROOT = Path(__file__).resolve().parent.parent
BASE = os.environ.get("KOANS_BASE_URL", "http://127.0.0.1:3100")
EXERCISES = json.loads((ROOT / "lib/curriculum.json").read_text())

with sync_playwright() as p:
    browser = p.chromium.launch(executable_path="/usr/bin/chromium", headless=True, args=["--no-sandbox"])
    context = browser.new_context(accept_downloads=True, viewport={"width": 1280, "height": 900})
    page = context.new_page()
    errors = []
    page.on("pageerror", lambda error: errors.append(str(error)))
    page.goto(BASE)
    expect(page.get_by_role("heading", name="Assert truth", exact=True)).to_be_visible()
    assert page.locator(".curriculum li button").count() == 120
    assert page.locator(".curriculum details").count() == 30
    editor = page.get_by_label("Go test code")
    expect(editor).to_have_value(EXERCISES[0]["code"])
    page.get_by_role("button", name="Run tests", exact=True).click()
    expect(page.locator(".result.failed")).to_be_visible(timeout=30000)
    expect(page.locator(".output pre")).to_contain_text("got true; want false")
    editor.fill(EXERCISES[0]["code"].replace("want := false", "want := true"))
    page.get_by_role("button", name="Run tests", exact=True).click()
    expect(page.locator(".result.passed")).to_be_visible(timeout=30000)
    expect(page.locator(".progress")).to_have_text("1 / 120 passed")
    page.reload()
    expect(editor).to_have_value(EXERCISES[0]["code"].replace("want := false", "want := true"))
    expect(page.locator(".progress")).to_have_text("1 / 120 passed")
    with page.expect_download() as download:
        page.get_by_role("button", name="Download test", exact=True).click()
    assert Path(download.value.path()).read_text() == editor.input_value()
    page.once("dialog", lambda dialog: dialog.accept())
    page.get_by_role("button", name="Reset", exact=True).click()
    expect(editor).to_have_value(EXERCISES[0]["code"])
    expect(page.locator(".progress")).to_have_text("0 / 120 passed")
    editor.fill("package main\nfunc main() {}")
    page.get_by_role("button", name="Run tests", exact=True).click()
    expect(page.locator(".result.error")).to_be_visible(timeout=30000)
    expect(page.locator(".output pre")).to_contain_text("No tests completed")
    editor.fill('package main\nimport "testing"\nfunc TestKoan(t *testing.T) { t.Skip("skip") }')
    page.get_by_role("button", name="Run tests", exact=True).click()
    expect(page.locator(".result.failed")).to_be_visible(timeout=30000)
    expect(page.locator(".progress")).to_have_text("0 / 120 passed")
    editor.fill('package main\nimport "testing"\nfunc TestKoan(t *testing.T) { invalid code }')
    page.get_by_role("button", name="Run tests", exact=True).click()
    expect(page.locator(".result.failed")).to_be_visible(timeout=30000)
    expect(page.locator(".output pre")).to_contain_text("syntax error")
    page.route("**/api/run", lambda route: route.fulfill(status=503, content_type="application/json", body=json.dumps({"status":"error","output":"Go Playground is unavailable.","passed":0,"failed":0,"skipped":0})))
    page.get_by_role("button", name="Run tests", exact=True).click()
    expect(page.locator(".result.error")).to_be_visible()
    expect(page.locator(".progress")).to_have_text("0 / 120 passed")
    page.unroute("**/api/run")
    page.get_by_role("button", name="Next exercise", exact=True).click()
    expect(page.get_by_role("heading", name="Assert equality", exact=True)).to_be_visible()
    expect(editor).to_have_value(EXERCISES[1]["code"])
    page.reload()
    expect(page.get_by_role("heading", name="Assert equality", exact=True)).to_be_visible()
    assert page.evaluate("document.documentElement.scrollWidth <= window.innerWidth")
    page.screenshot(path="/tmp/go-koans-desktop.png", full_page=True)
    page.set_viewport_size({"width":375,"height":812})
    assert page.evaluate("document.documentElement.scrollWidth <= window.innerWidth")
    page.screenshot(path="/tmp/go-koans-mobile.png", full_page=True)
    assert not errors, errors
    context.close()
    # Corrupt storage must not make the editor unusable.
    context = browser.new_context()
    context.add_init_script("localStorage.setItem('go-koans:exercises:v1', '{invalid')")
    page = context.new_page(); page.goto(BASE)
    expect(page.get_by_role("button", name="Run tests", exact=True)).to_be_enabled()
    expect(page.locator(".notice")).to_contain_text("Saved edits could not be read")
    browser.close()
print("Browser checks passed: failing/passing tests, editing, persistence, reset, download, compiler errors, skipped/zero-test rejection, outage handling, navigation, mobile layout, and corrupt storage.")
