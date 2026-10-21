"""Inject a real compiler download failure, then verify resume and reload reuse.

Run against a production server using KOANS_BASE_URL. Requires Python Playwright
and Chromium. A local proxy fails the last standard-library chunk exactly once.
"""
from collections import Counter
from hashlib import sha256
from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer
import json
import os
from pathlib import Path
from threading import Event, Lock, Thread
from urllib.error import HTTPError
from urllib.request import urlopen
from playwright.sync_api import sync_playwright, expect

ROOT = Path(__file__).resolve().parent.parent
BACKEND = os.environ.get("KOANS_BASE_URL", "http://127.0.0.1:3000").rstrip("/")
stdlib = (ROOT / "node_modules/@wasm-oj/toolchain-go/assets/go-1.26.5-wasip1.stdlib.gz.bin").read_bytes()
size = 4 * 1024 * 1024
last_chunk = f"/wasm/chunks/{sha256(stdlib[(len(stdlib)-1)//size*size:]).hexdigest()}.bin"
first_chunk = f"/wasm/chunks/{sha256(stdlib[:size]).hexdigest()}.bin"
second_chunk = f"/wasm/chunks/{sha256(stdlib[size:2*size]).hexdigest()}.bin"
first_gate, second_gate, started = Event(), Event(), Event()
calls = Counter()
lock = Lock()
version = 1

class Proxy(BaseHTTPRequestHandler):
    def log_message(self, *_args):
        pass

    def do_GET(self):
        with lock:
            calls[self.path] += 1
            fail = self.path == last_chunk and calls[self.path] == 1
        if self.path == first_chunk:
            started.set()
            first_gate.wait(30)
        if self.path == second_chunk:
            second_gate.wait(30)
        if fail:
            self.send_response(503)
            self.send_header("Content-Length", "0")
            self.end_headers()
            return
        try:
            try:
                response = urlopen(BACKEND + self.path, timeout=60)
            except HTTPError as error:
                response = error
            with response:
                body = response.read()
                self.send_response(response.status)
                for key, value in response.headers.items():
                    if key.lower() not in ["connection", "transfer-encoding", "content-length"]:
                        self.send_header(key, value)
                self.send_header("Content-Length", str(len(body)))
                self.send_header("X-Download-Test-Version", str(version))
                self.end_headers()
                self.wfile.write(body)
        except (BrokenPipeError, ConnectionResetError):
            pass

server = ThreadingHTTPServer(("127.0.0.1", 0), Proxy)
thread = Thread(target=server.serve_forever, daemon=True)
thread.start()
try:
    with sync_playwright() as p:
        browser = p.chromium.launch(executable_path=os.environ.get("CHROMIUM_PATH", "/usr/bin/chromium"), headless=True, args=["--no-sandbox"])
        context = browser.new_context()
        page = context.new_page()
        page.clock.install()
        base = f"http://127.0.0.1:{server.server_port}"
        page.goto(base)
        run = page.get_by_role("button", name="Run tests", exact=True)
        run.click()
        assert started.wait(20), "The compiler download did not start"
        expect(page.locator(".output pre")).to_contain_text(" MB (")
        page.clock.fast_forward(170_000)
        before = page.locator(".output pre").inner_text()
        first_gate.set()
        page.wait_for_function("previous => { const value = document.querySelector('.output pre').textContent; return value !== previous && value.includes(' MB ('); }", arg=before)
        page.clock.fast_forward(170_000)
        # More than three minutes overall has elapsed, with intervening progress.
        expect(page.get_by_role("button", name="Cancel", exact=True)).to_be_visible()
        expect(page.locator(".result.error")).to_have_count(0)
        second_gate.set()
        expect(page.locator(".result.error")).to_be_visible(timeout=180_000)
        page.wait_for_function("async () => (await (await caches.open('go-koans:wasm-chunks:v1')).keys()).length > 1")
        solutions = json.loads((ROOT / "scripts/solutions.json").read_text())
        page.get_by_label("Go test code", exact=True).fill(solutions["001-assert-truth"])
        run.click()
        expect(page.locator(".result.passed")).to_be_visible(timeout=180_000)
        chunks = {url: count for url, count in calls.items() if "/wasm/chunks/" in url}
        assert chunks[last_chunk] == 2, chunks
        assert all(count == (2 if url == last_chunk else 1) for url, count in chunks.items()), chunks
        context.new_cdp_session(page).send("Network.clearBrowserCache")
        version = 2
        response = page.reload()
        assert response.headers["x-download-test-version"] == "2", "The worker cached the application page"
        run.click()
        expect(page.locator(".result.passed")).to_be_visible(timeout=180_000)
        assert chunks == {url: count for url, count in calls.items() if "/wasm/chunks/" in url}, "Reload fetched compiler chunks again"
        browser.close()
        print(f"Download checks passed: progressing downloads survive >3 minutes, recovered from 503, downloaded each successful chunk once, and reused all {len(chunks)} chunks after HTTP cache clear + reload. App HTML reached the server on reload.")
finally:
    first_gate.set(); second_gate.set()
    server.shutdown()
    server.server_close()
