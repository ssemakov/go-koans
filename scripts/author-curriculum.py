"""Author the starter files and reference fixtures. Requires gofmt on PATH.

Go tests use only testing.T and standard-library comparisons. The helpers in
this file generate source; none are shipped as assertion helpers in Go.
"""
import json
import shutil
import subprocess
import textwrap
from collections import Counter
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
METADATA = json.loads((ROOT / "lib/curriculum.json").read_text())
LESSONS = {}


def add(number, instruction, body, edits, *, imports=(), declarations="",
        title=None, kind="implement", revisits=()):
    item = dict(METADATA[number - 1])
    assert item["id"].startswith(f"{number:03}-")
    if title:
        item["title"] = title
    names = sorted(set(("testing", *imports)))
    comments = f'// {item["title"]}\n' + "".join(f"// {line}\n" for line in textwrap.wrap(instruction, width=88))
    source = "package main\n\n" + comments + '\nimport (\n' + "".join(f'\t"{name}"\n' for name in names) + ")\n\n"
    source += declarations.strip() + "\n\n" if declarations else ""
    source += "func TestKoan(t *testing.T) {\n" + body.strip() + "\n}\n"
    solution = source
    for old, new in edits:
        assert solution.count(old) == 1, (number, old, solution.count(old))
        solution = solution.replace(old, new)
    assert source != solution
    item.update(instruction=instruction, code=source, kind=kind,
                revisits=[METADATA[n - 1]["id"] for n in revisits])
    LESSONS[number] = (item, solution)


def observe_steps(number, instruction, steps, *, imports=(), declarations="",
                  title=None, revisits=()):
    body = ""
    edits = []
    for setup, checks in steps:
        body += setup + "\n"
        for name, expression, blank, answer in checks:
            line = f"want{name} := {blank} // TODO: replace this value."
            body += f"""
{line}
if got := {expression}; got != want{name} {{
    t.Errorf("{name}: got %#v; want %#v", got, want{name})
}}
"""
            edits.append((line, f"want{name} := {answer}"))
    add(number, instruction, body, edits, imports=imports, declarations=declarations,
        title=title, kind="observe", revisits=revisits)


def observe(number, instruction, setup, checks, *, imports=(), declarations="",
            title=None, revisits=()):
    observe_steps(number, instruction, [(setup, checks)], imports=imports,
                  declarations=declarations, title=title, revisits=revisits)


add(1, "Replace the TODO value with true. Run the test to see how a failed check is reported.", """
want := false // TODO: replace this value.
got := true
if got != want {
    t.Fatalf("got %v; want %v", got, want)
}
""", [("want := false // TODO: replace this value.", "want := true")], kind="complete")

observe(2, "Fill both expected values. t.Errorf reports a failed check and continues, so one run can report both failures.", "", [
    ("Sum", "1 + 1", "-999", "2"),
    ("Different", "2 == 3", "true", "false"),
])

observe(3, "Fill the expected values. Compare an omitted integer field with one initialized explicitly.", """
var count int
record := struct{ Count int }{}
initialized := struct{ Count int }{Count: 7}
""", [("Variable", "count", "-999", "0"), ("OmittedField", "record.Count", "-999", "0"),
      ("InitializedField", "initialized.Count", "-999", "7")])

observe(4, "Fill the expected strings. An empty input and an omitted field do not supply a default name automatically.", """
var name string
record := struct{ Name string }{}
initialized := struct{ Name string }{Name: "Ada"}
""", [("Variable", "name", '"TODO"', '""'), ("OmittedField", "record.Name", '"TODO"', '""'),
      ("InitializedField", "initialized.Name", '"TODO"', '"Ada"')], revisits=(3,))

add(5, "Remove the explicit initializer so enabled receives Go's default boolean value. Do not replace true with false. Keep the assertion.", """
var enabled bool = true

if enabled {
    t.Error("an uninitialized boolean should be false")
}
""", [("var enabled bool = true", "var enabled bool")],
    kind="repair", revisits=(3, 4))

add(6, "Repair the assignment inside the block. Update the existing total instead of declaring a shadowing variable. Keep the checks.", """
for _, delta := range []int{1, 4, -2} {
    total := 10
    {
        total := total + delta // TODO: update the outer variable.
        _ = total
    }
    if total != 10 + delta { t.Errorf("delta=%d: total=%d; want %d", delta, total, 10+delta) }
}
""", [("total := total + delta // TODO: update the outer variable.", "total = total + delta")], kind="repair")

add(7, "Repair the swap using one simultaneous assignment. Both right-hand values must come from the original pair.", """
for _, pair := range [][2]int{{2,5}, {0,7}, {-3,4}, {8,8}} {
    a, b := pair[0], pair[1]
    // TODO: replace the two assignments with a simultaneous swap.
    a = b
    b = a
    if a != pair[1] || b != pair[0] {
        t.Errorf("swap(%v) = [%d %d]; want [%d %d]", pair, a, b, pair[1], pair[0])
    }
}
""", [("// TODO: replace the two assignments with a simultaneous swap.\n    a = b\n    b = a", "a, b = b, a")],
    kind="repair", revisits=(6,))

add(8, "Fill the conversion and declaration. An untyped constant can initialize several compatible types; a typed variable needs an explicit conversion.", """
const unit = 3
var small int8 = unit
var wide int64 = unit
typed := int8(7)
converted := int64(0) // TODO: convert typed to int64.
_ = typed
if small != 3 || wide != 3 { t.Error("constant initialization changed the value") }
if converted != 7 { t.Errorf("converted=%d; want 7", converted) }
const third = 0 // TODO: use an exact constant expression: 1.0 / 3.
const restored = third * 3
if restored != 1 { t.Errorf("restored=%v; want 1", restored) }
""", [("converted := int64(0) // TODO: convert typed to int64.", "converted := int64(typed)"),
      ("const third = 0 // TODO: use an exact constant expression: 1.0 / 3.", "const third = 1.0 / 3")],
    kind="complete")

observe(9, "Fill the values for each constant block. Check a skipped entry, a reset, and bit flags generated with shifts.", """
const (first = iota; _; third)
const (reset = iota; next)
const (read = 1 << iota; write; execute)
""", [("First", "first", "-999", "0"), ("Third", "third", "-999", "2"),
      ("Reset", "reset", "-999", "0"), ("Next", "next", "-999", "1"),
      ("CombinedFlags", "read | execute", "-999", "5")], revisits=(8,))

observe(10, "Fill the quotient and remainder values for both signs, then compare floating-point division.", """
positive, negative, divisor := 7, -7, 2
floating := float64(positive) / float64(divisor)
""", [("Positive", "positive / divisor", "-999", "3"), ("Negative", "negative / divisor", "-999", "-3"),
      ("NegativeRemainder", "negative % divisor", "-999", "-1"), ("Floating", "floating", "-999.0", "3.5")])

add(11, "Implement wrapIndex for a positive size. Return an index from 0 through size-1, including for negative positions.", """
cases := []struct{ position, size, want int }{
    {0,5,0}, {7,5,2}, {5,5,0}, {-1,5,4}, {-6,5,4}, {12,1,0},
}
for _, tc := range cases {
    if got := wrapIndex(tc.position, tc.size); got != tc.want {
        t.Errorf("wrapIndex(%d,%d)=%d; want %d", tc.position, tc.size, got, tc.want)
    }
}
""", [("return 0 // TODO: implement.", "remainder := position % size\nif remainder < 0 { remainder += size }\nreturn remainder")],
    declarations="func wrapIndex(position, size int) int { return 0 // TODO: implement.\n}", revisits=(10,))

add(12, "Implement toggle using XOR. Change only the bits selected by mask; applying the same mask twice must restore the value.", """
cases := []struct{ value, mask, want uint8 }{
    {5,2,7}, {7,2,5}, {5,1,4}, {0,8,8}, {255,0,255},
}
for _, tc := range cases {
    got := toggle(tc.value,tc.mask)
    if got != tc.want { t.Errorf("toggle(%08b,%08b)=%08b; want %08b", tc.value,tc.mask,got,tc.want) }
    if restored := toggle(got,tc.mask); restored != tc.value { t.Errorf("toggling twice gave %d; want %d", restored,tc.value) }
}
""", [("return value // TODO: implement.", "return value ^ mask")],
    declarations="func toggle(value, mask uint8) uint8 { return value // TODO: implement.\n}", revisits=(9,))

add(13, "Complete the guarded expression. Return true only when the slice has a first element and that element is positive. Empty input must not panic.", """
cases := []struct{ values []int; want bool }{
    {nil,false}, {[]int{},false}, {[]int{0},false}, {[]int{-1,2},false}, {[]int{3},true},
}
for _, tc := range cases {
    got := false // TODO: use a short-circuit expression.
    if got != tc.want { t.Errorf("values=%v: got %v; want %v", tc.values,got,tc.want) }
}
""", [("got := false // TODO: use a short-circuit expression.", "got := len(tc.values) > 0 && tc.values[0] > 0")],
    kind="complete", revisits=(5,))

observe(14, "Fill the results for both signs. Compare conversion to int with math.Floor; they differ for negative fractions.", """
positive, negative := 3.75, -3.75
""", [("PositiveInt", "int(positive)", "-999", "3"), ("NegativeInt", "int(negative)", "-999", "-3"),
      ("PositiveFloor", "int(math.Floor(positive))", "-999", "3"), ("NegativeFloor", "int(math.Floor(negative))", "-999", "-4")],
    imports=("math",), revisits=(10,))

add(15, "Implement addBonus. Score is a defined type, so convert between Score and int explicitly when combining them.", """
cases := []struct{ score Score; bonus int; want Score }{{4,3,7}, {0,5,5}, {10,-2,8}}
for _, tc := range cases {
    if got := addBonus(tc.score,tc.bonus); got != tc.want { t.Errorf("addBonus(%d,%d)=%d; want %d",tc.score,tc.bonus,got,tc.want) }
}
""", [("return score // TODO: combine the values with an explicit conversion.", "return score + Score(bonus)")],
    declarations="type Score int\nfunc addBonus(score Score, bonus int) Score { return score // TODO: combine the values with an explicit conversion.\n}",
    revisits=(14,))

add(16, "Complete the aliases. Count should be the same type as int, not a distinct defined type. Bytes should be the same type as []byte.", """
var plain int = 4
var count Count = Count(plain)
count += 3
plain = int(count)
if got := reflect.TypeOf(count).Name(); got != "int" { t.Errorf("Count type name=%q; want int",got) }
if plain != 7 { t.Errorf("plain=%d; want 7",plain) }
var bytes Bytes = []byte("Go")
var ordinary []byte = bytes
ordinary[0] = 'g'
if reflect.TypeOf(bytes) != reflect.TypeOf(ordinary) { t.Error("Bytes must be the same type as []byte") }
if string(bytes) != "go" { t.Errorf("bytes=%q; want go",bytes) }
""", [("type Count int // TODO: make Count an alias of int.", "type Count = int"),
      ("type Bytes []byte // TODO: make Bytes an alias of []byte.", "type Bytes = []byte")],
    declarations="type Count int // TODO: make Count an alias of int.\ntype Bytes []byte // TODO: make Bytes an alias of []byte.",
    imports=("reflect",), kind="complete", revisits=(15,))

add(17, "Implement joinName using concatenation. Add one space between nonempty parts; an empty part must not add a leading or trailing space.", """
cases := []struct{ first,last,want string }{
    {"Ada","Lovelace","Ada Lovelace"}, {"Go","","Go"}, {"","Gopher","Gopher"}, {"","",""},
}
for _, tc := range cases {
    first,last := tc.first,tc.last
    if got := joinName(first,last); got != tc.want { t.Errorf("joinName(%q,%q)=%q; want %q",first,last,got,tc.want) }
    if first != tc.first || last != tc.last { t.Error("the inputs changed") }
}
""", [("return first // TODO: implement.", 'if first == "" { return last }\nif last == "" { return first }\nreturn first + " " + last')],
    declarations="func joinName(first,last string) string { return first // TODO: implement.\n}")

observe(18, "Fill each substring. The upper bound is exclusive; omitted bounds and equal bounds produce different results.", 'text := "gopher"', [
    ("Middle", "text[2:5]", '"TODO"', '"phe"'), ("Prefix", "text[:2]", '"TODO"', '"go"'),
    ("Suffix", "text[2:]", '"TODO"', '"pher"'), ("Empty", "text[2:2]", '"TODO"', '""'),
    ("Whole", "text[:]", '"TODO"', '"gopher"'),
])

observe(19, "Fill the byte lengths and indexed byte value. Go string indexing reads a byte, even for UTF-8 text.", """
ascii := "Go"
accent := "é"
mixed := "é世"
""", [("ASCIIBytes", "len(ascii)", "-999", "2"), ("AccentBytes", "len(accent)", "-999", "2"),
      ("MixedBytes", "len(mixed)", "-999", "5"), ("FirstByte", "int(accent[0])", "-999", "195")],
    revisits=(18,))

observe(20, "Fill the rune counts. A rune is a Unicode code point; a base letter plus a combining mark can display as one character but contain two runes.", """
accent := "é"
decomposed := "e\\u0301"
mixed := "é世"
""", [("AccentRunes", "utf8.RuneCountInString(accent)", "-999", "1"),
      ("DecomposedRunes", "utf8.RuneCountInString(decomposed)", "-999", "2"),
      ("DecomposedBytes", "len(decomposed)", "-999", "3"),
      ("MixedRunes", "utf8.RuneCountInString(mixed)", "-999", "2")],
    imports=("unicode/utf8",), title="Count Unicode code points", revisits=(19,))

add(21, "Complete the two conversions. Keep the original string independent of later byte edits, and make an editable byte copy of it.", """
data := []byte("Go")
original := "" // TODO: convert data to a string.
data[0] = 'N'
editable := []byte{} // TODO: convert original to bytes.
if original != "Go" { t.Fatalf("original=%q; want Go after mutating data",original) }
if len(editable) != 2 { t.Fatalf("editable=%v; want two bytes",editable) }
editable[1] = '!'
if string(data) != "No" { t.Errorf("data=%q; want No",data) }
if string(editable) != "G!" { t.Errorf("editable=%q; want G!",editable) }
if original != "Go" { t.Errorf("editing a byte copy changed original=%q",original) }
""", [('original := "" // TODO: convert data to a string.', "original := string(data)"),
      ("editable := []byte{} // TODO: convert original to bytes.", "editable := []byte(original)")],
    kind="complete", revisits=(19,))

add(22, "Implement splitFields with strings.Split. The separator is an exact string, not a set of delimiter characters. Preserve empty fields.", """
cases := []struct{ text,sep string; want []string }{
    {"a,b,c",",",[]string{"a","b","c"}},
    {"a,,b,",",",[]string{"a","","b",""}},
    {"a,|;b;c",",|;",[]string{"a","b;c"}},
    {"",",",[]string{""}},
    {"abc",",",[]string{"abc"}},
}
for _, tc := range cases {
    if got := splitFields(tc.text,tc.sep); !slices.Equal(got,tc.want) { t.Errorf("splitFields(%q,%q)=%q; want %q",tc.text,tc.sep,got,tc.want) }
}
""", [("return nil // TODO: use strings.Split.", "return strings.Split(text,sep)")],
    imports=("slices","strings"), declarations="var _ = strings.Split\nfunc splitFields(text,sep string) []string { return nil // TODO: use strings.Split.\n}")

add(23, "Implement formatEntry with fmt.Sprintf. Quote the name with %q and print the count as a decimal integer.", """
cases := []struct{ name string; count int; want string }{
    {"Go",7,"\\\"Go\\\":7"}, {"two words",0,"\\\"two words\\\":0"}, {"a\\\"b",-2,"\\\"a\\\\\\\"b\\\":-2"},
}
for _, tc := range cases {
    if got := formatEntry(tc.name,tc.count); got != tc.want { t.Errorf("formatEntry(%q,%d)=%q; want %q",tc.name,tc.count,got,tc.want) }
}
""", [('return "" // TODO: format the name and count.', 'return fmt.Sprintf("%q:%d",name,count)')],
    imports=("fmt",), declarations='var _ = fmt.Sprintf\nfunc formatEntry(name string,count int) string { return "" // TODO: format the name and count.\n}',
    revisits=(17,))

