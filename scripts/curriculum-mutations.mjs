// Deliberately wrong implementations check that the learning fixtures reject
// plausible mistakes, rather than merely agreeing with their own solutions.
export const mutations = [
  { number: 34, name: "shared-clone", edits: [["result := make([]int, len(values))", "result := values"]] },
  { number: 38, name: "delete-shares-storage", edits: [["result := make([]int, 0, len(values)-1)", "result := values[:0]"]] },
  { number: 41, name: "zero-as-missing", edits: [["if ok {", "if ok && value != 0 {"]] },
  { number: 78, name: "wrong-method-set", edits: [["func (c *Counter) Increment()", "func (c Counter) Increment()"]] },
  { number: 80, name: "lost-error-identity", edits: [['"parse: %w"', '"parse: %v"']] },
  { number: 93, name: "unprotected-update", edits: [["\tc.mu.Lock()\n\tdefer c.mu.Unlock()\n\tc.n += delta", "\tc.n += delta"]] },
  { number: 96, name: "discarded-read-error", edits: [["data, err := io.ReadAll(reader)\n\treturn string(data), err", "data, _ := io.ReadAll(reader)\n\treturn string(data), nil"]] },
  { number: 97, name: "discarded-copy-error", edits: [["return io.Copy(writer, reader)", "count, _ := io.Copy(writer, reader)\n\treturn count, nil"]] },
  { number: 108, name: "shallow-value-copy", edits: [["result[key] = slices.Clone(values)", "result[key] = values"]] },
  { number: 111, name: "wrong-upper-bound", edits: [["if value > 10 {\n\t\treturn 10", "if value > 10 {\n\t\treturn 0"]] },
  { number: 117, name: "overflowing-side-sums", edits: [["a <= c-b || b <= a-c || c <= b-a", "a+b <= c || b+c <= a || c+a <= b"]] },
  { number: 119, name: "reused-roll-storage", edits: [["d.Values = make([]int, n)", "if cap(d.Values) >= n { d.Values = d.Values[:n] } else { d.Values = make([]int, n) }"]] },
  { number: 119, name: "ignored-generator", edits: [["d.Values[index] = d.Next()", "d.Values[index] = index%6 + 1"]] },
  {
    number: 120, name: "waits-for-input-closure",
    edits: [[
      `func squares(input <-chan int) <-chan int {
\toutput := make(chan int)
\tgo func() {
\t\tdefer close(output)
\t\tfor value := range input {
\t\t\toutput <- value * value
\t\t}
\t}()
\treturn output
}`,
      `func squares(input <-chan int) <-chan int {
\toutput := make(chan int, len(input))
\tfor value := range input { output <- value * value }
\tclose(output)
\treturn output
}`,
    ]],
  },
];
