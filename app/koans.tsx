"use client";

import dynamic from "next/dynamic";
import { useEffect, useRef, useState } from "react";
import curriculum from "@/lib/curriculum.json";
import type { RunResult } from "@/lib/runner";

type Entry = { code: string; passed: boolean };
const topics = [...new Set(curriculum.map((exercise) => exercise.topic))];
const storageKey = "go-koans:exercises:v1";
const vimStorageKey = "go-koans:vim:v1";
const VimEditor = dynamic(() => import("./vim-editor"), { ssr: false, loading: () => <p role="status">Loading editor…</p> });

export default function Koans() {
  const [selected, setSelected] = useState(curriculum[0].id);
  const [entries, setEntries] = useState<Record<string, Entry>>({});
  const [loaded, setLoaded] = useState(false);
  const [vimEnabled, setVimEnabled] = useState(false);
  const [storageNotice, setStorageNotice] = useState("");
  const [result, setResult] = useState<RunResult | null>(null);
  const [running, setRunning] = useState(false);
  const sequence = useRef(0);
  const controller = useRef<AbortController | null>(null);
  const exercise = curriculum.find((item) => item.id === selected) || curriculum[0];
  const code = entries[exercise.id]?.code ?? exercise.code;
  const index = curriculum.indexOf(exercise);
  const completed = curriculum.filter((item) => entries[item.id]?.passed).length;

  useEffect(() => {
    try {
      const saved = JSON.parse(localStorage.getItem(storageKey) || "{}");
      const valid: Record<string, Entry> = {};
      for (const item of curriculum) {
        const entry = saved?.[item.id];
        if (entry && typeof entry.code === "string" && entry.code.length <= 40_000) valid[item.id] = { code: entry.code, passed: entry.passed === true };
      }
      setEntries(valid);
    } catch { setStorageNotice("Saved edits could not be read. You can still edit and run tests."); }
    try { setVimEnabled(localStorage.getItem(vimStorageKey) === "true"); }
    catch { /* The ordinary editor remains available without browser storage. */ }
    try {
      const id = decodeURIComponent(location.hash.slice(1));
      if (curriculum.some((item) => item.id === id)) setSelected(id);
    } catch { /* An invalid URL starts at the first exercise. */ }
    setLoaded(true);
    return () => controller.current?.abort();
  }, []);

  useEffect(() => {
    if (!loaded) return;
    try {
      localStorage.setItem(storageKey, JSON.stringify(entries));
      localStorage.setItem(vimStorageKey, String(vimEnabled));
    }
    catch { setStorageNotice("Browser storage is unavailable. Download your edited test to keep a copy."); }
  }, [entries, loaded, vimEnabled]);

  function stopRun() {
    controller.current?.abort(); sequence.current++; setRunning(false); setResult(null);
  }
  function select(id: string) {
    stopRun(); setSelected(id); history.replaceState(null, "", `#${id}`);
  }
  function edit(value: string) {
    stopRun(); setEntries((previous) => ({ ...previous, [exercise.id]: { code: value, passed: false } }));
  }
  async function run() {
    controller.current?.abort();
    const current = ++sequence.current;
    const abort = new AbortController(); controller.current = abort;
    setRunning(true); setResult(null);
    const timeout = setTimeout(() => abort.abort(), 25_000);
    try {
      const response = await fetch("/api/run", { method: "POST", headers: { "Content-Type": "application/json" }, body: JSON.stringify({ code }), signal: abort.signal });
      const output: RunResult = await response.json();
      if (current !== sequence.current) return;
      if (!["passed", "failed", "error"].includes(output.status) || typeof output.output !== "string") throw new Error("Invalid runner response");
      setResult(output);
      setEntries((previous) => ({ ...previous, [exercise.id]: { code, passed: output.status === "passed" } }));
    } catch {
      if (current === sequence.current) setResult({ status: "error", output: "The test runner did not respond. Try again, or download the test and run it locally.", passed: 0, failed: 0, skipped: 0 });
    } finally { clearTimeout(timeout); if (current === sequence.current) setRunning(false); }
  }
  function download() {
    const url = URL.createObjectURL(new Blob([code], { type: "text/plain" }));
    const anchor = document.createElement("a"); anchor.href = url; anchor.download = `${exercise.id}_test.go`; anchor.click();
    setTimeout(() => URL.revokeObjectURL(url), 1000);
  }

  return (
    <>
      <a className="skip-link" href="#exercise">Skip to exercise</a>
      <header className="header"><h1>Go Koans</h1><div className="header-links"><a href="/api/download">Download all tests</a><a href="https://github.com/ssemakov/go-koans">Source</a></div></header>
      <p className="instructions">Run the tests. Replace the TODO values or implement the missing code. Run again until the tests pass.</p>
      <div className="workspace">
        <nav className="curriculum" aria-label="Exercises">
          <p className="progress" aria-live="polite">{completed} / {curriculum.length} passed</p>
          {topics.map((topic) => (
            <details key={topic} open={topic === exercise.topic}>
              <summary>{topic} <span>{curriculum.filter((item) => item.topic === topic && entries[item.id]?.passed).length}/{curriculum.filter((item) => item.topic === topic).length}</span></summary>
              <ol>{curriculum.filter((item) => item.topic === topic).map((item) => (
                <li key={item.id}><button type="button" aria-current={item.id === exercise.id ? "step" : undefined} onClick={() => select(item.id)}><span className="exercise-number">{item.id.slice(0, 3)}</span> {item.title}{entries[item.id]?.passed && <span className="passed-label"> (passed)</span>}</button></li>
              ))}</ol>
            </details>
          ))}
        </nav>
        <main className="exercise" id="exercise">
          <p className="topic">{exercise.topic} · {index + 1} of {curriculum.length}</p><h2>{exercise.title}</h2><p>{exercise.instruction}</p>
          <div className="editor-heading">
            <label htmlFor="code">Go test code</label>
            <label className="vim-toggle"><input type="checkbox" checked={vimEnabled} onChange={(event) => setVimEnabled(event.target.checked)} /> Vim mode</label>
          </div>
          {vimEnabled ? <VimEditor key={exercise.id} value={code} onChange={edit} onRun={() => { if (!running) void run(); }} /> : <textarea id="code" aria-describedby="editor-help" spellCheck={false} autoCapitalize="off" autoCorrect="off" wrap="off" value={code} maxLength={40_000} onChange={(event) => edit(event.target.value)} onKeyDown={(event) => { if ((event.ctrlKey || event.metaKey) && event.key === "Enter") { event.preventDefault(); if (!running) void run(); } }} />}
          <div className="toolbar">
            <button className="run" type="button" onClick={() => void run()} disabled={running}>{running ? "Running…" : "Run tests"}</button>
            {running && <button type="button" onClick={stopRun}>Cancel</button>}
            <button type="button" onClick={download}>Download test</button>
            <button type="button" onClick={() => { if (window.confirm("Reset this exercise? Your edits and passed status will be removed.")) edit(exercise.code); }}>Reset</button>
            <a href={exercise.reference}>Go documentation</a>
          </div>
          <p className="help" id="editor-help">Ctrl/Cmd + Enter runs the tests. Edits and progress are saved in this browser.</p>
          {vimEnabled && <p className="help">Vim: i enters insert mode; Esc returns to normal mode. Edits save automatically.</p>}
          {storageNotice && <p className="notice" role="status">{storageNotice}</p>}
          <section className="output" aria-labelledby="output-title" aria-live="polite" aria-busy={running}>
            <h3 id="output-title">Test output{result && <span className={`result ${result.status}`}> — {result.status === "passed" ? "Passed" : result.status === "failed" ? "Failed" : "Runner error"}</span>}</h3>
            <pre>{running ? "Running tests…" : result?.output || "Run the tests to see the result."}</pre>
          </section>
          <div className="exercise-navigation"><button type="button" disabled={index === 0} onClick={() => select(curriculum[index - 1].id)}>Previous</button><button type="button" disabled={index === curriculum.length - 1} onClick={() => select(curriculum[index + 1].id)}>Next exercise</button></div>
          <details className="local-instructions"><summary>Run on your computer</summary><p>Install Go 1.23 or later. Download this test and run:</p><pre>go test -v {exercise.id}_test.go</pre><p>Or download all tests, extract the archive, and run the first exercise:</p><pre>cd go-koans{`\n`}go test -v ./{curriculum[0].id}</pre><p>Tests run on the Go Playground when you use the Run tests button.</p></details>
        </main>
      </div>
      <footer>References: <a href="https://github.com/edgecase/ruby_koans">Ruby Koans</a>, <a href="https://github.com/cdarwin/go-koans">cdarwin/go-koans</a>, <a href="https://github.com/VictorDenisov/go_koans">VictorDenisov/go_koans</a>, <a href="https://go.dev/wiki/">Go wiki</a>.</footer>
    </>
  );
}