add(24, "Fill the expected indices and runes for each string. range yields byte offsets and decoded runes, not one index per byte.", """
cases := []struct{ text string; indices []int; runes []rune }{
    {"Go",[]int{-999},[]rune{'?'}}, // TODO: indices and runes.
    {"é世",[]int{-999},[]rune{'?'}}, // TODO: indices and runes.
    {"e\\u0301",[]int{-999},[]rune{'?'}}, // TODO: indices and runes.
}
for _, tc := range cases {
    var indices []int
    var runes []rune
    for index,value := range tc.text { indices=append(indices,index); runes=append(runes,value) }
    if !slices.Equal(indices,tc.indices) { t.Errorf("%q: indices=%v; want %v",tc.text,indices,tc.indices) }
    if !slices.Equal(runes,tc.runes) { t.Errorf("%q: runes=%U; want %U",tc.text,runes,tc.runes) }
}
""", [('{"Go",[]int{-999},[]rune{\'?\'}}, // TODO: indices and runes.', '{"Go",[]int{0,1},[]rune{\'G\',\'o\'}},'),
      ('{"é世",[]int{-999},[]rune{\'?\'}}, // TODO: indices and runes.', '{"é世",[]int{0,2},[]rune{\'é\',\'世\'}},'),
      ('{"e\\u0301",[]int{-999},[]rune{\'?\'}}, // TODO: indices and runes.', '{"e\\u0301",[]int{0,1},[]rune{\'e\',\'\\u0301\'}},')],
    imports=("slices",), kind="observe", revisits=(19,20))

add(25, "Complete the array declaration. It must hold exactly three readings, including an omitted zero-valued reading, and match the function's [3]int parameter.", """
readings := [3]int{} // TODO: initialize the first two readings to 4 and 7.
if got := totalReadings(readings); got != 11 { t.Errorf("total=%d; want 11",got) }
if readings != [3]int{4,7,0} { t.Errorf("readings=%v; want [4 7 0]",readings) }
if len(readings) != 3 || cap(readings) != 3 { t.Error("an array has a fixed length and capacity") }
""", [("readings := [3]int{} // TODO: initialize the first two readings to 4 and 7.", "readings := [3]int{4,7}")],
    declarations="func totalReadings(values [3]int) int { return values[0]+values[1]+values[2] }",
    kind="complete", revisits=(3,))

observe(26, "Fill the complete expected arrays. Omitted and explicitly indexed elements leave other positions at their zero values.", """
partial := [4]int{8,3}
indexed := [4]int{2:9}
""", [("Partial", "partial", "[4]int{-1,-1,-1,-1}", "[4]int{8,3,0,0}"),
      ("Indexed", "indexed", "[4]int{-1,-1,-1,-1}", "[4]int{0,0,9,0}")], revisits=(3,25))

observe(27, "Fill both expected arrays after the copy is modified. Check the whole values, not just one element.", """
original := [3]int{1,2,3}
copy := original
copy[0] = 9
""", [("Original", "original", "[3]int{-1,-1,-1}", "[3]int{1,2,3}"),
      ("Copy", "copy", "[3]int{-1,-1,-1}", "[3]int{9,2,3}")])

add(28, "Complete the array literal using [...] so its length is inferred. Check that its type is [3]string and that it contains all three names.", """
names := [...]string{"?"} // TODO: include Ada, Grace, and Edsger.
if got := fmt.Sprintf("%T",names); got != "[3]string" { t.Errorf("type=%s; want [3]string",got) }
if got := fmt.Sprint(names); got != "[Ada Grace Edsger]" { t.Errorf("names=%s; want [Ada Grace Edsger]",got) }
""", [('names := [...]string{"?"} // TODO: include Ada, Grace, and Edsger.', 'names := [...]string{"Ada","Grace","Edsger"}')],
    imports=("fmt",), kind="complete", revisits=(23,25,26))

observe(29, "Track the two slice headers. Appending within capacity changes the returned slice's length, not the old header's length.", """
values := make([]int,2,4)
old := values
values = append(values,7)
""", [("OldLength", "len(old)", "-999", "2"), ("NewLength", "len(values)", "-999", "3"),
      ("OldCapacity", "cap(old)", "-999", "4"), ("NewCapacity", "cap(values)", "-999", "4"),
      ("AppendedValue", "old[:3][2]", "-999", "7")])

observe(30, "Fill the length, capacity, and original array after a slice is modified. Capacity starts at the slice's first element and extends to the end of the array.", """
array := [5]int{1,2,3,4,5}
part := array[1:3]
part[0] = 9
""", [("Length", "len(part)", "-999", "2"), ("Capacity", "cap(part)", "-999", "4"),
      ("Array", "array", "[5]int{-1,-1,-1,-1,-1}", "[5]int{1,9,3,4,5}")], revisits=(27,))

add(31, "Fill the complete expected contents after modifying the slice view. Compare this with array assignment in exercise 027.", """
values := []int{1,2,3,4}
view := values[1:3]
view[0] = 9
wantValues := []int{-999} // TODO: all values.
wantView := []int{-999} // TODO: all values in the view.
if !slices.Equal(values,wantValues) { t.Errorf("values=%v; want %v",values,wantValues) }
if !slices.Equal(view,wantView) { t.Errorf("view=%v; want %v",view,wantView) }
""", [("wantValues := []int{-999} // TODO: all values.", "wantValues := []int{1,9,3,4}"),
      ("wantView := []int{-999} // TODO: all values in the view.", "wantView := []int{9,3}")],
    imports=("slices",), kind="observe", revisits=(27,30))

add(32, "Fill the complete expected slices after both appends. Keep the setup and append operations unchanged. Compare capacities 4 and 2: can first change without being reassigned?", """
cases := []struct {
    capacity int
    wantBase []int
    wantFirst []int
    wantSecond []int
}{
    {4, []int{-999}, []int{-999}, []int{-999}}, // TODO: base, first, second.
    {2, []int{-999}, []int{-999}, []int{-999}}, // TODO: base, first, second.
}
for _, tc := range cases {
    base := make([]int,2,tc.capacity)
    base[0] = 10
    base[1] = 20

    first := append(base,30)
    second := append(base,40)

    if !slices.Equal(base,tc.wantBase) {
        t.Errorf("capacity=%d: base=%v; want %v",tc.capacity,base,tc.wantBase)
    }
    if !slices.Equal(first,tc.wantFirst) {
        t.Errorf("capacity=%d: first=%v; want %v",tc.capacity,first,tc.wantFirst)
    }
    if !slices.Equal(second,tc.wantSecond) {
        t.Errorf("capacity=%d: second=%v; want %v",tc.capacity,second,tc.wantSecond)
    }
}
""", [("{4, []int{-999}, []int{-999}, []int{-999}}, // TODO: base, first, second.",
       "{4, []int{10,20}, []int{10,20,40}, []int{10,20,40}},"),
      ("{2, []int{-999}, []int{-999}, []int{-999}}, // TODO: base, first, second.",
       "{2, []int{10,20}, []int{10,20,30}, []int{10,20,40}},")],
    imports=("slices",), title="Append and shared storage", kind="observe", revisits=(29,30,31))

observe_steps(33, "Compare append with spare capacity and append beyond capacity. Fill the values observed through each alias. Do not assume an exact growth factor.", [("""
spare := make([]int,1,3)
spare[0] = 1
spareAlias := spare[:3]
spare = append(spare,7)
""", [("ReusedAppended", "spareAlias[1]", "-999", "7")]),
("""
spare[0] = 9
""", [("ReusedFirst", "spareAlias[0]", "-999", "9")]),
("""
full := []int{1}
fullAlias := full
grown := append(full,7)
grown[0] = 9
""", [("AllocatedOld", "fullAlias[0]", "-999", "1"), ("AllocatedNew", "grown[0]", "-999", "9"),
      ("AllocatedLength", "len(grown)", "-999", "2")])], revisits=(29,31))

add(34, "Implement cloneInts. Preserve every element and give the result independent storage. Mutating either slice must leave the other unchanged; nil or empty input may return nil.", """
for _, input := range [][]int{nil, {}, {7}, {1,2,3}} {
    original := slices.Clone(input)
    cloned := cloneInts(input)
    if !slices.Equal(cloned,original) { t.Errorf("cloneInts(%v)=%v",original,cloned); continue }
    if len(input)>0 {
        cloned[0]=99
        if !slices.Equal(input,original) { t.Error("editing clone changed input") }
        input[len(input)-1]=-7
        if cloned[len(cloned)-1] == -7 { t.Error("editing input changed clone") }
    }
}
""", [("return values // TODO: allocate and copy.", "result := make([]int,len(values))\ncopy(result,values)\nreturn result")],
    imports=("slices",), declarations="func cloneInts(values []int) []int { return values // TODO: allocate and copy.\n}",
    revisits=(27,31))

add(35, "Predict each copy's return value and the resulting destination. The first two destinations are shorter and longer than source. For the overlapping copy, predict the entire sharedValues slice.", """
source := []int{1,2,3}

shortDestination := []int{9,9}
wantShortCopyCount := -999 // TODO: number of elements copied.
if copiedCount := copy(shortDestination,source); copiedCount != wantShortCopyCount {
    t.Errorf("copy to shortDestination returned %d; want %d",copiedCount,wantShortCopyCount)
}
wantShortDestination := []int{-999} // TODO: complete destination after copy.
if !slices.Equal(shortDestination,wantShortDestination) {
    t.Errorf("shortDestination=%v; want %v",shortDestination,wantShortDestination)
}

longDestination := []int{9,9,9,9}
wantLongCopyCount := -999 // TODO: number of elements copied.
if copiedCount := copy(longDestination,source); copiedCount != wantLongCopyCount {
    t.Errorf("copy to longDestination returned %d; want %d",copiedCount,wantLongCopyCount)
}
wantLongDestination := []int{-999} // TODO: complete destination after copy.
if !slices.Equal(longDestination,wantLongDestination) {
    t.Errorf("longDestination=%v; want %v",longDestination,wantLongDestination)
}

sharedValues := []int{1,2,3,4}
overlappingSource := sharedValues[:3]
overlappingDestination := sharedValues[1:]
wantOverlapCopyCount := -999 // TODO: number of elements copied.
if copiedCount := copy(overlappingDestination,overlappingSource); copiedCount != wantOverlapCopyCount {
    t.Errorf("copy to overlappingDestination returned %d; want %d",copiedCount,wantOverlapCopyCount)
}
wantSharedValues := []int{-999} // TODO: entire shared slice after copy.
if !slices.Equal(sharedValues,wantSharedValues) {
    t.Errorf("sharedValues=%v; want %v",sharedValues,wantSharedValues)
}
""", [("wantShortCopyCount := -999 // TODO: number of elements copied.", "wantShortCopyCount := 2"),
      ("wantShortDestination := []int{-999} // TODO: complete destination after copy.", "wantShortDestination := []int{1,2}"),
      ("wantLongCopyCount := -999 // TODO: number of elements copied.", "wantLongCopyCount := 3"),
      ("wantLongDestination := []int{-999} // TODO: complete destination after copy.", "wantLongDestination := []int{1,2,3,9}"),
      ("wantOverlapCopyCount := -999 // TODO: number of elements copied.", "wantOverlapCopyCount := 3"),
      ("wantSharedValues := []int{-999} // TODO: entire shared slice after copy.", "wantSharedValues := []int{1,1,2,3}")],
    imports=("slices",), kind="observe", revisits=(34,))

add(36, "Complete the appends to the nil slice and the non-nil empty slice. Both should contain 7 afterward. Keep the checks that show their initial difference and their contents after append.", """
var absent []int
if absent != nil || len(absent) != 0 { t.Fatal("a nil slice starts nil with length zero") }
// TODO: append 7 to absent and store the returned slice in absent.
if absent == nil || !slices.Equal(absent,[]int{7}) { t.Errorf("after append, absent=%v; want a non-nil slice containing [7]",absent) }

empty := []int{}
if empty == nil || len(empty) != 0 { t.Fatal("an empty slice literal starts non-nil with length zero") }
// TODO: append 7 to empty and store the returned slice in empty.
if empty == nil || !slices.Equal(empty,[]int{7}) { t.Errorf("after append, empty=%v; want a non-nil slice containing [7]",empty) }
""", [("// TODO: append 7 to absent and store the returned slice in absent.", "absent = append(absent,7)"),
      ("// TODO: append 7 to empty and store the returned slice in empty.", "empty = append(empty,7)")],
    imports=("slices",), kind="complete", revisits=(3,29))

observe_steps(37, "Compare an unrestricted view with a full slice expression that caps capacity. Check capacities before append, then predict the slice length and values after append and mutation.", [("""
array := [4]int{1,2,3,4}
unrestricted := array[1:2]
limited := array[1:2:2]
""", [("UnrestrictedCapacity", "cap(unrestricted)", "-999", "3"),
      ("LimitedCapacity", "cap(limited)", "-999", "1")]),
("""
unrestricted = append(unrestricted,8)
limited = append(limited,9)
""", [("LimitedLength", "len(limited)", "-999", "2")]),
("""
limited[0] = 7
""", [("OriginalArray", "array", "[4]int{-1,-1,-1,-1}", "[4]int{1,2,8,4}"),
      ("LimitedFirst", "limited[0]", "-999", "7"),
      ("LimitedSecond", "limited[1]", "-999", "9")])], revisits=(30,33))

add(38, "Implement deleteAt for a valid index. Keep the remaining elements in order and leave the input unchanged. A result with zero elements may be nil.", """
cases := []struct{ values []int; index int; want []int }{
    {[]int{1,2,3},0,[]int{2,3}}, {[]int{1,2,3},1,[]int{1,3}},
    {[]int{1,2,3},2,[]int{1,2}}, {[]int{7},0,nil},
}
for _, tc := range cases {
    before := slices.Clone(tc.values)
    got := deleteAt(tc.values,tc.index)
    if !slices.Equal(got,tc.want) { t.Errorf("deleteAt(%v,%d)=%v; want %v",before,tc.index,got,tc.want) }
    if len(got)>0 { got[0]=99 }
    if !slices.Equal(tc.values,before) { t.Error("result shares mutable storage with input") }
}
""", [("return values // TODO: construct an independent result.", "result := make([]int,0,len(values)-1)\nresult=append(result,values[:index]...)\nresult=append(result,values[index+1:]...)\nreturn result")],
    imports=("slices",), declarations="func deleteAt(values []int,index int) []int { return values // TODO: construct an independent result.\n}",
    revisits=(32,34,37))

add(39, "Implement frequencies. Count every word, including an empty string, and return an initialized map for empty input.", """
cases := []struct{ words []string; want map[string]int }{
    {nil,map[string]int{}}, {[]string{"go"},map[string]int{"go":1}},
    {[]string{"go","code","go"},map[string]int{"go":2,"code":1}},
    {[]string{"","go",""},map[string]int{"":2,"go":1}},
}
for _, tc := range cases {
    got := frequencies(tc.words)
    if got == nil { t.Error("return an initialized map"); continue }
    if !maps.Equal(got,tc.want) { t.Errorf("frequencies(%q)=%v; want %v",tc.words,got,tc.want) }
}
""", [("return nil // TODO: initialize and populate the map.", "counts := make(map[string]int)\nfor _,word := range words { counts[word]++ }\nreturn counts")],
    imports=("maps",), declarations="func frequencies(words []string) map[string]int { return nil // TODO: initialize and populate the map.\n}")

observe(40, "Fill both values and presence flags. A missing key and a present key storing zero have the same value but different presence.", """
counts := map[string]int{"zero":0,"go":2}
zero,hasZero := counts["zero"]
missing,hasMissing := counts["missing"]
""", [("ZeroValue", "zero", "-999", "0"), ("MissingValue", "missing", "-999", "0"),
      ("ZeroPresent", "hasZero", "false", "true"), ("MissingPresent", "hasMissing", "true", "false")],
    revisits=(3,39))

add(41, "Implement lookupOr. Return the stored value when the key exists, including zero; use fallback only for an absent key. Nil maps are allowed.", """
counts := map[string]int{"zero":0,"go":2}
cases := []struct{ values map[string]int; key string; fallback,want int }{
    {counts,"zero",9,0}, {counts,"go",9,2}, {counts,"missing",9,9}, {nil,"go",-1,-1},
}
for _, tc := range cases {
    if got := lookupOr(tc.values,tc.key,tc.fallback); got != tc.want { t.Errorf("lookupOr(%v,%q,%d)=%d; want %d",tc.values,tc.key,tc.fallback,got,tc.want) }
}
""", [("return fallback // TODO: distinguish missing from zero.", "value,ok := values[key]\nif ok { return value }\nreturn fallback")],
    declarations="func lookupOr(values map[string]int,key string,fallback int) int { return fallback // TODO: distinguish missing from zero.\n}",
    revisits=(40,))

