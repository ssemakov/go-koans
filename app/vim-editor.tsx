"use client";

import { defaultKeymap, history, historyKeymap, indentLess, insertTab } from "@codemirror/commands";
import { go } from "@codemirror/lang-go";
import { bracketMatching, defaultHighlightStyle, indentOnInput, indentUnit, syntaxHighlighting } from "@codemirror/language";
import { Annotation, EditorState, Prec, StateEffect, StateField, Transaction } from "@codemirror/state";
import { drawSelection, EditorView, gutter, GutterMarker, highlightActiveLineGutter, highlightSpecialChars, keymap } from "@codemirror/view";
import { getCM, Vim, vim } from "@replit/codemirror-vim";
import { useEffect, useRef } from "react";

type Props = { value: string; onChange: (value: string) => void; onRun: () => void };
const externalChange = Annotation.define<boolean>();
const setRelativeNumbers = StateEffect.define<boolean>();
let defaultRelativeNumbers = true;
const relativeNumbers = StateField.define<boolean>({
  create: () => defaultRelativeNumbers,
  update(value, transaction) {
    for (const effect of transaction.effects) if (effect.is(setRelativeNumbers)) value = effect.value;
    return value;
  },
});

Vim.defineOption("relativenumber", true, "boolean", ["rnu"], (value, cm) => {
  if (!cm) {
    if (value !== undefined) defaultRelativeNumbers = value;
    return defaultRelativeNumbers;
  }
  if (value === undefined) return cm.cm6.state.field(relativeNumbers);
  cm.cm6.dispatch({ effects: setRelativeNumbers.of(value) });
  return value;
});

class LineNumber extends GutterMarker {
  constructor(readonly label: string) { super(); }
  eq(other: LineNumber) { return this.label === other.label; }
  toDOM() { return document.createTextNode(this.label); }
}

const vimLineNumbers = gutter({
  class: "cm-lineNumbers",
  lineMarker(view, line) {
    const number = view.state.doc.lineAt(line.from).number;
    const current = view.state.doc.lineAt(view.state.selection.main.head).number;
    return new LineNumber(String(view.state.field(relativeNumbers) && number !== current ? Math.abs(number - current) : number));
  },
  // A formatNumber callback alone would leave cached labels stale after motions.
  lineMarkerChange: (update) => update.selectionSet || update.docChanged || update.startState.field(relativeNumbers) !== update.state.field(relativeNumbers),
  initialSpacer: (view) => new LineNumber(String(view.state.doc.lines)),
  updateSpacer: (_spacer, update) => new LineNumber(String(update.state.doc.lines)),
});

function vimHistory(view: EditorView, redo: boolean) {
  const cm = getCM(view);
  if (!cm?.state.vim || cm.state.vim.insertMode) return false;
  Vim.handleKey(cm, "<Esc>", "user");
  // Native undo restores Vim operators' temporary selections as Visual mode.
  // Vim's history commands collapse those selections and restore a normal caret.
  Vim.handleKey(cm, redo ? "<C-r>" : "u", "user");
  return true;
}

export default function VimEditor({ value, onChange, onRun }: Props) {
  const container = useRef<HTMLDivElement>(null);
  const view = useRef<EditorView | null>(null);
  const initialValue = useRef(value);
  const callbacks = useRef({ onChange, onRun });
  callbacks.current = { onChange, onRun };

  useEffect(() => {
    if (!container.current) return;
    const editor = new EditorView({
      parent: container.current,
      state: EditorState.create({
        doc: initialValue.current,
        extensions: [
          // Raising a keymap's priority raises the shared keymap handler, including
          // default bindings that would steal Vim's Enter, Ctrl-D, and Ctrl-U.
          Prec.highest(EditorView.domEventHandlers({ keydown(event) {
            if (event.key !== "Enter" || !(event.ctrlKey || event.metaKey) || event.shiftKey || event.altKey || event.isComposing) return false;
            event.preventDefault();
            if (!event.repeat) callbacks.current.onRun();
            return true;
          } })),
          vim({ status: true }),
          history(), drawSelection(), relativeNumbers, vimLineNumbers, highlightActiveLineGutter(), highlightSpecialChars(),
          keymap.of([
            { key: "Mod-z", run: (view) => vimHistory(view, false) },
            { key: "Mod-Shift-z", run: (view) => vimHistory(view, true) },
            { key: "Mod-y", run: (view) => vimHistory(view, true) },
            {
              key: "Tab",
              run: (view) => { if (getCM(view)?.state.vim?.insertMode) insertTab(view); return true; },
              shift: (view) => { if (getCM(view)?.state.vim?.insertMode) indentLess(view); return true; },
            }, ...defaultKeymap, ...historyKeymap]),
          go(), indentOnInput(), bracketMatching(), syntaxHighlighting(defaultHighlightStyle),
          EditorState.tabSize.of(4), indentUnit.of("\t"), EditorState.allowMultipleSelections.of(true),
          EditorView.contentAttributes.of({ id: "code", "aria-label": "Go test code", "aria-describedby": "editor-help", spellcheck: "false", autocapitalize: "off", autocorrect: "off" }),
          EditorState.transactionFilter.of((transaction) => transaction.newDoc.length <= 40_000 ? transaction : []),
          EditorView.updateListener.of((update) => {
            if (update.docChanged && !update.transactions.some((transaction) => transaction.annotation(externalChange))) callbacks.current.onChange(update.state.doc.toString());
          }),
          EditorView.theme({
            "&": { fontSize: "13px", border: "1px solid #bfc4ca", borderRadius: "3px", backgroundColor: "#fafbfc" },
            "&.cm-focused": { outline: "2px solid #174ea6", outlineOffset: "3px" },
            ".cm-scroller": { fontFamily: "ui-monospace, SFMono-Regular, Consolas, monospace", lineHeight: "1.65", overflow: "auto", maxHeight: "600px" },
            ".cm-content": { minHeight: "380px", padding: "14px 0" },
            ".cm-gutters": { border: "none", backgroundColor: "#f2f4f6", color: "#5f6368" },
            ".cm-activeLineGutter": { color: "#174ea6", backgroundColor: "#e8f0fe", fontWeight: "600" },
            ".cm-panels": { backgroundColor: "#f2f4f6", fontSize: "12px" },
            ".cm-vim-panel": { padding: "4px 10px" },
          }),
        ],
      }),
    });
    view.current = editor;
    return () => { view.current = null; editor.destroy(); };
  }, []);

  useEffect(() => {
    const editor = view.current;
    if (editor && editor.state.doc.toString() !== value) {
      editor.dispatch({ changes: { from: 0, to: editor.state.doc.length, insert: value }, selection: { anchor: 0 }, annotations: [externalChange.of(true), Transaction.addToHistory.of(false)] });
    }
  }, [value]);

  return <div className="vim-editor" ref={container} />;
}
