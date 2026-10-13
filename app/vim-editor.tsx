"use client";

import { defaultKeymap, history, historyKeymap } from "@codemirror/commands";
import { defaultHighlightStyle, StreamLanguage, syntaxHighlighting } from "@codemirror/language";
import { go } from "@codemirror/legacy-modes/mode/go";
import { Annotation, EditorState, Prec, Transaction } from "@codemirror/state";
import { drawSelection, EditorView, highlightSpecialChars, keymap, lineNumbers } from "@codemirror/view";
import { vim } from "@replit/codemirror-vim";
import { useEffect, useRef } from "react";

type Props = { value: string; onChange: (value: string) => void; onRun: () => void };
const externalChange = Annotation.define<boolean>();

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
          Prec.highest(keymap.of([{ key: "Mod-Enter", run: () => { callbacks.current.onRun(); return true; } }])),
          vim({ status: true }),
          history(), drawSelection(), lineNumbers(), highlightSpecialChars(),
          keymap.of([...defaultKeymap, ...historyKeymap]),
          StreamLanguage.define(go), syntaxHighlighting(defaultHighlightStyle),
          EditorState.tabSize.of(4),
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