add(42, "Implement removeKeys. Delete the requested keys in place, preserve other entries, and accept absent keys or a nil map.", """
cases := []struct{ values map[string]int; keys []string; want map[string]int }{
    {map[string]int{"a":1,"b":2},[]string{"a"},map[string]int{"b":2}},
    {map[string]int{"a":1},[]string{"missing","a","a"},map[string]int{}},
    {map[string]int{"a":1},nil,map[string]int{"a":1}},
    {nil,[]string{"a"},nil},
}
for _, tc := range cases {
    removeKeys(tc.values,tc.keys)
    if !maps.Equal(tc.values,tc.want) { t.Errorf("removeKeys(%q): got %v; want %v",tc.keys,tc.values,tc.want) }
}
""", [("// TODO: delete only the requested keys.", "for _,key := range keys { delete(values,key) }")],
    declarations="func removeKeys(values map[string]int,keys []string) { // TODO: delete only the requested keys.\n}",
    imports=("maps",), revisits=(39,40))

add(43, "Fill the map contents and lengths after modifying an assigned map. Assignment shares the entries; it does not clone the map.", """
original := map[string]int{"a":1,"b":2}
alias := original
alias["a"]=9
delete(alias,"b")
original["c"]=3
wantOriginal := map[string]int{} // TODO: all entries.
wantAlias := map[string]int{} // TODO: all entries.
if !maps.Equal(original,wantOriginal) { t.Errorf("original=%v; want %v",original,wantOriginal) }
if !maps.Equal(alias,wantAlias) { t.Errorf("alias=%v; want %v",alias,wantAlias) }
""", [("wantOriginal := map[string]int{} // TODO: all entries.", 'wantOriginal := map[string]int{"a":9,"c":3}'),
      ("wantAlias := map[string]int{} // TODO: all entries.", 'wantAlias := map[string]int{"a":9,"c":3}')],
    imports=("maps",), kind="observe", revisits=(27,31,42))

add(44, "Repair putCount so a nil map can receive its first entry. Return the resulting map and preserve existing entries when it is already initialized.", """
var missing map[string]int
if got := missing["go"]; got != 0 { t.Errorf("nil-map read=%d; want 0",got) }
created := putCount(missing,"go",2)
if !maps.Equal(created,map[string]int{"go":2}) { t.Errorf("created=%v; want map[go:2]",created) }
existing := map[string]int{"old":1}
updated := putCount(existing,"go",3)
if !maps.Equal(updated,map[string]int{"old":1,"go":3}) { t.Errorf("updated=%v",updated) }
if existing["go"] != 3 { t.Error("an existing map must be updated in place") }
""", [("if values == nil { return nil } // TODO: initialize before inserting.", "if values == nil { values=make(map[string]int) }")],
    declarations="""func putCount(values map[string]int,key string,count int) map[string]int {
if values == nil { return nil } // TODO: initialize before inserting.
values[key]=count
return values
}""", imports=("maps",), kind="repair", revisits=(39,43))

add(45, "Implement sign with if/else. Return negative, zero, or positive according to the input.", """
cases := []struct{ value int; want string }{{-7,"negative"}, {-1,"negative"}, {0,"zero"}, {1,"positive"}, {8,"positive"}}
for _, tc := range cases {
    if got := sign(tc.value); got != tc.want { t.Errorf("sign(%d)=%q; want %q",tc.value,got,tc.want) }
}
""", [('return "" // TODO: implement with if/else.', 'if value < 0 { return "negative" }\nif value == 0 { return "zero" }\nreturn "positive"')],
    declarations='func sign(value int) string { return "" // TODO: implement with if/else.\n}')

add(46, "Complete the if initializer and condition. Use comma-ok so a stored zero succeeds, and keep the lookup variable scoped to the if.", """
cases := []struct{ key string; want string }{{"zero","found:0"}, {"go","found:2"}, {"missing","missing"}}
for _, tc := range cases {
    if got := describeLookup(map[string]int{"zero":0,"go":2},tc.key); got != tc.want { t.Errorf("key=%q: got %q; want %q",tc.key,got,tc.want) }
}
""", [("if value,ok := values[key]; value != 0 { // TODO: use the presence flag.\n_ = ok", "if value,ok := values[key]; ok {")],
    imports=("fmt",), declarations="""func describeLookup(values map[string]int,key string) string {
if value,ok := values[key]; value != 0 { // TODO: use the presence flag.
_ = ok
return fmt.Sprintf("found:%d",value)
}
return "missing"
}""", kind="complete", revisits=(40,41))

add(47, "Implement commandName with switch. Start and run share a case; stop has another; any other input returns unknown. Go cases do not fall through automatically.", """
cases := []struct{ command,want string }{{"start","running"}, {"run","running"}, {"stop","stopped"}, {"","unknown"}, {"restart","unknown"}}
for _, tc := range cases {
    if got := commandName(tc.command); got != tc.want { t.Errorf("commandName(%q)=%q; want %q",tc.command,got,tc.want) }
}
""", [('return "" // TODO: implement a switch.', 'switch command { case "start","run": return "running"; case "stop": return "stopped"; default: return "unknown" }')],
    declarations='func commandName(command string) string { return "" // TODO: implement a switch.\n}')

add(48, "Repair the expressionless switch. Test invalid values first, then high scores. Return invalid outside 0..100, pass for 60..100, and fail otherwise.", """
cases := []struct{ score int; want string }{{-1,"invalid"}, {0,"fail"}, {59,"fail"}, {60,"pass"}, {100,"pass"}, {101,"invalid"}}
for _, tc := range cases {
    if got := grade(tc.score); got != tc.want { t.Errorf("grade(%d)=%q; want %q",tc.score,got,tc.want) }
}
""", [('case score >= 60: return "pass" // TODO: put invalid inputs first.\ncase score < 0 || score > 100: return "invalid"', 'case score < 0 || score > 100: return "invalid"\ncase score >= 60: return "pass"')],
    declarations="""func grade(score int) string {
switch {
case score >= 60: return "pass" // TODO: put invalid inputs first.
case score < 0 || score > 100: return "invalid"
default: return "fail"
}
}""", kind="repair", revisits=(13,45))

add(49, "Implement sum using a for loop. Handle empty input, a single value, and negative numbers without changing the input.", """
cases := []struct{ values []int; want int }{{nil,0}, {[]int{7},7}, {[]int{1,2,3},6}, {[]int{-4,2,-1},-3}}
for _, tc := range cases {
    before := slices.Clone(tc.values)
    if got := sum(tc.values); got != tc.want { t.Errorf("sum(%v)=%d; want %d",tc.values,got,tc.want) }
    if !slices.Equal(tc.values,before) { t.Error("sum changed its input") }
}
""", [("return 0 // TODO: accumulate with a loop.", "total := 0\nfor i:=0;i<len(values);i++ { total+=values[i] }\nreturn total")],
    imports=("slices",), declarations="func sum(values []int) int { return 0 // TODO: accumulate with a loop.\n}", revisits=(29,))

add(50, "Implement takeUntilZero. Use break to stop before the first zero; ignore values after it. Return an independent result.", """
cases := []struct{ values,want []int }{
    {nil,nil}, {[]int{0,1},nil}, {[]int{2,3,0,4},[]int{2,3}},
    {[]int{2,3},[]int{2,3}}, {[]int{-1,0},[]int{-1}},
}
for _, tc := range cases {
    before := slices.Clone(tc.values)
    got := takeUntilZero(tc.values)
    if !slices.Equal(got,tc.want) { t.Errorf("takeUntilZero(%v)=%v; want %v",before,got,tc.want) }
    if len(got)>0 { got[0]=99 }
    if !slices.Equal(tc.values,before) { t.Error("the result must have independent storage") }
}
""", [("return nil // TODO: stop at the first zero.", "var result []int\nfor _,value := range values { if value == 0 { break }; result=append(result,value) }\nreturn result")],
    declarations="func takeUntilZero(values []int) []int { return nil // TODO: stop at the first zero.\n}",
    imports=("slices",), revisits=(34,49))

add(51, "Implement positiveValues with continue. Skip nonpositive elements, keep all later positive values in order, and leave the input unchanged.", """
cases := []struct{ values,want []int }{
    {nil,nil}, {[]int{0,-1},nil}, {[]int{2,0,3,-4,5},[]int{2,3,5}}, {[]int{7},[]int{7}},
}
for _, tc := range cases {
    before := slices.Clone(tc.values)
    got := positiveValues(tc.values)
    if !slices.Equal(got,tc.want) { t.Errorf("positiveValues(%v)=%v; want %v",before,got,tc.want) }
    if len(got)>0 { got[0]=99 }
    if !slices.Equal(tc.values,before) { t.Error("input changed") }
}
""", [("return nil // TODO: skip, rather than stop, on nonpositive values.", "var result []int\nfor _,value := range values { if value <= 0 { continue }; result=append(result,value) }\nreturn result")],
    declarations="func positiveValues(values []int) []int { return nil // TODO: skip, rather than stop, on nonpositive values.\n}",
    imports=("slices",), revisits=(50,))

add(52, "Implement positions with range. Return the indices where the target occurs, not the matching values. An absent target returns no indices.", """
cases := []struct{ values []int; target int; want []int }{
    {nil,7,nil}, {[]int{7,2,7},7,[]int{0,2}}, {[]int{2,7},7,[]int{1}},
    {[]int{1,2},7,nil}, {[]int{0,0},0,[]int{0,1}},
}
for _, tc := range cases {
    if got := positions(tc.values,tc.target); !slices.Equal(got,tc.want) { t.Errorf("positions(%v,%d)=%v; want %v",tc.values,tc.target,got,tc.want) }
}
""", [("return nil // TODO: collect the matching indices.", "var result []int\nfor index,value := range values { if value == target { result=append(result,index) } }\nreturn result")],
    declarations="func positions(values []int,target int) []int { return nil // TODO: collect the matching indices.\n}",
    imports=("slices",), revisits=(24,51))

add(53, "Implement indices using range over an integer. Return 0 through n-1; zero or negative n should produce no values.", """
cases := []struct{ n int; want []int }{{-2,nil}, {0,nil}, {1,[]int{0}}, {4,[]int{0,1,2,3}}}
for _, tc := range cases {
    if got := indices(tc.n); !slices.Equal(got,tc.want) { t.Errorf("indices(%d)=%v; want %v",tc.n,got,tc.want) }
}
""", [("return nil // TODO: range over n.", "var result []int\nfor index := range n { result=append(result,index) }\nreturn result")],
    declarations="func indices(n int) []int { return nil // TODO: range over n.\n}", imports=("slices",), revisits=(52,))

add(54, "Implement clamp. Return the nearest bound when value is outside lower..upper, and the value itself otherwise. lower is never greater than upper.", """
cases := []struct{ value,lower,upper,want int }{
    {-1,0,10,0}, {0,0,10,0}, {4,0,10,4}, {10,0,10,10}, {11,0,10,10},
    {-7,-5,-2,-5}, {-3,-5,-2,-3}, {9,4,4,4},
}
for _, tc := range cases {
    if got := clamp(tc.value,tc.lower,tc.upper); got != tc.want { t.Errorf("clamp(%d,%d,%d)=%d; want %d",tc.value,tc.lower,tc.upper,got,tc.want) }
}
""", [("return value // TODO: implement the bounds.", "if value < lower { return lower }\nif value > upper { return upper }\nreturn value")],
    declarations="func clamp(value,lower,upper int) int { return value // TODO: implement the bounds.\n}", revisits=(45,48))

add(55, "Implement first. Return the first element and true for a nonempty slice. Return zero and false for empty input; a stored zero is still present.", """
cases := []struct{ values []int; want int; ok bool }{{nil,0,false}, {[]int{},0,false}, {[]int{0,7},0,true}, {[]int{4,5},4,true}}
for _, tc := range cases {
    got,ok := first(tc.values)
    if got != tc.want || ok != tc.ok { t.Errorf("first(%v)=(%d,%v); want (%d,%v)",tc.values,got,ok,tc.want,tc.ok) }
}
""", [("return 0,false // TODO: return both results.", "if len(values)==0 { return 0,false }\nreturn values[0],true")],
    declarations="func first(values []int) (int,bool) { return 0,false // TODO: return both results.\n}", revisits=(13,40))

add(56, "Complete the named result assignments. Return a quotient and remainder with a bare return; divisor is nonzero.", """
cases := []struct{ value,divisor,quotient,remainder int }{{7,2,3,1}, {-7,2,-3,-1}, {8,4,2,0}, {0,3,0,0}}
for _, tc := range cases {
    quotient,remainder := divide(tc.value,tc.divisor)
    if quotient != tc.quotient || remainder != tc.remainder { t.Errorf("divide(%d,%d)=(%d,%d); want (%d,%d)",tc.value,tc.divisor,quotient,remainder,tc.quotient,tc.remainder) }
}
""", [("// TODO: assign quotient and remainder before the bare return.", "quotient = value / divisor\nremainder = value % divisor")],
    declarations="""func divide(value,divisor int) (quotient,remainder int) {
// TODO: assign quotient and remainder before the bare return.
return
}""", kind="complete", revisits=(10,55))

add(57, "Implement total as a variadic function. Zero arguments return zero; positive and negative values are included.", """
if got := total(); got != 0 { t.Errorf("total()=%d; want 0",got) }
if got := total(7); got != 7 { t.Errorf("total(7)=%d; want 7",got) }
if got := total(1,2,3); got != 6 { t.Errorf("total(1,2,3)=%d; want 6",got) }
if got := total(-4,2); got != -2 { t.Errorf("total(-4,2)=%d; want -2",got) }
""", [("return 0 // TODO: sum the variadic values.", "sum := 0\nfor _,value := range values { sum+=value }\nreturn sum")],
    declarations="func total(values ...int) int { return 0 // TODO: sum the variadic values.\n}", revisits=(49,))

add(58, "Complete addAll by expanding the slice into total's variadic arguments. Preserve the input and include base even for an empty slice.", """
cases := []struct{ base int; values []int; want int }{{5,nil,5}, {1,[]int{2,3},6}, {0,[]int{-4,2},-2}}
for _, tc := range cases {
    before := slices.Clone(tc.values)
    if got := addAll(tc.base,tc.values); got != tc.want { t.Errorf("addAll(%d,%v)=%d; want %d",tc.base,tc.values,got,tc.want) }
    if !slices.Equal(tc.values,before) { t.Error("input changed") }
}
""", [("return base // TODO: forward values with slice expansion.", "return base + total(values...)")],
    declarations="""func total(values ...int) int { sum:=0; for _,value:=range values { sum+=value }; return sum }
func addAll(base int,values []int) int { return base // TODO: forward values with slice expansion.
}""", imports=("slices",), kind="complete", revisits=(57,))

observe_steps(59, "Compare a closure reading an outer variable with a function using a copied argument. Check the closure before and after changing the outer variable, then check the copied argument and the caller's value.", [("""
value := 1
captured := func() int { return value }
copied := func(n int) int { n++; return n }
""", [("Before", "captured()", "-999", "1")]),
("""
value = 5
""", [("After", "captured()", "-999", "5"),
      ("ArgumentResult", "copied(value)", "-999", "6"), ("OuterValue", "value", "-999", "5")])],
    revisits=(6,54))

add(60, "Implement counter. Each factory call returns a closure with its own state; each call adds step and returns the updated value.", """
a := counter(0,1)
b := counter(10,2)
got := []int{a(),a(),b(),a(),b()}
want := []int{1,2,12,3,14}
if !slices.Equal(got,want) { t.Errorf("interleaved counters=%v; want %v",got,want) }
negative := counter(3,-2)
if got := []int{negative(),negative()}; !slices.Equal(got,[]int{1,-1}) { t.Errorf("negative step=%v",got) }
""", [("return func() int { return start } // TODO: retain independent state.", "value := start\nreturn func() int { value+=step; return value }")],
    declarations="func counter(start,step int) func() int { return func() int { return start } // TODO: retain independent state.\n}",
    imports=("slices",), revisits=(59,))

