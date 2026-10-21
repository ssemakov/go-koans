"use client";

import { defaultKeymap, history, historyKeymap, indentLess, insertTab } from "@codemirror/commands";
import { go } from "@codemirror/lang-go";
import { bracketMatching, HighlightStyle, indentOnInput, indentUnit, syntaxHighlighting } from "@codemirror/language";
import { Annotation, EditorState, Prec, StateEffect, StateField, Transaction } from "@codemirror/state";
import { drawSelection, EditorView, gutter, GutterMarker, highlightActiveLineGutter, highlightSpecialChars, keymap } from "@codemirror/view";
import { tags } from "@lezer/highlight";
import { getCM, Vim, vim } from "@replit/codemirror-vim";
import { useEffect, useRef } from "react";

type Props = { value: string; onChange: (value: string) => void; onRun: () => void };
const externalChange = Annotation.define<boolean>();
const highlightStyle = HighlightStyle.define([
  { tag: tags.keyword, color: "var(--syntax-keyword)" },
  { tag: [tags.atom, tags.bool, tags.number], color: "var(--syntax-literal)" },
  { tag: [tags.string, tags.character], color: "var(--syntax-string)" },
  { tag: [tags.typeName, tags.className, tags.namespace], color: "var(--syntax-type)" },
  { tag: [tags.definition(tags.variableName), tags.function(tags.variableName), tags.propertyName], color: "var(--syntax-name)" },
  { tag: [tags.special(tags.variableName), tags.standard(tags.variableName)], color: "var(--syntax-builtin)" },
  { tag: tags.comment, color: "var(--syntax-comment)" },
  { tag: tags.invalid, color: "var(--error)" },
]);
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
          keymap.of([{
            key: "Tab",
            run: (view) => { if (getCM(view)?.state.vim?.insertMode) insertTab(view); return true; },
            shift: (view) => { if (getCM(view)?.state.vim?.insertMode) indentLess(view); return true; },
          }, ...defaultKeymap, ...historyKeymap]),
          go(), indentOnInput(), bracketMatching(), syntaxHighlighting(highlightStyle),
          EditorState.tabSize.of(4), indentUnit.of("\t"), EditorState.allowMultipleSelections.of(true),
          EditorView.contentAttributes.of({ id: "code", "aria-label": "Go test code", "aria-describedby": "editor-help", spellcheck: "false", autocapitalize: "off", autocorrect: "off" }),
          EditorState.transactionFilter.of((transaction) => transaction.newDoc.length <= 40_000 ? transaction : []),
          EditorView.updateListener.of((update) => {
            if (update.docChanged && !update.transactions.some((transaction) => transaction.annotation(externalChange))) callbacks.current.onChange(update.state.doc.toString());
          }),
          EditorView.theme({
            "&": { fontSize: "13px", border: "1px solid var(--border)", borderRadius: "3px", backgroundColor: "var(--surface)", color: "var(--foreground)" },
            "&.cm-focused": { outline: "2px solid var(--accent)", outlineOffset: "3px" },
            ".cm-scroller": { fontFamily: "ui-monospace, SFMono-Regular, Consolas, monospace", lineHeight: "1.65", overflow: "auto", maxHeight: "600px" },
            ".cm-content": { minHeight: "380px", padding: "14px 0", caretColor: "var(--foreground)" },
            ".cm-gutters": { border: "none", backgroundColor: "var(--surface-raised)", color: "var(--muted)" },
            ".cm-activeLineGutter": { color: "var(--accent)", backgroundColor: "var(--selected)", fontWeight: "600" },
            ".cm-panels": { backgroundColor: "var(--surface-raised)", color: "var(--foreground)", borderColor: "var(--divider)", fontSize: "12px" },
            ".cm-vim-panel": { padding: "4px 10px" },
            ".cm-vim-panel input": { color: "var(--foreground)", caretColor: "var(--foreground)" },
            ".cm-cursor, .cm-dropCursor": { borderLeftColor: "var(--foreground)" },
            ".cm-selectionBackground, &.cm-focused > .cm-scroller > .cm-selectionLayer .cm-selectionBackground": { backgroundColor: "var(--selection)" },
            ".cm-searchMatch": { backgroundColor: "var(--search-match)" },
            "&.cm-focused .cm-matchingBracket": { backgroundColor: "var(--selected)" },
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
