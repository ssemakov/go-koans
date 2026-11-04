// Work around codemirror-vim-core 0.1.0 selecting a multiline enclosing
// block before looking for a brace text object later on the current line.
import { readFileSync, writeFileSync } from "node:fs";

const file = new URL("../node_modules/@replit/codemirror-vim-core/vim.js", import.meta.url);
const source = readFileSync(file, "utf8");
const before = "    var start = startBracket.pos; \n    var end = endBracket.pos;";
const after = `    // Prefer the next brace pair on this line over a multiline outer block.
    if (openSym === '{' && startBracket.pos.line < head.line) {
      var nextBrace = /[{}]/.exec(cm.getLine(head.line).slice(head.ch));
      if (nextBrace && nextBrace[0] === '{') {
        return selectCompanionObject(cm,
          new Pos(head.line, head.ch + nextBrace.index), symb, inclusive);
      }
    }

${before}`;
if (!source.includes(after)) {
  if (!source.includes(before)) {
    throw new Error("Vim text-object patch no longer applies; review the dependency implementation.");
  }
  writeFileSync(file, source.replace(before, after));
}