add(61, "Implement transform. Apply the supplied function once to each value, in order, and return independent storage. The function is non-nil; empty input produces no calls.", """
cases := []struct{ values []int; fn func(int) int; want []int }{
    {nil,func(n int) int { return n*2 },nil},
    {[]int{1,2,3},func(n int) int { return n*2 },[]int{2,4,6}},
    {[]int{1,2,3},func(n int) int { return -n },[]int{-1,-2,-3}},
}
for _, tc := range cases {
    before := slices.Clone(tc.values)
    calls := 0
    got := transform(tc.values,func(n int) int { calls++; return tc.fn(n) })
    if !slices.Equal(got,tc.want) { t.Errorf("transform(%v)=%v; want %v",before,got,tc.want) }
    if calls != len(tc.values) { t.Errorf("calls=%d; want %d",calls,len(tc.values)) }
    if len(got)>0 { got[0]=99 }
    if !slices.Equal(tc.values,before) { t.Error("input changed") }
}
""", [("return nil // TODO: call the supplied function for each value.", "result := make([]int,len(values))\nfor i,value := range values { result[i]=fn(value) }\nreturn result")],
    declarations="func transform(values []int,fn func(int) int) []int { return nil // TODO: call the supplied function for each value.\n}",
    imports=("slices",), revisits=(34,52,59))

add(62, "Implement swap through the two non-nil pointers. It must also work when both pointers refer to the same variable.", """
for _, pair := range [][2]int{{2,5},{0,7},{-3,4}} {
    a,b := pair[0],pair[1]
    swap(&a,&b)
    if a != pair[1] || b != pair[0] { t.Errorf("swap(%v)=[%d %d]",pair,a,b) }
}
value := 7
swap(&value,&value)
if value != 7 { t.Errorf("same-pointer swap=%d; want 7",value) }
""", [("// TODO: swap the pointed-to values.", "*a,*b = *b,*a")],
    declarations="func swap(a,b *int) { // TODO: swap the pointed-to values.\n}", revisits=(7,))

observe_steps(63, "Fill the caller's values and the returned value. Changing a copied integer or array parameter does not update its caller.", [("""
number := 3
returned := increment(number)
""", [("Caller", "number", "-999", "3"), ("Returned", "returned", "-999", "4")]),
("""
array := [2]int{1,2}
changeArray(array)
""", [("Array", "array", "[2]int{-1,-1}", "[2]int{1,2}")])],
    declarations="func increment(value int) int { value++; return value }\nfunc changeArray(values [2]int) { values[0]=9 }",
    revisits=(27,54))

add(64, "Repair reset. Changing a local pointer variable does not update the object it originally pointed to; write through the pointer instead. Inputs are non-nil.", """
for _, initial := range []int{7,-3,0} {
    value := initial
    pointer := &value
    originalAddress := pointer
    reset(pointer)
    if value != 0 { t.Errorf("reset(%d): value=%d; want 0",initial,value) }
    if pointer != originalAddress { t.Error("reset changed the caller's pointer") }
}
""", [("value = new(int) // TODO: reset the pointed-to integer.", "*value = 0")],
    declarations="func reset(value *int) { value = new(int) // TODO: reset the pointed-to integer.\n}",
    kind="repair", revisits=(62,63))

add(65, "Implement newCounter using new(int). Each call must return a non-nil pointer to a fresh integer initialized to zero. Updating one counter must not change another. Keep the checks.", """
allocated := newCounter()
if allocated == nil { t.Fatal("newCounter must allocate an integer") }
if *allocated != 0 { t.Errorf("new counter=%d; want 0",*allocated) }
*allocated = 4

second := newCounter()
if second == nil { t.Fatal("each call must allocate an integer") }
if *second != 0 { t.Errorf("second counter=%d; want 0 after changing the first",*second) }
*second = 9
if *allocated != 4 { t.Errorf("first counter=%d; want 4 after changing the second",*allocated) }
""", [("return nil // TODO: allocate a zero-initialized int with new.", "return new(int)")],
    declarations="func newCounter() *int { return nil // TODO: allocate a zero-initialized int with new.\n}",
    revisits=(3,62))

add(66, "Complete the allocation: two zero-valued elements, with capacity for five. Append three values and check that an existing view observes them without changing its own length.", """
buffer := []int{} // TODO: allocate the required length and capacity.
if len(buffer) != 2 || cap(buffer) != 5 { t.Fatalf("len=%d cap=%d; want len=2 cap=5",len(buffer),cap(buffer)) }
if !slices.Equal(buffer,[]int{0,0}) { t.Error("the initial elements must be zero") }
old := buffer
view := buffer[:cap(buffer)]
buffer=append(buffer,3,4,5)
if len(old) != 2 { t.Error("append changed the old slice header") }
if !slices.Equal(view,[]int{0,0,3,4,5}) { t.Errorf("view=%v; want [0 0 3 4 5]",view) }
""", [("buffer := []int{} // TODO: allocate the required length and capacity.", "buffer := make([]int,2,5)")],
    imports=("slices",), kind="complete", revisits=(29,33,65))

add(67, "Implement newPerson with a struct literal. Store the supplied name and age; leave Active at its zero value.", """
cases := []struct{ name string; age int }{{"Ada",30},{"Grace",42},{"",0}}
for _, tc := range cases {
    got := newPerson(tc.name,tc.age)
    want := Person{Name:tc.name,Age:tc.age}
    if got != want { t.Errorf("newPerson(%q,%d)=%+v; want %+v",tc.name,tc.age,got,want) }
    if got.Active { t.Error("Active must default to false") }
}
""", [("return Person{} // TODO: initialize the supplied fields.", "return Person{Name:name,Age:age}")],
    declarations="type Person struct { Name string; Age int; Active bool }\nfunc newPerson(name string,age int) Person { return Person{} // TODO: initialize the supplied fields.\n}",
    revisits=(3,4,5))

add(68, "Fill the expected fields after struct assignment. Scalar fields are copied, while copied slices and pointers still refer to shared storage.", """
number := 4
original := Record{Count:1,Values:[]int{2,3},Pointer:&number}
copied := original
copied.Count=9
wantOriginalCount := -999 // TODO: original scalar.
if original.Count != wantOriginalCount { t.Errorf("original.Count=%d; want %d",original.Count,wantOriginalCount) }
if copied.Count != 9 { t.Error("the copied scalar was not updated") }

copied.Values[0]=7
wantOriginalValues := []int{-999} // TODO: original slice.
if !slices.Equal(original.Values,wantOriginalValues) { t.Errorf("original.Values=%v; want %v",original.Values,wantOriginalValues) }

*copied.Pointer=8
wantOriginalPointedValue := -999 // TODO: original pointed-to value.
if *original.Pointer != wantOriginalPointedValue { t.Errorf("*original.Pointer=%d; want %d",*original.Pointer,wantOriginalPointedValue) }
""", [("wantOriginalCount := -999 // TODO: original scalar.", "wantOriginalCount := 1"),
      ("wantOriginalValues := []int{-999} // TODO: original slice.", "wantOriginalValues := []int{7,3}"),
      ("wantOriginalPointedValue := -999 // TODO: original pointed-to value.", "wantOriginalPointedValue := 8")],
    declarations="type Record struct { Count int; Values []int; Pointer *int }", imports=("slices",),
    kind="observe", revisits=(27,31,63,67))

add(69, "Implement the Kilometers method on Meters. Use the receiver value and a floating-point conversion; do not lose the fractional kilometer.", """
cases := []struct{ meters Meters; want float64 }{{0,0},{1000,1},{1500,1.5},{250,0.25}}
for _, tc := range cases {
    if got := tc.meters.Kilometers(); got != tc.want { t.Errorf("%d meters=%v km; want %v",tc.meters,got,tc.want) }
}
""", [("return 0 // TODO: use the receiver and preserve the fraction.", "return float64(m) / 1000")],
    declarations="type Meters int\nfunc (m Meters) Kilometers() float64 { return 0 // TODO: use the receiver and preserve the fraction.\n}",
    revisits=(10,14,15))

add(70, "Repair the method receiver so Add updates the caller's Counter. Repeated positive and negative additions must accumulate.", """
counter := Counter{N:2}
for _, step := range []struct{ delta,want int }{{3,5},{-2,3},{0,3}} {
    counter.Add(step.delta)
    if counter.N != step.want { t.Errorf("Add(%d): N=%d; want %d",step.delta,counter.N,step.want) }
}
""", [("func (c Counter) Add(delta int) { // TODO: use a receiver that updates the caller.", "func (c *Counter) Add(delta int) {")],
    declarations="type Counter struct{ N int }\nfunc (c Counter) Add(delta int) { // TODO: use a receiver that updates the caller.\nc.N+=delta\n}",
    kind="repair", revisits=(64,67))

observe(71, "Fill the caller state and method result. A value receiver copies the struct, but its slice field still shares an array.", """
counter := Counter{N:2,Values:[]int{4}}
returned := counter.Change()
""", [("CallerNumber", "counter.N", "-999", "2"), ("ReturnedNumber", "returned", "-999", "3"),
      ("SharedElement", "counter.Values[0]", "-999", "5")],
    declarations="""type Counter struct { N int; Values []int }
func (c Counter) Change() int { c.N++; c.Values[0]++; return c.N }""",
    revisits=(63,68,70))

add(72, "Complete newEmployee. Initialize the embedded Person through its field name, then read Name through both promoted and explicit selectors.", """
cases := []struct{ name string; id int }{{"Ada",7},{"Grace",9},{"",0}}
for _, tc := range cases {
    employee := newEmployee(tc.name,tc.id)
    if employee.Name != tc.name || employee.Person.Name != tc.name || employee.ID != tc.id { t.Errorf("employee=%+v; want name=%q id=%d",employee,tc.name,tc.id) }
}
""", [("return Employee{} // TODO: initialize the embedded value and ID.", "return Employee{Person:Person{Name:name},ID:id}")],
    declarations="type Person struct{ Name string }\ntype Employee struct { Person; ID int }\nfunc newEmployee(name string,id int) Employee { return Employee{} // TODO: initialize the embedded value and ID.\n}",
    kind="complete", revisits=(67,))

add(73, "Implement Tagged.Label to override the promoted method. Append a nonempty tag in brackets; keep Base.Label available without the tag.", """
plain := struct{ Base }{Base{Name:"Ada"}}
if got := plain.Label(); got != "Ada" { t.Errorf("promoted label=%q; want Ada",got) }
for _, tc := range []struct{ name,tag,want string }{{"Ada","admin","Ada [admin]"},{"Grace","","Grace"},{"Go","test","Go [test]"}} {
    value := Tagged{Base:Base{Name:tc.name},Tag:tc.tag}
    if got := value.Label(); got != tc.want { t.Errorf("overridden label=%q; want %q",got,tc.want) }
    if got := value.Base.Label(); got != tc.name { t.Errorf("base label=%q; want %q",got,tc.name) }
}
""", [("return t.Base.Label() // TODO: add the tag without changing Base.", 'if t.Tag == "" { return t.Base.Label() }\nreturn t.Base.Label()+" ["+t.Tag+"]"')],
    declarations="""type Base struct{ Name string }
func (b Base) Label() string { return b.Name }
type Tagged struct { Base; Tag string }
func (t Tagged) Label() string { return t.Base.Label() // TODO: add the tag without changing Base.
}""", revisits=(17,69,72))

add(74, "Complete Robot.Name so Robot satisfies Named's behavior. The same consumer must accept people and robots; robot names are robot: followed by the ID.", """
values := []Named{Person{Value:"Ada"},Robot{ID:"r1"},Person{Value:"Grace"},Robot{ID:"r2"}}
want := []string{"Hello Ada","Hello robot:r1","Hello Grace","Hello robot:r2"}
if got := greetings(values); !slices.Equal(got,want) { t.Errorf("greetings=%q; want %q",got,want) }
""", [('return "" // TODO: implement the interface method.', 'return "robot:"+r.ID')],
    declarations="""type Named interface{ Name() string }
type Person struct{ Value string }
func (p Person) Name() string { return p.Value }
type Robot struct{ ID string }
func (r Robot) Name() string { return "" // TODO: implement the interface method.
}
func greetings(values []Named) []string {
var result []string
for _,value := range values { result=append(result,"Hello "+value.Name()) }
return result
}""", imports=("slices",), revisits=(69,73))

add(75, "Implement stringValue with a comma-ok type assertion. Return the string and true only for dynamic type string, including an empty string. Other types must return empty and false without panicking.", """
cases := []struct{ value any; want string; ok bool }{
    {"Go","Go",true}, {"","",true}, {7,"",false}, {nil,"",false}, {Label("Go"),"",false}, {[]byte("Go"),"",false},
}
for _, tc := range cases {
    got,ok := stringValue(tc.value)
    if got != tc.want || ok != tc.ok { t.Errorf("stringValue(%T)=(%q,%v); want (%q,%v)",tc.value,got,ok,tc.want,tc.ok) }
}
""", [('return "",false // TODO: use a safe type assertion.', "text,ok := value.(string)\nreturn text,ok")],
    declarations='type Label string\nfunc stringValue(value any) (string,bool) { return "",false // TODO: use a safe type assertion.\n}',
    revisits=(15,55,74))

add(76, "Implement describeType using a type switch. Recognize nil, int, string, and []int; other dynamic types, including Label, return other.", """
cases := []struct{ value any; want string }{{nil,"nil"}, {7,"int"}, {"Go","string"}, {[]int{1},"slice"}, {Label("Go"),"other"}, {3.5,"other"}}
for _, tc := range cases {
    if got := describeType(tc.value); got != tc.want { t.Errorf("describeType(%T)=%q; want %q",tc.value,got,tc.want) }
}
""", [('return "other" // TODO: switch on the dynamic type.', 'switch value.(type) { case nil: return "nil"; case int: return "int"; case string: return "string"; case []int: return "slice"; default: return "other" }')],
    declarations='type Label string\nfunc describeType(value any) string { return "other" // TODO: switch on the dynamic type.\n}',
    revisits=(47,75))

observe(77, "Fill the nil and equality checks. An interface containing a nil *Failure still has a dynamic type; an unassigned error interface does not.", """
var pointer *Failure
var boxed error = pointer
var empty error
""", [("PointerNil", "pointer == nil", "false", "true"), ("BoxedNil", "boxed == nil", "true", "false"),
      ("EmptyNil", "empty == nil", "false", "true"), ("SameTypedPointer", "boxed == pointer", "false", "true")],
    declarations='type Failure struct{}\nfunc (*Failure) Error() string { return "failure" }', revisits=(65,74,75))

add(78, "Repair the receiver so only *Counter implements Incrementer. The pointer must mutate its counter; a Counter value must fail the runtime interface assertion.", """
_,valueOK := any(Counter{}).(Incrementer)
if valueOK { t.Error("Counter value must not implement Incrementer") }
counter := &Counter{}
incrementer,pointerOK := any(counter).(Incrementer)
if !pointerOK { t.Fatal("*Counter must implement Incrementer") }
incrementer.Increment()
incrementer.Increment()
if counter.N != 2 { t.Errorf("counter.N=%d; want 2",counter.N) }
""", [("func (c Counter) Increment() { // TODO: choose the pointer receiver.", "func (c *Counter) Increment() {")],
    declarations="type Incrementer interface{ Increment() }\ntype Counter struct{ N int }\nfunc (c Counter) Increment() { // TODO: choose the pointer receiver.\nc.N++\n}",
    kind="repair", revisits=(70,71,74,75))

add(79, "Implement requirePositive. Return the provided ErrNonpositive for zero or negative input, and an actual nil error for positive input.", """
for _, tc := range []struct{ value int; invalid bool }{{-7,true},{0,true},{1,false},{9,false}} {
    err := requirePositive(tc.value)
    if tc.invalid && err != ErrNonpositive { t.Errorf("requirePositive(%d)=%v; want ErrNonpositive",tc.value,err) }
    if !tc.invalid && err != nil { t.Errorf("requirePositive(%d)=%v; want nil",tc.value,err) }
}
""", [("return nil // TODO: validate the value.", "if value <= 0 { return ErrNonpositive }\nreturn nil")],
    declarations='var ErrNonpositive = errors.New("value must be positive")\nfunc requirePositive(value int) error { return nil // TODO: validate the value.\n}',
    imports=("errors",), revisits=(45,77))

