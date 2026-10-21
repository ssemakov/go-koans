"""Exercise the real browser compiler against a running production server.

KOANS_BASE_URL=http://127.0.0.1:3000 python scripts/wasm-smoke.py [--curriculum]
Requires Python Playwright and Chromium (CHROMIUM_PATH overrides its path).
No Google Playground or external execution service is used by these checks.
"""
import json
import os
import re
import sys
import time
from pathlib import Path
from urllib.parse import urlparse
from playwright.sync_api import sync_playwright, expect

ROOT = Path(__file__).resolve().parent.parent
BASE = os.environ.get("KOANS_BASE_URL", "http://127.0.0.1:3000")
CURRICULUM = json.loads((ROOT / "lib/curriculum.json").read_text())
SOLUTIONS = json.loads((ROOT / "scripts/solutions.json").read_text())

with sync_playwright() as p:
    browser = p.chromium.launch(executable_path=os.environ.get("CHROMIUM_PATH", "/usr/bin/chromium"), headless=True, args=["--no-sandbox"])
    context = browser.new_context()
    page = context.new_page()
    requests, errors = [], []
    context.on("request", lambda request: requests.append(request.url))
    page.on("pageerror", lambda error: errors.append(str(error)))
    response = page.goto(BASE)
    assert response.headers["cross-origin-opener-policy"] == "same-origin"
    assert response.headers["cross-origin-embedder-policy"] == "require-corp"
    assert page.evaluate("crossOriginIsolated && typeof SharedArrayBuffer !== 'undefined'")
    expect(page.get_by_role("combobox", name="Runner", exact=True)).to_have_value("wasm")
    editor = page.get_by_label("Go test code", exact=True)
    run = page.get_by_role("button", name="Run tests", exact=True)

    def execute(code, status, output=None):
        editor.fill(code)
        run.click()
        expect(page.locator(".result")).to_be_visible(timeout=180_000)
        expect(page.locator(f".result.{status}")).to_be_visible()
        if output:
            expect(page.locator(".output pre")).to_contain_text(output)
        print(f"Browser check: {status} ({output or 'test run'})", flush=True)

    # Cancelling initial preparation must not mark a test passed or poison retry.
    run.click()
    page.get_by_role("button", name="Cancel", exact=True).click()
    expect(run).to_be_enabled()
    expect(page.locator(".result")).to_have_count(0)
    start = time.monotonic()
    execute(CURRICULUM[0]["code"], "failed", "--- FAIL: TestKoan")
    cold = time.monotonic() - start
    start = time.monotonic()
    execute(SOLUTIONS[CURRICULUM[0]["id"]], "passed", "--- PASS: TestKoan")
    warm = time.monotonic() - start
    expect(page.locator(".progress")).to_have_text(f"1 / {len(CURRICULUM)} passed")
    downloaded = [url for url in requests if "/wasm/chunks/" in url]
    assert downloaded and len(downloaded) == len(set(downloaded)), downloaded
    # Discard HTTP cache so reload tests exercise persistent chunk storage.
    context.new_cdp_session(page).send("Network.clearBrowserCache")
    before_reload = len(requests)
    page.reload()
    expect(editor).to_have_value(SOLUTIONS[CURRICULUM[0]["id"]])
    expect(page.locator(".progress")).to_have_text(f"1 / {len(CURRICULUM)} passed")

    execute('package main\nimport "testing"\nfunc TestKoan(t *testing.T) { missing() }', "failed", "undefined")
    assert not any("/wasm/chunks/" in url for url in requests[before_reload:]), "Reload downloaded compiler bytes again"
    execute('package main\nimport "testing"\nfunc TestKoan(t *testing.T) {t.Skip("skip")}', "failed", "--- SKIP")
    execute('package main\nfunc main() {println("PASS")}', "error", "No tests found")
    expect(page.locator(".progress")).to_have_text(f"0 / {len(CURRICULUM)} passed")
    execute('package main\nimport "testing"\nfunc TestKoan(t *testing.T) {t.Run("child",func(t *testing.T){t.Fatal("child failed")})}', "failed", "child failed")
    execute('package main\nimport "testing"\nfunc TestKoan(t *testing.T) {}\nfunc TestAdded(t *testing.T) {t.Fatal("added test failed")}', "failed", "added test failed")
    execute('package main\nimport "testing"\nfunc TestKoan(t *testing.T) {panic("panic failed")}', "failed", "panic failed")
    execute('package main\nimport("testing";"os")\nfunc TestKoan(t *testing.T) {t.Cleanup(func(){if err:=os.WriteFile(t.TempDir()+"/test", []byte("ok"),0600); err!=nil {t.Fatal(err)}})}', "passed")
    execute('package main\nimport "testing"\nfunc TestKoan(t *testing.T) {for {}}', "error", "Browser runner stopped")

    # Terminate active compiler/runtime preparation, then retry. Guest loops
    # can hit the deterministic instruction limit before a UI polling interval.
    editor.fill('package main\nimport "testing"\nfunc TestKoan(t *testing.T) {for {}}')
    run.click()
    expect(page.locator(".output pre")).to_contain_text(re.compile("Compiling Go tests|Preparing compiler and runtime"), timeout=120_000)
    page.get_by_role("button", name="Cancel", exact=True).click()
    execute(SOLUTIONS[CURRICULUM[0]["id"]], "passed")
    editor.fill(CURRICULUM[0]["code"])
    run.click()
    page.get_by_role("button", name="Next exercise", exact=True).click()
    expect(page.get_by_role("heading", name=CURRICULUM[1]["title"], exact=True)).to_be_visible()
    expect(run).to_be_enabled()
    expect(page.locator(".result")).to_have_count(0)

    # Everything fetched so far is a same-origin asset. User code stays local.
    assert not any("/api/run" in url for url in requests), requests
    external = [url for url in requests if urlparse(url.removeprefix("blob:")).netloc != urlparse(BASE).netloc]
    assert not external, external
    module_url = next(url for url in requests if "/wasm/" in url and url.endswith("/runtime/client.js"))
    asset = context.request.get(module_url)
    assert "immutable" in asset.headers["cache-control"]
    worker = context.request.get(f"{BASE.rstrip('/')}/wasm-cache-sw.js")
    assert worker.headers["cache-control"] == "no-store"

    if "--curriculum" in sys.argv:
        for exercise in CURRICULUM:
            for kind, code in [("starter", exercise["code"]), ("solution", SOLUTIONS[exercise["id"]])]:
                result = page.evaluate("""async ({url,code}) => {
                    const runner = await import(url);
                    return runner.runWasmTests(code, AbortSignal.timeout(180000), () => {});
                }""", {"url": module_url, "code": code})
                expected = "failed" if kind == "starter" else "passed"
                assert result["status"] == expected, (exercise["id"], kind, result)
            print(f"{exercise['id']}: starter fails, solution passes", flush=True)

    # The remote runner is used only when explicitly selected (stub the HTTP API).
    remote = []
    def playground(route):
        remote.append(route.request.post_data_json)
        route.fulfill(json={"status": "passed", "output": "Explicit Playground selection", "passed": 1, "failed": 0, "skipped": 0})
    page.route("**/api/run", playground)
    page.get_by_role("combobox", name="Runner", exact=True).select_option("playground")
    run.click()
    expect(page.locator(".result.passed")).to_be_visible()
    assert remote == [{"code": CURRICULUM[1]["code"]}]
    page.get_by_role("combobox", name="Runner", exact=True).select_option("wasm")
    page.get_by_role("combobox", name="Theme", exact=True).select_option("dark")
    page.set_viewport_size({"width": 375, "height": 812})
    assert page.evaluate("document.documentElement.scrollWidth <= innerWidth")
    page.screenshot(path="/tmp/go-koans-wasm-mobile.png", full_page=True)

    # A host missing isolation headers must explain the manual alternative.
    unisolated = context.new_page()
    def remove_isolation(route):
        response = route.fetch()
        headers = dict(response.headers)
        headers.pop("cross-origin-opener-policy", None)
        headers.pop("cross-origin-embedder-policy", None)
        route.fulfill(response=response, headers=headers)
    unisolated.route(f"{BASE.rstrip('/')}/", remove_isolation)
    unisolated.goto(BASE)
    assert not unisolated.evaluate("crossOriginIsolated")
    unisolated.get_by_role("button", name="Run tests", exact=True).click()
    expect(unisolated.locator(".result.error")).to_be_visible()
    expect(unisolated.locator(".output pre")).to_contain_text("cross-origin isolation")
    assert not errors, errors
    browser.close()
    print(f"WASM smoke checks passed. Cold run {cold:.1f}s; warm edit + run {warm:.1f}s (local Chromium; not a performance guarantee).")