add(80, "Repair readNumber's error wrapping. Add parse context while preserving strconv's underlying error so callers can recognize syntax and range errors with errors.Is.", """
cases := []struct{ text string; want int; err error }{
    {"42",42,nil}, {"-7",-7,nil}, {"bad",0,strconv.ErrSyntax}, {"",0,strconv.ErrSyntax},
    {"999999999999999999999999999999",0,strconv.ErrRange},
}
for _, tc := range cases {
    value,err := readNumber(tc.text)
    if tc.err == nil {
        if err != nil || value != tc.want { t.Errorf("readNumber(%q)=(%d,%v); want (%d,nil)",tc.text,value,err,tc.want) }
    } else {
        if !errors.Is(err,tc.err) { t.Errorf("readNumber(%q): error %v does not wrap %v",tc.text,err,tc.err) }
        if err == nil || !strings.HasPrefix(err.Error(),"parse: ") { t.Errorf("missing parse context: %v",err) }
    }
}
sameMessage := errors.New(strconv.ErrSyntax.Error())
if errors.Is(sameMessage,strconv.ErrSyntax) { t.Error("matching error text does not imply matching error identity") }
""", [('fmt.Errorf("parse: %v",err) // TODO: preserve the underlying error.', 'fmt.Errorf("parse: %w",err)')],
    declarations="""func readNumber(text string) (int,error) {
value,err := strconv.Atoi(text)
if err != nil { return 0,fmt.Errorf("parse: %v",err) // TODO: preserve the underlying error.
}
return value,nil
}""", imports=("errors","fmt","strconv","strings"), kind="repair", revisits=(55,79))

add(81, "Implement errorPath with errors.As. Find a *PathError through wrapping and return its Path; unrelated errors, nil errors, or a typed nil *PathError return empty and false.", """
cases := []struct{ err error; path string; ok bool }{
    {&PathError{Path:"a.txt"},"a.txt",true},
    {fmt.Errorf("read: %w",&PathError{Path:"b.txt"}),"b.txt",true},
    {errors.New("a.txt"),"",false}, {nil,"",false}, {(*PathError)(nil),"",false},
}
for _, tc := range cases {
    path,ok := errorPath(tc.err)
    if path != tc.path || ok != tc.ok { t.Errorf("errorPath(%v)=(%q,%v); want (%q,%v)",tc.err,path,ok,tc.path,tc.ok) }
}
""", [('return "",false // TODO: extract the typed error and its path.', 'var target *PathError\nif !errors.As(err,&target) || target == nil { return "",false }\nreturn target.Path,true')],
    declarations="""type PathError struct{ Path string }
func (e *PathError) Error() string { return e.Path }
func errorPath(err error) (string,bool) { return "",false // TODO: extract the typed error and its path.
}""", imports=("errors","fmt"), revisits=(75,77,80))

add(82, "Repair runWithCleanup. Register cleanup with defer so it runs once after the work, on both success and early return.", """
for _, fail := range []bool{false,true} {
    var trace []string
    err := runWithCleanup(func() error {
        trace=append(trace,"work")
        if fail { return ErrWork }
        return nil
    },func() { trace=append(trace,"cleanup") })
    if !slices.Equal(trace,[]string{"work","cleanup"}) { t.Errorf("fail=%v: trace=%q; want [work cleanup]",fail,trace) }
    if fail && err != ErrWork { t.Errorf("error=%v; want ErrWork",err) }
    if !fail && err != nil { t.Errorf("error=%v; want nil",err) }
}
""", [("cleanup() // TODO: run cleanup when this function returns.", "defer cleanup()")],
    declarations="""var ErrWork = errors.New("work failed")
func runWithCleanup(work func() error,cleanup func()) error {
cleanup() // TODO: run cleanup when this function returns.
return work()
}""", imports=("errors","slices"), kind="repair", revisits=(61,79))

observe(83, "Fill the complete execution trace. Include ordinary statements and all deferred calls in the order they actually run.", """
trace := ""
func() {
    trace += "start "
    defer func() { trace += "first " }()
    defer func() { trace += "second " }()
    trace += "return "
}()
""", [("Trace", "trace", '"TODO"', '"start return second first "')], revisits=(82,))

observe_steps(84, "Compare a deferred argument with a deferred closure reading the variable later. Then check how defer changes a named result before it reaches the caller.", [("""
value := 1
argument,closure := 0,0
func() {
    defer func(n int) { argument=n }(value)
    defer func() { closure=value }()
    value=2
}()
""", [("Argument", "argument", "-999", "1"), ("Closure", "closure", "-999", "2")]),
("", [("NamedResult", "adjusted()", "-999", "5")])],
    declarations="func adjusted() (result int) { defer func(){result++}(); return 4 }",
    revisits=(56,59,82,83))

add(85, "Implement catchPanic. Run work once; a panic becomes an error containing its value, while normal completion returns a real nil error. Use a deferred recovery in the same goroutine.", """
calls := 0
if err := catchPanic(func(){calls++}); err != nil { t.Errorf("normal work returned %v",err) }
if calls != 1 { t.Errorf("calls=%d; want 1",calls) }
for _, message := range []string{"stop","bad input"} {
    err := catchPanic(func(){panic(message)})
    if err == nil || err.Error() != message { t.Errorf("panic(%q): error=%v",message,err) }
}
if err := catchPanic(func(){panic(7)}); err == nil || err.Error() != "7" { t.Errorf("numeric panic: error=%v",err) }
""", [("return fmt.Errorf(\"not implemented\") // TODO: recover only when work panics.", 'defer func(){ if value:=recover(); value != nil { err=fmt.Errorf("%v",value) } }()\nwork()\nreturn nil')],
    declarations="""func catchPanic(work func()) (err error) {
return fmt.Errorf("not implemented") // TODO: recover only when work panics.
}""", imports=("fmt",), revisits=(56,77,79,82))

observe_steps(86, "Track queue length separately from capacity as values are sent and received. Fill the first received value and each queue length at the point where it is checked.", [("""
ch := make(chan int,2)
""", [("Initial", "len(ch)", "-999", "0"), ("Capacity", "cap(ch)", "-999", "2")]),
("""
ch <- 7
ch <- 0
""", [("Filled", "len(ch)", "-999", "2"), ("First", "<-ch", "-999", "7"),
      ("Remaining", "len(ch)", "-999", "1")])], revisits=(29,66))

add(87, "Complete the send and receive loops. Use the supplied buffered channel to round-trip every string in order; return the values and the number still queued.", """
for _, values := range [][]string{nil,{},{"Go"},{"first","","last"}} {
    got,remaining := roundTrip(values)
    if !slices.Equal(got,values) { t.Errorf("roundTrip(%q)=%q",values,got) }
    if remaining != 0 { t.Errorf("remaining=%d; want 0",remaining) }
}
""", [("// TODO: send each value to ch.", "ch <- value"),
      ("// TODO: receive one value and append it to result.", "result=append(result,<-ch)")],
    declarations="""func roundTrip(values []string) ([]string,int) {
ch := make(chan string,len(values))
for _,value := range values {
_ = value
// TODO: send each value to ch.
}
var result []string
for range values {
// TODO: receive one value and append it to result.
}
return result,len(ch)
}""", imports=("slices",), kind="complete", revisits=(52,53,86))

observe_steps(88, "Fill the value and ok result immediately after each receive. Closing a buffered channel does not discard queued values; only a closed and drained channel returns zero and false.", [("""
ch := make(chan int,2)
ch <- 7
ch <- 8
close(ch)
value,ok := <-ch
""", [("First", "value", "-999", "7"), ("FirstOK", "ok", "false", "true")]),
("""
value,ok = <-ch
""", [("Second", "value", "-999", "8"), ("SecondOK", "ok", "false", "true")]),
("""
value,ok = <-ch
""", [("Drained", "value", "-999", "0"), ("DrainedOK", "ok", "true", "false")])],
    revisits=(40,86,87))

add(89, "Implement drain with range over the channel. Return all queued values, including zero, and finish when the sender closes the channel.", """
for _, values := range [][]int{nil,{0},{2,0,3}} {
    input := make(chan int,len(values))
    for _,value := range values { input <- value }
    close(input)
    finished := make(chan []int,1)
    go func(){finished <- drain(input)}()
    select {
    case got := <-finished:
        if !slices.Equal(got,values) { t.Errorf("drain(%v)=%v",values,got) }
    case <-time.After(time.Second):
        t.Fatal("drain did not stop at channel closure")
    }
}
""", [("return nil // TODO: receive until the input closes.", "var result []int\nfor value := range input { result=append(result,value) }\nreturn result")],
    declarations="func drain(input <-chan int) []int { return nil // TODO: receive until the input closes.\n}",
    imports=("slices","time"), revisits=(52,88))

add(90, "Implement tryReceive with select and default. Return value, ready, and open: a closed channel is ready but not open; an unready or nil channel returns zero, false, false.", """
buffered := make(chan int,2)
type result struct { value int; ready,open bool }
buffered <- 7
buffered <- 0
closed := make(chan int)
close(closed)
cases := []struct{ name string; ch <-chan int; value int; ready,open bool }{
    {"nil",nil,0,false,false}, {"unready",make(chan int),0,false,false},
    {"buffered",buffered,7,true,true}, {"buffered zero",buffered,0,true,true},
    {"closed",closed,0,true,false},
}
for _, tc := range cases {
    t.Run(tc.name,func(t *testing.T){
        finished := make(chan result,1)
        go func(){
            value,ready,open := tryReceive(tc.ch)
            finished <- result{value,ready,open}
        }()
        select {
        case got := <-finished:
            if got.value != tc.value || got.ready != tc.ready || got.open != tc.open { t.Errorf("got %+v; want (value=%d,ready=%v,open=%v)",got,tc.value,tc.ready,tc.open) }
        case <-time.After(time.Second): t.Fatal("tryReceive blocked")
        }
    })
}
""", [("return 0,false,false // TODO: perform a nonblocking receive.", "select { case value,open := <-ch: return value,true,open; default: return 0,false,false }")],
    declarations="func tryReceive(ch <-chan int) (int,bool,bool) { return 0,false,false // TODO: perform a nonblocking receive.\n}",
    imports=("time",), revisits=(88,))

add(91, "Implement squareAsync. Launch a goroutine that sends the square on the supplied unbuffered output, then closes it. Return the channel before the receiver starts.", """
for _, tc := range []struct{ name string; value,want int }{{"positive",3,9},{"negative",-4,16},{"zero",0,0}} {
    t.Run(tc.name,func(t *testing.T){
        returned := make(chan (<-chan int),1)
        go func(){returned <- squareAsync(tc.value)}()
        var output <-chan int
        select {
        case output=<-returned:
        case <-time.After(time.Second): t.Fatal("squareAsync did not return before a receiver started")
        }
        select {
        case got,ok:=<-output:
            if !ok || got != tc.want { t.Errorf("got (%d,%v); want (%d,true)",got,ok,tc.want) }
        case <-time.After(time.Second): t.Fatal("no squared value arrived")
        }
        select {
        case _,ok:=<-output: if ok { t.Error("output should close after one value") }
        case <-time.After(time.Second): t.Fatal("output did not close")
        }
    })
}
""", [("close(output) // TODO: send and close from a goroutine.", "go func(){ defer close(output); output <- value*value }()")],
    declarations="""func squareAsync(value int) <-chan int {
output:=make(chan int)
close(output) // TODO: send and close from a goroutine.
return output
}""", imports=("time",), revisits=(61,87,88))

add(92, "Complete parallelMap using one goroutine per value and a WaitGroup. Keep results in input order and wait for every worker before returning. The barrier test requires the workers to run concurrently.", """
if got := parallelMap(nil,func(n int) int{return n}); len(got) != 0 { t.Errorf("empty result=%v",got) }
values:=[]int{2,3,4}
started:=make(chan int,len(values))
release:=make(chan struct{})
defer close(release)
finished:=make(chan []int,1)
go func(){
    finished <- parallelMap(values,func(n int) int { started<-n; <-release; return n*n })
}()
for range values {
    select {
    case <-started:
    case <-time.After(time.Second): t.Fatal("all workers must start before any worker finishes")
    }
}
// Signal workers without closing release twice on the failure path.
for range values { release <- struct{}{} }
select {
case got:=<-finished:
    if !slices.Equal(got,[]int{4,9,16}) { t.Errorf("results=%v; want [4 9 16]",got) }
case <-time.After(time.Second): t.Fatal("parallelMap did not wait for workers")
}
if !slices.Equal(values,[]int{2,3,4}) { t.Error("input changed") }
""", [("_ = wg\n// TODO: launch workers, write each result to its own index, and wait.", "wg.Add(len(values))\nfor index,value := range values {\ngo func(){ defer wg.Done(); result[index]=fn(value) }()\n}\nwg.Wait()")],
    declarations="""func parallelMap(values []int,fn func(int) int) []int {
result:=make([]int,len(values))
var wg sync.WaitGroup
_ = wg
// TODO: launch workers, write each result to its own index, and wait.
return result
}""", imports=("slices","sync","time"), kind="complete", revisits=(52,61,91))

add(93, "Implement Add and Value using the counter's mutex to protect all accesses. Multiple goroutines update and read the same counter. The downloaded test can also be checked with go test -race.", """
var counter Counter
var wg sync.WaitGroup
for range 200 {
    wg.Add(1)
    go func(){ defer wg.Done(); counter.Add(3); counter.Add(-1); _=counter.Value() }()
}
wg.Wait()
if got:=counter.Value(); got != 400 { t.Errorf("counter=%d; want 400",got) }
counter.Add(-400)
if got:=counter.Value(); got != 0 { t.Errorf("counter after reset=%d; want 0",got) }
""", [("// TODO: protect and update n.", "c.mu.Lock()\ndefer c.mu.Unlock()\nc.n+=delta"),
      ("return 0 // TODO: read n under the mutex.", "c.mu.Lock()\ndefer c.mu.Unlock()\nreturn c.n")],
    declarations="""type Counter struct{ mu sync.Mutex; n int }
func (c *Counter) Add(delta int) { // TODO: protect and update n.
}
func (c *Counter) Value() int { return 0 // TODO: read n under the mutex.
}""", imports=("sync",), revisits=(70,78,92))

add(94, "Implement waitForValue with select. Return a received value, or return ctx.Err when cancellation ends the wait. A nil input disables the receive branch.", """
cancelled,cancel:=context.WithCancel(context.Background())
cancel()
ready:=make(chan int,1)
ready<-7
zero:=make(chan int,1)
zero<-0
expired,stop:=context.WithDeadline(context.Background(),time.Unix(0,0))
defer stop()
cases:=[]struct{ name string; ctx context.Context; input <-chan int; value int; err error }{
    {"ready",context.Background(),ready,7,nil}, {"zero",context.Background(),zero,0,nil},
    {"cancelled",cancelled,nil,0,context.Canceled}, {"expired",expired,nil,0,context.DeadlineExceeded},
}
for _,tc:=range cases {
    t.Run(tc.name,func(t *testing.T){
        type result struct{ value int; err error }
        finished:=make(chan result,1)
        go func(){value,err:=waitForValue(tc.ctx,tc.input); finished<-result{value,err}}()
        select {
        case got:=<-finished:
            if got.value != tc.value || !errors.Is(got.err,tc.err) { t.Errorf("got (%d,%v); want (%d,%v)",got.value,got.err,tc.value,tc.err) }
        case <-time.After(time.Second): t.Fatal("waitForValue did not respond")
        }
    })
}
""", [("return 0,nil // TODO: receive a value or respond to cancellation.", "select { case value:=<-input: return value,nil; case <-ctx.Done(): return 0,ctx.Err() }")],
    declarations="func waitForValue(ctx context.Context,input <-chan int) (int,error) { return 0,nil // TODO: receive a value or respond to cancellation.\n}",
    imports=("context","errors","time"), revisits=(55,80,90,91))

add(95, "Repair cancellable so its child retains the parent's deadline and cancellation. An expired deadline must remain DeadlineExceeded rather than being replaced with Canceled.", """
expired,stop:=context.WithDeadline(context.Background(),time.Unix(0,0))
defer stop()
child,cancelChild:=cancellable(expired)
defer cancelChild()
deadline,ok:=child.Deadline()
if !ok || !deadline.Equal(time.Unix(0,0)) { t.Errorf("child deadline=(%v,%v); want Unix(0,0)",deadline,ok) }
select {
case <-child.Done():
    if !errors.Is(child.Err(),context.DeadlineExceeded) { t.Errorf("child error=%v; want DeadlineExceeded",child.Err()) }
case <-time.After(time.Second): t.Error("expired parent did not stop the child")
}
parent,cancelParent:=context.WithCancel(context.Background())
other,cancelOther:=cancellable(parent)
defer cancelOther()
cancelParent()
select {
case <-other.Done():
    if !errors.Is(other.Err(),context.Canceled) { t.Errorf("child error=%v; want Canceled",other.Err()) }
case <-time.After(time.Second): t.Error("parent cancellation did not stop the child")
}
""", [("return context.WithCancel(context.Background()) // TODO: derive from parent.", "return context.WithCancel(parent)")],
    declarations="func cancellable(parent context.Context) (context.Context,context.CancelFunc) { return context.WithCancel(context.Background()) // TODO: derive from parent.\n}",
    imports=("context","errors","time"), kind="repair", revisits=(82,94))

add(96, "Implement readText with io.ReadAll. Return the text and read error, including any partial data received together with an error.", """
cases:=[]struct{ name string; reader io.Reader; text string; err error }{
    {"ordinary",strings.NewReader("Go"),"Go",nil}, {"empty",strings.NewReader(""),"",nil},
    {"unicode",strings.NewReader("é世"),"é世",nil}, {"partial error",brokenReader{},"partial",ErrRead},
}
for _,tc:=range cases {
    text,err:=readText(tc.reader)
    if text != tc.text || !errors.Is(err,tc.err) { t.Errorf("%s: got (%q,%v); want (%q,%v)",tc.name,text,err,tc.text,tc.err) }
}
""", [('return "",nil // TODO: read all bytes without discarding errors or partial data.', "data,err:=io.ReadAll(reader)\nreturn string(data),err")],
    declarations="""var ErrRead=errors.New("read failed")
type brokenReader struct{}
func (brokenReader) Read(p []byte) (int,error) { return copy(p,"partial"),ErrRead }
func readText(reader io.Reader) (string,error) { return "",nil // TODO: read all bytes without discarding errors or partial data.
}""", imports=("errors","io","strings"), revisits=(19,21,55,74,80))

add(97, "Implement copyAll with io.Copy. Return the actual byte count and preserve reader or writer errors instead of treating partial copying as success.", """
for _,text:=range []string{"","hello","é世"} {
    var out bytes.Buffer
    count,err:=copyAll(&out,strings.NewReader(text))
    if err != nil || count != int64(len(text)) || out.String() != text { t.Errorf("copyAll(%q)=(%d,%v,%q)",text,count,err,out.String()) }
}
count,err:=copyAll(brokenWriter{},strings.NewReader("hello"))
if count != 0 || !errors.Is(err,ErrWrite) { t.Errorf("writer failure=(%d,%v); want (0,ErrWrite)",count,err) }
var partial bytes.Buffer
count,err=copyAll(&partial,brokenReader{})
if count != 7 || !errors.Is(err,ErrRead) || partial.String() != "partial" { t.Errorf("partial read=(%d,%v,%q)",count,err,partial.String()) }
""", [("return 0,nil // TODO: copy and return both results.", "return io.Copy(writer,reader)")],
    declarations="""var ErrRead=errors.New("read failed")
var ErrWrite=errors.New("write failed")
type brokenReader struct{}
func (brokenReader) Read(p []byte) (int,error) { return copy(p,"partial"),ErrRead }
type brokenWriter struct{}
func (brokenWriter) Write([]byte) (int,error) { return 0,ErrWrite }
func copyAll(writer io.Writer,reader io.Reader) (int64,error) { return 0,nil // TODO: copy and return both results.
}""", imports=("bytes","errors","io","strings"), revisits=(55,74,80,96))

add(98, "Repair copyPrefix with io.CopyN. Copy at most n bytes; if the reader ends early, return the bytes copied and io.EOF. n is nonnegative.", """
cases:=[]struct{ text string; n int64; want string; count int64; err error }{
    {"hello world",5,"hello",5,nil}, {"hello",0,"",0,nil},
    {"Go",2,"Go",2,nil}, {"Go",5,"Go",2,io.EOF}, {"",1,"",0,io.EOF},
}
for _,tc:=range cases {
    var out bytes.Buffer
    count,err:=copyPrefix(&out,strings.NewReader(tc.text),tc.n)
    if count != tc.count || !errors.Is(err,tc.err) || out.String() != tc.want { t.Errorf("text=%q n=%d: got (%d,%v,%q); want (%d,%v,%q)",tc.text,tc.n,count,err,out.String(),tc.count,tc.err,tc.want) }
}
""", [("return io.Copy(writer,reader) // TODO: limit the copy to n bytes.", "return io.CopyN(writer,reader,n)")],
    declarations="func copyPrefix(writer io.Writer,reader io.Reader,n int64) (int64,error) { return io.Copy(writer,reader) // TODO: limit the copy to n bytes.\n}",
    imports=("bytes","errors","io","strings"), kind="repair", revisits=(19,80,97))

add(99, "Implement storeText and loadText with os.WriteFile and os.ReadFile. Overwrite existing contents, use mode 0600 when creating a file, and return filesystem errors.", """
directory:=t.TempDir()
filename:=filepath.Join(directory,"sample.txt")
for _,text:=range []string{"Go","","new contents","é世"} {
    if err:=storeText(filename,text); err != nil { t.Fatalf("storeText: %v",err) }
    got,err:=loadText(filename)
    if err != nil || got != text { t.Errorf("loadText=(%q,%v); want (%q,nil)",got,err,text) }
}
if _,err:=loadText(filepath.Join(directory,"missing.txt")); !errors.Is(err,os.ErrNotExist) { t.Errorf("missing file error=%v",err) }
if err:=storeText(filepath.Join(directory,"missing-directory","file.txt"),"data"); !errors.Is(err,os.ErrNotExist) { t.Errorf("missing directory error=%v",err) }
""", [("return nil // TODO: write the file and return its error.", "return os.WriteFile(filename,[]byte(text),0600)"),
      ('return "",nil // TODO: read the file and preserve its error.', "data,err:=os.ReadFile(filename)\nreturn string(data),err")],
    declarations="""func storeText(filename,text string) error { return nil // TODO: write the file and return its error.
}
func loadText(filename string) (string,error) { return "",nil // TODO: read the file and preserve its error.
}""", imports=("errors","os","path/filepath"), revisits=(21,79,80,96))

add(100, "Complete the JSON tags and encodeRecord. Use lowercase name and tags keys; omit an empty nickname. Nil tags encode as null, empty non-nil tags as []. Unexported fields are excluded.", """
cases:=[]struct{ value Record; want string }{
    {Record{Name:"Ada",hidden:"private"},"{\\\"name\\\":\\\"Ada\\\",\\\"tags\\\":null}"},
    {Record{Name:"Go",Tags:[]string{}},"{\\\"name\\\":\\\"Go\\\",\\\"tags\\\":[]}"},
    {Record{Name:"Grace",Tags:[]string{"admin"},Nickname:"G"},"{\\\"name\\\":\\\"Grace\\\",\\\"tags\\\":[\\\"admin\\\"],\\\"nickname\\\":\\\"G\\\"}"},
}
for _,tc:=range cases {
    data,err:=encodeRecord(tc.value)
    if err != nil || string(data) != tc.want { t.Errorf("encodeRecord(%+v)=(%s,%v); want %s",tc.value,data,err,tc.want) }
}
""", [('`json:"Name"` // TODO: use the name key.', '`json:"name"`'),
      ('`json:"nickname"` // TODO: omit an empty nickname.', '`json:"nickname,omitempty"`'),
      ("return json.Marshal(record.Name) // TODO: encode the struct.", "return json.Marshal(record)")],
    declarations="""type Record struct {
Name string `json:"Name"` // TODO: use the name key.
Tags []string `json:"tags"`
Nickname string `json:"nickname"` // TODO: omit an empty nickname.
hidden string
}
func encodeRecord(record Record) ([]byte,error) { return json.Marshal(record.Name) // TODO: encode the struct.
}""", imports=("encoding/json",), kind="complete", revisits=(36,67,68))

add(101, "Implement decodePerson with json.Unmarshal into a pointer. Return decoding errors for malformed JSON or a nonnumeric age; omitted fields keep zero values.", """
cases:=[]struct{ data string; want Person; invalid bool }{
    {"{\\\"name\\\":\\\"Ada\\\",\\\"age\\\":30}",Person{Name:"Ada",Age:30},false},
    {"{}",Person{},false},
    {"{\\\"name\\\":\\\"Grace\\\",\\\"extra\\\":true}",Person{Name:"Grace"},false},
    {"{",Person{},true},
    {"{\\\"age\\\":\\\"30\\\"}",Person{},true},
}
for _,tc:=range cases {
    got,err:=decodePerson([]byte(tc.data))
    if tc.invalid {
        if err == nil { t.Errorf("decodePerson(%q) should return an error",tc.data) }
        continue
    }
    if err != nil || got != tc.want { t.Errorf("decodePerson(%q)=(%+v,%v); want (%+v,nil)",tc.data,got,err,tc.want) }
}
""", [("return Person{},nil // TODO: decode into a pointer and return the error.", "var person Person\nerr:=json.Unmarshal(data,&person)\nreturn person,err")],
    declarations="""type Person struct { Name string `json:"name"`; Age int `json:"age"` }
func decodePerson(data []byte) (Person,error) {
_ = json.Unmarshal
return Person{},nil // TODO: decode into a pointer and return the error.
}""", imports=("encoding/json",), revisits=(3,4,64,79,100))

add(102, "Write the regexp for a complete ID: lowercase go followed by one or more ASCII digits. Reject prefixes, suffixes, spaces, and non-ASCII digits.", """
cases:=[]struct{ text string; want bool }{
    {"go0",true},{"go123",true},{"go",false},{"Go1",false},{"xgo1",false},
    {"go1x",false},{"go-1",false},{"go1\\n",false},{"go١",false},{"",false},
}
for _,tc:=range cases {
    if got:=idPattern.MatchString(tc.text); got != tc.want { t.Errorf("MatchString(%q)=%v; want %v",tc.text,got,tc.want) }
}
""", [('regexp.MustCompile("TODO") // TODO: supply an anchored pattern.', 'regexp.MustCompile(`^go[0-9]+$`)')],
    declarations='var idPattern=regexp.MustCompile("TODO") // TODO: supply an anchored pattern.',
    imports=("regexp",), kind="complete", revisits=(19,20))

add(103, "Complete firstNumber using the input text. Find the first run of ASCII digits, parse it, and return false when no number is present. Numbers in these inputs fit in int.", """
cases:=[]struct{ text string; want int; ok bool }{
    {"go123x",123,true},{"a7b42",7,true},{"zero0",0,true},{"x007",7,true},{"none",0,false},{"",0,false},
}
for _,tc:=range cases {
    got,ok:=firstNumber(tc.text)
    if got != tc.want || ok != tc.ok { t.Errorf("firstNumber(%q)=(%d,%v); want (%d,%v)",tc.text,got,ok,tc.want,tc.ok) }
}
""", [('digits.FindString("") // TODO: search the input text.', "digits.FindString(text)")],
    declarations="""var digits=regexp.MustCompile(`[0-9]+`)
func firstNumber(text string) (int,bool) {
match:=digits.FindString("") // TODO: search the input text.
value,err:=strconv.Atoi(match)
return value,err==nil
}""", imports=("regexp","strconv"), kind="complete", revisits=(55,80,102))

add(104, "Implement generic first. Return the first value and true, or the type's zero value and false for empty input. The same function must handle numbers, strings, and slices.", """
for _,tc:=range []struct{ values []int; want int; ok bool }{{nil,0,false},{[]int{0,7},0,true},{[]int{7,8},7,true}} {
    got,ok:=first(tc.values)
    if got != tc.want || ok != tc.ok { t.Errorf("first(%v)=(%d,%v)",tc.values,got,ok) }
}
for _,tc:=range []struct{ values []string; want string; ok bool }{{nil,"",false},{[]string{"","Go"},"",true},{[]string{"Go","code"},"Go",true}} {
    got,ok:=first(tc.values)
    if got != tc.want || ok != tc.ok { t.Errorf("first(%q)=(%q,%v)",tc.values,got,ok) }
}
got,ok:=first([][]int{{1,2},{3}})
if !ok || !slices.Equal(got,[]int{1,2}) { t.Errorf("first of slices=(%v,%v)",got,ok) }
""", [("return zero,false // TODO: handle both empty and nonempty input.", "if len(values)==0 { return zero,false }\nreturn values[0],true")],
    declarations="func first[T any](values []T) (T,bool) { var zero T; return zero,false // TODO: handle both empty and nonempty input.\n}",
    imports=("slices",), revisits=(3,4,55,75))

add(105, "Implement generic sum. The constraint permits addition for int and float64, and ~ includes defined types with those underlying types.", """
if got:=sum([]int{2,3,-1}); got != 4 { t.Errorf("integer sum=%d; want 4",got) }
if got:=sum([]float64{1.5,2.25}); got != 3.75 { t.Errorf("float sum=%v; want 3.75",got) }
if got:=sum([]Distance{2,5}); got != Distance(7) { t.Errorf("defined int sum=%v; want 7",got) }
if got:=sum([]Weight{0.5,1.25}); got != Weight(1.75) { t.Errorf("defined float sum=%v; want 1.75",got) }
if got:=sum([]int(nil)); got != 0 { t.Errorf("empty int sum=%d; want 0",got) }
if got:=sum([]Weight(nil)); got != 0 { t.Errorf("empty Weight sum=%v; want 0",got) }
""", [("return zero // TODO: accumulate using the constraint's allowed addition.", "for _,value:=range values { zero+=value }\nreturn zero")],
    declarations="""type Distance int
type Weight float64
type Number interface{ ~int | ~float64 }
func sum[T Number](values []T) T { var zero T; return zero // TODO: accumulate using the constraint's allowed addition.
}""", revisits=(14,15,49,104))

add(106, "Implement generic indexOf using ==. Return the first matching index and true; return -1 and false if absent. comparable supports strings, numbers, and comparable structs, but not slices.", """
for _,tc:=range []struct{ values []int; target,index int; ok bool }{
    {nil,7,-1,false},{[]int{7,2,7},7,0,true},{[]int{2,7},7,1,true},{[]int{1,2},7,-1,false},
} {
    index,ok:=indexOf(tc.values,tc.target)
    if index != tc.index || ok != tc.ok { t.Errorf("indexOf(%v,%d)=(%d,%v); want (%d,%v)",tc.values,tc.target,index,ok,tc.index,tc.ok) }
}
if index,ok:=indexOf([]string{"Go","code"},"code"); index != 1 || !ok { t.Errorf("string index=(%d,%v)",index,ok) }
if index,ok:=indexOf([]Point{{1,2},{3,4}},Point{3,4}); index != 1 || !ok { t.Errorf("struct index=(%d,%v)",index,ok) }
""", [("return 0,false // TODO: compare values and return the first index.", "for index,value:=range values { if value == target { return index,true } }\nreturn -1,false")],
    declarations="type Point struct{ X,Y int }\nfunc indexOf[T comparable](values []T,target T) (int,bool) { return 0,false // TODO: compare values and return the first index.\n}",
    revisits=(52,55,67,104,105))

add(107, "Implement sortedCopy with slices.Sort. Return a sorted independent copy; preserve the input and its storage. Empty input may produce nil.", """
cases:=[]struct{ values,want []int }{
    {nil,nil},{[]int{7},[]int{7}},{[]int{3,1,2},[]int{1,2,3}},
    {[]int{3,1,3,1},[]int{1,1,3,3}},{[]int{-1,0,-2},[]int{-2,-1,0}},
}
for _,tc:=range cases {
    before:=slices.Clone(tc.values)
    got:=sortedCopy(tc.values)
    if !slices.Equal(got,tc.want) { t.Errorf("sortedCopy(%v)=%v; want %v",before,got,tc.want) }
    if len(got)>0 { got[0]=99 }
    if !slices.Equal(tc.values,before) { t.Error("sorting or editing the result changed the input") }
}
""", [("return values // TODO: clone before sorting.", "result:=slices.Clone(values)\nslices.Sort(result)\nreturn result")],
    declarations="func sortedCopy(values []int) []int { return values // TODO: clone before sorting.\n}",
    imports=("slices",), revisits=(34,38))

add(108, "Repair cloneGroups. maps.Clone copies entries shallowly; clone each slice too. Preserve nil maps and nil slice values, and make edits independent in both directions.", """
if got:=cloneGroups(nil); got != nil { t.Error("nil input should stay nil") }
original:=map[string][]int{"go":{1,2},"empty":{},"nil":nil}
clone:=cloneGroups(original)
if !maps.EqualFunc(original,clone,func(a,b []int) bool{return slices.Equal(a,b)}) { t.Fatalf("clone=%v; want %v",clone,original) }
if clone["nil"] != nil { t.Error("nil slice became non-nil") }
clone["new"]=[]int{7}
if _,ok:=original["new"]; ok { t.Error("the cloned map shares entries") }
clone["go"][0]=9
if original["go"][0] != 1 { t.Error("the cloned slice still shares its array") }
original["go"][1]=8
if clone["go"][1] != 2 { t.Error("editing the original changed the clone") }
""", [("return maps.Clone(groups) // TODO: also clone the reference-bearing values.", "result:=maps.Clone(groups)\nfor key,values:=range result { result[key]=slices.Clone(values) }\nreturn result")],
    declarations="func cloneGroups(groups map[string][]int) map[string][]int { return maps.Clone(groups) // TODO: also clone the reference-bearing values.\n}",
    imports=("maps","slices"), kind="repair", revisits=(34,36,43,68))

add(109, "Implement parseNumbers. Parse every string with strconv.Atoi. If any conversion fails, return no values and wrap the error with item N: using the zero-based index.", """
cases:=[]struct{ texts []string; want []int; badIndex int; err error }{
    {nil,nil,-1,nil},{[]string{"42","-7","0"},[]int{42,-7,0},-1,nil},
    {[]string{"1","bad","3"},nil,1,strconv.ErrSyntax},{[]string{""},nil,0,strconv.ErrSyntax},
    {[]string{"99999999999999999999999999999"},nil,0,strconv.ErrRange},
}
for _,tc:=range cases {
    got,err:=parseNumbers(tc.texts)
    if tc.err == nil {
        if err != nil || !slices.Equal(got,tc.want) { t.Errorf("parseNumbers(%q)=(%v,%v); want (%v,nil)",tc.texts,got,err,tc.want) }
    } else {
        if got != nil || !errors.Is(err,tc.err) { t.Errorf("invalid input: got (%v,%v)",got,err) }
        prefix:=fmt.Sprintf("item %d:",tc.badIndex)
        if err == nil || !strings.HasPrefix(err.Error(),prefix) { t.Errorf("error=%v; want prefix %q",err,prefix) }
    }
}
""", [("return nil,nil // TODO: parse all values and preserve the first error.", 'var result []int\nfor index,text:=range texts {\nvalue,err:=strconv.Atoi(text)\nif err != nil { return nil,fmt.Errorf("item %d: %w",index,err) }\nresult=append(result,value)\n}\nreturn result,nil')],
    declarations="func parseNumbers(texts []string) ([]int,error) { return nil,nil // TODO: parse all values and preserve the first error.\n}",
    imports=("errors","fmt","slices","strconv","strings"), revisits=(49,80,96,103))

add(110, "Repair inspect. Return the type's Name and its Kind string, not its full type string. A nil interface has no type; a typed nil pointer still does.", """
var nilPointer *Count
cases:=[]struct{ value any; name,kind string }{
    {7,"int","int"},{Count(7),"Count","int"},{&CountValue,"","ptr"},
    {[]int{1},"","slice"},{nil,"","nil"},{nilPointer,"","ptr"},
}
for _,tc:=range cases {
    name,kind:=inspect(tc.value)
    if name != tc.name || kind != tc.kind { t.Errorf("inspect(%T)=(%q,%q); want (%q,%q)",tc.value,name,kind,tc.name,tc.kind) }
}
""", [("return typ.String(),typ.String() // TODO: distinguish type name from kind.", "return typ.Name(),typ.Kind().String()")],
    declarations="""type Count int
var CountValue Count=7
func inspect(value any) (string,string) {
typ:=reflect.TypeOf(value)
if typ == nil { return "","nil" }
return typ.String(),typ.String() // TODO: distinguish type name from kind.
}""", imports=("reflect",), kind="repair", revisits=(15,16,75,77))

add(111, "Write the per-case comparison inside t.Run. The table supplies cases below, inside, and above the range, including both bounds. Keep clamp and the expected results; make failures name the input and expected result.", """
cases:=[]struct{ name string; value,want int }{
    {"below",-1,0},{"lower bound",0,0},{"inside",4,4},{"upper bound",10,10},{"above",11,10},
}
for _,tc:=range cases {
    t.Run(tc.name,func(t *testing.T){
        got:=clamp(tc.value)
        t.Errorf("TODO: compare got with tc.want; got=%d",got)
    })
}
""", [('t.Errorf("TODO: compare got with tc.want; got=%d",got)', 'if got != tc.want { t.Errorf("clamp(%d)=%d; want %d",tc.value,got,tc.want) }')],
    declarations="func clamp(value int) int { if value<0{return 0}; if value>10{return 10}; return value }",
    kind="complete", revisits=(2,48,54))

add(112, "Implement queryName. Read the first name query value, URL-decoded. Return false for a missing or empty name; repeated names use the first value.", """
cases:=[]struct{ url,name string; ok bool }{
    {"/?name=Ada","Ada",true},{"/?name=Ada+Lovelace","Ada Lovelace",true},
    {"/?name=%C3%A9","é",true},{"/","",false},{"/?name=","",false},
    {"/?name=first&name=second","first",true},
}
for _,tc:=range cases {
    request:=httptest.NewRequest("GET",tc.url,nil)
    name,ok:=queryName(request)
    if name != tc.name || ok != tc.ok { t.Errorf("queryName(%q)=(%q,%v); want (%q,%v)",tc.url,name,ok,tc.name,tc.ok) }
}
""", [('return "",false // TODO: read and validate the name query value.', 'name:=request.URL.Query().Get("name")\nreturn name,name!=""')],
    declarations='func queryName(request *http.Request) (string,bool) { return "",false // TODO: read and validate the name query value.\n}',
    imports=("net/http","net/http/httptest"), revisits=(4,55,101))

add(113, "Implement createItem. POST returns status 201 and created followed by a newline. Other methods return 405, Allow: POST, and method not allowed followed by a newline. Set status before writing the body.", """
cases:=[]struct{ method string; status int; body,allow string }{
    {"POST",http.StatusCreated,"created\\n",""},
    {"GET",http.StatusMethodNotAllowed,"method not allowed\\n","POST"},
    {"PUT",http.StatusMethodNotAllowed,"method not allowed\\n","POST"},
}
for _,tc:=range cases {
    recorder:=httptest.NewRecorder()
    createItem(recorder,httptest.NewRequest(tc.method,"/",nil))
    response:=recorder.Result()
    body,err:=io.ReadAll(response.Body)
    response.Body.Close()
    if err != nil { t.Fatal(err) }
    if response.StatusCode != tc.status || string(body) != tc.body || response.Header.Get("Allow") != tc.allow { t.Errorf("%s: status=%d body=%q Allow=%q; want %d %q %q",tc.method,response.StatusCode,body,response.Header.Get("Allow"),tc.status,tc.body,tc.allow) }
}
""", [("// TODO: implement method handling, status, and body.", 'if request.Method != http.MethodPost {\nwriter.Header().Set("Allow","POST")\nhttp.Error(writer,"method not allowed",http.StatusMethodNotAllowed)\nreturn\n}\nwriter.WriteHeader(http.StatusCreated)\n_,_ = writer.Write([]byte("created\\n"))')],
    declarations="func createItem(writer http.ResponseWriter,request *http.Request) { // TODO: implement method handling, status, and body.\n}",
    imports=("io","net/http","net/http/httptest"), revisits=(47,79,96,112))

add(114, "Repair jsonHandler by setting both headers before the body commits the response. Inspect the actual response snapshot, not just the recorder's mutable header map.", """
recorder:=httptest.NewRecorder()
jsonHandler(recorder,httptest.NewRequest("GET","/",nil))
response:=recorder.Result()
defer response.Body.Close()
if got:=response.Header.Get("Content-Type"); got != "application/json" { t.Errorf("Content-Type=%q; want application/json",got) }
if got:=response.Header.Get("Cache-Control"); got != "no-store" { t.Errorf("Cache-Control=%q; want no-store",got) }
body,err:=io.ReadAll(response.Body)
if err != nil { t.Fatal(err) }
if response.StatusCode != http.StatusOK || string(body) != "{\\\"ok\\\":true}\\n" { t.Errorf("status=%d body=%q",response.StatusCode,body) }
""", [('// TODO: move the header changes before the first write.\n_,_ = writer.Write([]byte("{\\\"ok\\\":true}\\n"))\nwriter.Header().Set("Content-Type","application/json")\nwriter.Header().Set("Cache-Control","no-store")',
       'writer.Header().Set("Content-Type","application/json")\nwriter.Header().Set("Cache-Control","no-store")\n_,_ = writer.Write([]byte("{\\\"ok\\\":true}\\n"))')],
    declarations="""func jsonHandler(writer http.ResponseWriter,request *http.Request) {
// TODO: move the header changes before the first write.
_,_ = writer.Write([]byte("{\\\"ok\\\":true}\\n"))
writer.Header().Set("Content-Type","application/json")
writer.Header().Set("Cache-Control","no-store")
}""", imports=("io","net/http","net/http/httptest"), kind="repair", revisits=(100,113))

add(115, "Repair parseDate's layout using Go's reference date. Parse YYYY-MM-DD in UTC, and preserve errors for invalid dates or a different input format.", """
cases:=[]struct{ text string; year int; month time.Month; day int; invalid bool }{
    {"2024-03-14",2024,time.March,14,false},{"2000-02-29",2000,time.February,29,false},
    {"2023-02-29",0,0,0,true},{"14-03-2024",0,0,0,true},{"2024-3-14",0,0,0,true},{"",0,0,0,true},
}
for _,tc:=range cases {
    value,err:=parseDate(tc.text)
    if tc.invalid {
        if err == nil { t.Errorf("parseDate(%q) should fail",tc.text) }
        continue
    }
    if err != nil { t.Errorf("parseDate(%q): %v",tc.text,err); continue }
    year,month,day:=value.Date()
    if year != tc.year || month != tc.month || day != tc.day || value.Location()!=time.UTC { t.Errorf("date=%v; want %04d-%02d-%02d UTC",value,tc.year,tc.month,tc.day) }
}
""", [('return time.Parse("YYYY-MM-DD",text) // TODO: use the Go reference-date layout.', 'return time.Parse("2006-01-02",text)')],
    declarations='func parseDate(text string) (time.Time,error) { return time.Parse("YYYY-MM-DD",text) // TODO: use the Go reference-date layout.\n}',
    imports=("time",), kind="repair", revisits=(55,79,103))

add(116, "Implement billableMinutes for intervals up to 24 hours. Round each positive partial minute up, keep zero at zero, and return ErrOrder when end is before start. Compare instants even if their timezones differ.", """
start:=time.Date(2024,1,1,12,0,0,0,time.UTC)
cases:=[]struct{ end time.Time; want int; err error }{
    {start,0,nil},{start.Add(time.Second),1,nil},{start.Add(time.Minute),1,nil},
    {start.Add(61*time.Second),2,nil},{start.Add(90*time.Minute+30*time.Second),91,nil},
    {start.Add(5*time.Minute).In(time.FixedZone("offset",2*60*60)),5,nil},
    {start.Add(-time.Second),0,ErrOrder},
}
for _,tc:=range cases {
    minutes,err:=billableMinutes(start,tc.end)
    if minutes != tc.want || !errors.Is(err,tc.err) { t.Errorf("end=%v: got (%d,%v); want (%d,%v)",tc.end,minutes,err,tc.want,tc.err) }
}
""", [("return 0,nil // TODO: calculate and round elapsed minutes.", "elapsed:=end.Sub(start)\nif elapsed < 0 { return 0,ErrOrder }\nminutes:=int(elapsed/time.Minute)\nif elapsed%time.Minute != 0 { minutes++ }\nreturn minutes,nil")],
    declarations='var ErrOrder=errors.New("end is before start")\nfunc billableMinutes(start,end time.Time) (int,error) { return 0,nil // TODO: calculate and round elapsed minutes.\n}',
    imports=("errors","time"), revisits=(10,11,14,80,115))

add(117, "Implement triangle. Reject nonpositive sides and degenerate or impossible triangles; return equilateral, isosceles, or scalene. Side order must not matter. Positive sides can be near the largest int, so avoid overflowing their sum.", """
largest:=int(^uint(0)>>1)
valid:=[]struct{ a,b,c int; want string }{
    {2,2,2,"equilateral"},{10,10,10,"equilateral"},{4,4,3,"isosceles"},
    {1,2,2,"isosceles"},{3,4,5,"scalene"},{2,3,4,"scalene"},
    {largest,largest,largest,"equilateral"},{largest-2,largest-1,largest,"scalene"},
}
for _,tc:=range valid {
    for _,sides:=range [][3]int{{tc.a,tc.b,tc.c},{tc.a,tc.c,tc.b},{tc.b,tc.a,tc.c},{tc.b,tc.c,tc.a},{tc.c,tc.a,tc.b},{tc.c,tc.b,tc.a}} {
        got,err:=triangle(sides[0],sides[1],sides[2])
        if err != nil || got != tc.want { t.Errorf("triangle(%v)=(%q,%v); want (%q,nil)",sides,got,err,tc.want) }
    }
}
for _,sides:=range [][3]int{{0,0,0},{0,2,2},{2,-1,2},{2,2,-1},{1,1,3},{3,1,1},{2,4,2},{4,2,2},{largest,1,1}} {
    if _,err:=triangle(sides[0],sides[1],sides[2]); err == nil { t.Errorf("triangle(%v) should return an error",sides) }
}
""", [('return "",errors.New("not implemented") // TODO: classify after validating all sides.',
       'if a<=0 || b<=0 || c<=0 { return "",errors.New("invalid triangle") }\n// Differences of positive ints cannot overflow, unlike their sums.\nif a<=c-b || b<=a-c || c<=b-a { return "",errors.New("invalid triangle") }\nif a==b && b==c { return "equilateral",nil }\nif a==b || a==c || b==c { return "isosceles",nil }\nreturn "scalene",nil')],
    declarations='func triangle(a,b,c int) (string,error) { return "",errors.New("not implemented") // TODO: classify after validating all sides.\n}',
    imports=("errors",), revisits=(13,48,54,79))

add(118, "Implement score for zero to five dice, each from 1 through 6. Three ones score 1000; other triples score 100 times the face. Remaining ones score 100, fives score 50, and other faces score zero. Order must not matter; preserve the input.", """
cases:=[]struct{ dice []int; want int }{
    {nil,0},{[]int{},0},{[]int{5},50},{[]int{1},100},{[]int{1,5,5,1},300},
    {[]int{2,3,4,6},0},{[]int{1,1,1},1000},{[]int{1,1,1,1},1100},{[]int{1,1,1,1,1},1200},
    {[]int{2,2,2},200},{[]int{3,3,3},300},{[]int{4,4,4},400},
    {[]int{5,5,5},500},{[]int{5,5,5,5},550},{[]int{5,5,5,5,5},600},{[]int{6,6,6},600},
    {[]int{1,1,1,5,1},1150},{[]int{3,5,3,1,3},450},{[]int{2,5,2,1,2},350},
}
for _,tc:=range cases {
    before:=slices.Clone(tc.dice)
    if got:=score(tc.dice); got != tc.want { t.Errorf("score(%v)=%d; want %d",before,got,tc.want) }
    if !slices.Equal(tc.dice,before) { t.Error("score changed its input") }
    reversed:=slices.Clone(before)
    slices.Reverse(reversed)
    if got:=score(reversed); got != tc.want { t.Errorf("score reversed(%v)=%d; want %d",before,got,tc.want) }
}
""", [("return -1 // TODO: count faces and score each die once.",
       "counts:=[7]int{}\nfor _,face:=range dice { counts[face]++ }\nresult:=0\nfor face:=1;face<=6;face++ {\ntriple:=face*100\nif face==1 { triple=1000 }\nresult+=(counts[face]/3)*triple\ncounts[face]%=3\n}\nreturn result+counts[1]*100+counts[5]*50")],
    declarations="func score(dice []int) int { return -1 // TODO: count faces and score each die once.\n}",
    imports=("slices",), revisits=(11,26,39,49,107))

add(119, "Implement Roll using Next exactly once per die. n is nonnegative and Next returns 1..6. Store n values in fresh storage, preserve every previous roll, and make no generator calls for zero dice.", """
calls:=0
d:=DiceSet{Next:func()int{calls++;return (calls-1)%6+1}}
d.Roll(3)
first:=d.Values
if !slices.Equal(first,[]int{1,2,3}) || calls != 3 { t.Fatalf("first=%v calls=%d; want [1 2 3],3",first,calls) }
d.Roll(2)
second:=d.Values
if !slices.Equal(second,[]int{4,5}) || calls != 5 { t.Fatalf("second=%v calls=%d; want [4 5],5",second,calls) }
d.Roll(0)
if len(d.Values) != 0 || calls != 5 { t.Errorf("zero dice=%v calls=%d",d.Values,calls) }
d.Roll(5)
if !slices.Equal(d.Values,[]int{6,1,2,3,4}) || calls != 10 { t.Errorf("third=%v calls=%d; want [6 1 2 3 4],10",d.Values,calls) }
if !slices.Equal(first,[]int{1,2,3}) || !slices.Equal(second,[]int{4,5}) { t.Errorf("previous rolls changed: first=%v second=%v",first,second) }
fixedCalls:=0
fixed:=DiceSet{Next:func()int{fixedCalls++;return 6}}
fixed.Roll(4)
if !slices.Equal(fixed.Values,[]int{6,6,6,6}) || fixedCalls != 4 { t.Errorf("must use the supplied generator: values=%v calls=%d",fixed.Values,fixedCalls) }
""", [("// TODO: allocate a fresh roll and call Next once per position.", "d.Values=make([]int,n)\nfor index:=range d.Values { d.Values[index]=d.Next() }")],
    declarations="""type DiceSet struct{ Values []int; Next func() int }
func (d *DiceSet) Roll(n int) { // TODO: allocate a fresh roll and call Next once per position.
}""", imports=("slices",), revisits=(33,34,36,60,61,66,70))

add(120, "Implement squares as a streaming pipeline. Return an output channel immediately; a goroutine squares and sends each input in order, then closes output when input closes. Consumers drain output, and input values fit safely when squared.", """
t.Run("closed empty",func(t *testing.T){
    input:=make(chan int)
    close(input)
    returned:=make(chan (<-chan int),1)
    go func(){returned<-squares(input)}()
    var output <-chan int
    select {
    case output=<-returned:
    case <-time.After(time.Second): t.Fatal("squares did not return for closed input")
    }
    select {
    case _,ok:=<-output: if ok { t.Error("empty input must not produce a value") }
    case <-time.After(time.Second): t.Fatal("empty output did not close")
    }
})
t.Run("buffered values",func(t *testing.T){
    input:=make(chan int,5)
    for _,value:=range []int{2,3,4,-2,0} { input<-value }
    close(input)
    finished:=make(chan []int,1)
    go func(){var values []int; for value:=range squares(input){values=append(values,value)}; finished<-values}()
    select {
    case got:=<-finished: if !slices.Equal(got,[]int{4,9,16,4,0}) { t.Errorf("squares=%v; want [4 9 16 4 0]",got) }
    case <-time.After(time.Second): t.Fatal("buffered pipeline did not finish")
    }
})
t.Run("stream before input closes",func(t *testing.T){
    input:=make(chan int)
    inputClosed:=false
    defer func(){if !inputClosed { close(input) }}()
    returned:=make(chan (<-chan int),1)
    go func(){returned<-squares(input)}()
    var output <-chan int
    select {
    case output=<-returned:
    case <-time.After(time.Second): t.Fatal("squares waited for input before returning")
    }
    for _,value:=range []int{3,-2,0,5} {
        select {
        case input<-value:
        case <-time.After(time.Second): t.Fatal("pipeline did not receive the next input")
        }
        select {
        case got,ok:=<-output:
            if !ok || got != value*value { t.Errorf("input=%d: got (%d,%v); want (%d,true)",value,got,ok,value*value) }
        case <-time.After(time.Second): t.Fatal("pipeline waited for input closure before producing output")
        }
    }
    close(input)
    inputClosed=true
    select {
    case _,ok:=<-output: if ok { t.Error("pipeline produced an extra value") }
    case <-time.After(time.Second): t.Fatal("output did not close after input")
    }
})
""", [("close(output) // TODO: transform input in a goroutine, then close output.", "go func(){ defer close(output); for value:=range input { output<-value*value } }()")],
    declarations="""func squares(input <-chan int) <-chan int {
output:=make(chan int)
close(output) // TODO: transform input in a goroutine, then close output.
return output
}""", imports=("slices","time"), revisits=(53,82,88,89,91,92))


# Seed metadata for this new topic without renumbering existing lessons. A new
# seed starts at version 0 so its first authored starter is published as version 1.
BIT_FLAGS = [
    (121, "read-binary-masks", "Read binary masks", [
        ("Binary integer literals", "Integer_literals"), ("Unsigned integer types", "Numeric_types")]),
    (122, "declare-bit-flags", "Declare bit flags", [
        ("Iota", "Iota"), ("Constant expressions and shifts", "Constant_expressions")]),
    (123, "set-bit-flags", "Set bit flags", [("Bitwise OR (|)", "Arithmetic_operators")]),
    (124, "check-bit-flags", "Check bit flags", [
        ("Bitwise AND (&)", "Arithmetic_operators"), ("Comparisons", "Comparison_operators")]),
    (125, "clear-bit-flags", "Clear bit flags", [("Bit clear / AND NOT (&^)", "Arithmetic_operators")]),
    (126, "toggle-bit-flags", "Toggle bit flags", [("Bitwise XOR (^)", "Arithmetic_operators")]),
    (127, "invert-bits", "Invert bits", [
        ("Unary bitwise complement (^)", "Arithmetic_operators"), ("Unsigned integer widths", "Numeric_types")]),
    (128, "pack-and-unpack-bit-fields", "Pack and unpack bit fields", [
        ("Shifts and bitwise masks", "Arithmetic_operators"), ("Unsigned integer widths", "Numeric_types")]),
]
for number, slug, title, sections in BIT_FLAGS:
    exercise_id = f"{number:03}-{slug}"
    if number <= len(METADATA):
        assert METADATA[number - 1]["id"] == exercise_id
        continue
    assert number == len(METADATA) + 1
    references = [{"title": label, "url": f"https://go.dev/ref/spec#{anchor}"}
                  for label, anchor in sections]
    METADATA.append(dict(id=exercise_id, version=0, topic="Bit flags and operations",
                         title=title, instruction="", reference=references[0]["url"],
                         references=references, code="", kind="implement", revisits=[]))


observe(121, "Fill the decimal values of these uint8 binary masks. The rightmost bit represents 1, and each position to its left doubles the value. Keep the binary literals and checks.", "", [
    ("Read", "uint8(0b00000001)", "uint8(0)", "uint8(1)"),
    ("Write", "uint8(0b00000010)", "uint8(0)", "uint8(2)"),
    ("Execute", "uint8(0b00000100)", "uint8(0)", "uint8(4)"),
    ("Combined", "uint8(0b00000101)", "uint8(0)", "uint8(5)"),
    ("HighBit", "uint8(0b10000000)", "uint8(0)", "uint8(128)"),
    ("AllBits", "uint8(0b11111111)", "uint8(0)", "uint8(255)"),
], revisits=(12,))

add(122, "Complete the first permission constant using 1 << iota. Read, Write, and Execute must occupy distinct bits with masks 1, 2, and 4. Keep the other declarations and checks.", """
cases := []struct{ name string; flag, want Permission }{
    {"Read",Read,0b001}, {"Write",Write,0b010}, {"Execute",Execute,0b100},
}
for _,tc := range cases {
    if tc.flag != tc.want { t.Errorf("%s=%08b; want %08b",tc.name,tc.flag,tc.want) }
}
""", [("Read Permission = 0 // TODO: use 1 << iota.", "Read Permission = 1 << iota")],
    declarations="""type Permission uint8
const (
Read Permission = 0 // TODO: use 1 << iota.
Write
Execute
)""", kind="complete", revisits=(9,121))

add(123, "Implement setFlags using bitwise OR (|). Set every bit selected by mask and preserve the others. Setting an already-set flag must leave it set. Keep the checks.", """
cases := []struct{ flags, mask, want uint8 }{
    {0,0,0}, {0,0b101,0b101}, {0b001,0b010,0b011},
    {0b101,0b001,0b101}, {0b101,0b101,0b101},
    {0b10000000,0b001,0b10000001}, {0xff,0b011,0xff},
}
for _,tc := range cases {
    if got:=setFlags(tc.flags,tc.mask); got != tc.want {
        t.Errorf("setFlags(%08b,%08b)=%08b; want %08b",tc.flags,tc.mask,got,tc.want)
    }
}
""", [("return flags // TODO: set the selected bits.", "return flags | mask")],
    declarations="func setFlags(flags,mask uint8) uint8 { return flags // TODO: set the selected bits.\n}",
    revisits=(121,122))

add(124, "Implement hasAny and hasAll using bitwise AND (&). hasAny reports whether at least one requested bit is set; hasAll requires every requested bit. An empty mask matches all, but does not match any. Keep the checks.", """
cases := []struct{ flags, mask uint8; any, all bool }{
    {0,0,false,true}, {0,0b001,false,false}, {0b001,0b001,true,true},
    {0b001,0b011,true,false}, {0b011,0b011,true,true}, {0b100,0b011,false,false},
    {0b10000000,0b10000000,true,true}, {0b10000001,0b011,true,false},
    {0xff,0xff,true,true}, {0xff,0,false,true},
}
for _,tc := range cases {
    if got:=hasAny(tc.flags,tc.mask); got != tc.any {
        t.Errorf("hasAny(%08b,%08b)=%v; want %v",tc.flags,tc.mask,got,tc.any)
    }
    if got:=hasAll(tc.flags,tc.mask); got != tc.all {
        t.Errorf("hasAll(%08b,%08b)=%v; want %v",tc.flags,tc.mask,got,tc.all)
    }
}
""", [("return false // TODO: test whether any requested bit is set.", "return flags&mask != 0"),
      ("return false // TODO: test whether all requested bits are set.", "return flags&mask == mask")],
    declarations="""func hasAny(flags,mask uint8) bool { return false // TODO: test whether any requested bit is set.
}
func hasAll(flags,mask uint8) bool { return false // TODO: test whether all requested bits are set.
}""", revisits=(13,123))

add(125, "Implement clearFlags using Go's bit-clear operator (&^). Clear every bit selected by mask, preserving all other bits. Clearing an absent flag or an empty mask must do nothing. Keep the checks.", """
cases := []struct{ flags, mask, want uint8 }{
    {0,0xff,0}, {0b111,0b010,0b101}, {0b101,0b010,0b101},
    {0xff,0b00001111,0b11110000}, {0b10000000,0b10000000,0},
    {0b10000001,0b001,0b10000000}, {0b111,0,0b111}, {0xff,0xff,0},
}
for _,tc := range cases {
    if got:=clearFlags(tc.flags,tc.mask); got != tc.want {
        t.Errorf("clearFlags(%08b,%08b)=%08b; want %08b",tc.flags,tc.mask,got,tc.want)
    }
}
""", [("return flags // TODO: clear only the selected bits.", "return flags &^ mask")],
    declarations="func clearFlags(flags,mask uint8) uint8 { return flags // TODO: clear only the selected bits.\n}",
    revisits=(123,124))

add(126, "Implement toggleFlags using XOR (^). Flip each selected bit and preserve unselected bits. Toggling the same mask twice must restore the original flags. Keep the checks.", """
cases := []struct{ flags, mask, want uint8 }{
    {0,0,0}, {0,0b101,0b101}, {0b101,0b001,0b100},
    {0b101,0b010,0b111}, {0b101,0b111,0b010}, {0xff,0xff,0},
    {0b10000001,0b10000000,0b001}, {0b101,0,0b101},
}
for _,tc := range cases {
    got:=toggleFlags(tc.flags,tc.mask)
    if got != tc.want { t.Errorf("toggleFlags(%08b,%08b)=%08b; want %08b",tc.flags,tc.mask,got,tc.want) }
    if restored:=toggleFlags(got,tc.mask); restored != tc.flags {
        t.Errorf("toggling mask %08b twice: got %08b; want %08b",tc.mask,restored,tc.flags)
    }
}
""", [("return flags // TODO: toggle the selected bits.", "return flags ^ mask")],
    declarations="func toggleFlags(flags,mask uint8) uint8 { return flags // TODO: toggle the selected bits.\n}",
    revisits=(12,123,125))

add(127, "Implement invertBits with unary ^. Flip all eight bits of a uint8, including leading zero bits. Unary ^ complements one value; binary ^ toggles bits using a second value. Keep the checks.", """
cases := []struct{ value, want uint8 }{
    {0b00000000,0b11111111}, {0b11111111,0b00000000},
    {0b00001111,0b11110000}, {0b10000000,0b01111111},
    {0b01010101,0b10101010}, {0b11000011,0b00111100},
}
for _,tc := range cases {
    got:=invertBits(tc.value)
    if got != tc.want { t.Errorf("invertBits(%08b)=%08b; want %08b",tc.value,got,tc.want) }
    if restored:=invertBits(got); restored != tc.value {
        t.Errorf("inverting %08b twice: got %08b; want %08b",tc.value,restored,tc.value)
    }
}
""", [("return value // TODO: complement all eight bits.", "return ^value")],
    declarations="func invertBits(value uint8) uint8 { return value // TODO: complement all eight bits.\n}",
    revisits=(121,126))

add(128, "Implement packFields and unpackFields using shifts and masks. A nibble is four bits: pack the low four bits of high into bits 7-4 and the low four bits of low into bits 3-0. Unpack both fields as values from 0 through 15. Keep the checks.", """
packedCases := []struct{ high, low, want uint8 }{
    {0,0,0}, {1,2,0x12}, {0x0a,0x05,0xa5}, {0x0f,0x0f,0xff},
    {0x12,0x34,0x24}, {0xff,0xff,0xff}, {0x80,0x08,0x08},
}
for _,tc := range packedCases {
    if got:=packFields(tc.high,tc.low); got != tc.want {
        t.Errorf("packFields(%08b,%08b)=%08b; want %08b",tc.high,tc.low,got,tc.want)
    }
}
unpackedCases := []struct{ packed, high, low uint8 }{
    {0,0,0}, {0x12,1,2}, {0xa5,0x0a,0x05}, {0xf0,0x0f,0},
    {0x0f,0,0x0f}, {0xff,0x0f,0x0f}, {0x80,8,0},
}
for _,tc := range unpackedCases {
    high,low:=unpackFields(tc.packed)
    if high != tc.high || low != tc.low {
        t.Errorf("unpackFields(%08b)=(%04b,%04b); want (%04b,%04b)",tc.packed,high,low,tc.high,tc.low)
    }
}
for value:=0; value<256; value++ {
    packed:=uint8(value)
    high,low:=unpackFields(packed)
    if high>15 || low>15 { t.Fatalf("fields must fit in four bits: high=%d low=%d",high,low) }
    if got:=packFields(high,low); got != packed {
        t.Fatalf("round trip for %08b: got %08b",packed,got)
    }
}
""", [("return 0 // TODO: mask and pack both fields.", "return (high&0x0f)<<4 | low&0x0f"),
      ("return 0,0 // TODO: extract the high and low nibbles.", "return packed >> 4, packed & 0x0f")],
    declarations="""func packFields(high,low uint8) uint8 { return 0 // TODO: mask and pack both fields.
}
func unpackFields(packed uint8) (uint8,uint8) { return 0,0 // TODO: extract the high and low nibbles.
}""", revisits=(7,121,123,124))


def write_curriculum():
    assert sorted(LESSONS) == list(range(1, len(METADATA) + 1))
    formatter = shutil.which("gofmt")
    if not formatter:
        raise SystemExit("gofmt is required. Add Go 1.23+ to PATH.")
    curriculum = []
    solutions = {}
    for number, (item, solution) in sorted(LESSONS.items()):
        item["code"] = subprocess.run([formatter], input=item["code"],
                                     text=True, capture_output=True, check=True).stdout
        previous = METADATA[number - 1]
        changed = any(item[field] != previous[field] for field in ("code", "instruction"))
        item["version"] = previous.get("version", 1) + int(changed)
        solution = subprocess.run([formatter], input=solution,
                                  text=True, capture_output=True, check=True).stdout
        assert all(int(id[:3]) < number for id in item["revisits"])
        curriculum.append(item)
        solutions[item["id"]] = solution
        directory = ROOT / "koans" / item["id"]
        directory.mkdir(exist_ok=True)
        (directory / "koan_test.go").write_text(item["code"])
    (ROOT / "lib/curriculum.json").write_text(json.dumps(curriculum, ensure_ascii=False, indent=2) + "\n")
    (ROOT / "scripts/solutions.json").write_text(json.dumps(solutions, ensure_ascii=False, indent=2) + "\n")
    print(f"Wrote {len(curriculum)} exercises:", dict(Counter(e["kind"] for e in curriculum)))


if __name__ == "__main__":
    write_curriculum()
