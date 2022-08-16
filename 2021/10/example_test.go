// https://adventofcode.com/2021/day/10
package d10_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"sort"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

var meta = [...]struct {
	open, closed         chunk
	ilglScore, cmplScore int
}{
	{'(', ')', 3, 1},
	{'{', '}', 57, 3},
	{'[', ']', 1197, 2},
	{'<', '>', 25137, 4},
}

type chunk byte

func (c chunk) opened() bool {
	return c.closedWith() != c
}

func (c chunk) closedWith() chunk {
	el, _, _ := c.meta()
	return el
}

func (c chunk) ilglScore() int {
	_, sc, _ := c.meta()
	return sc
}

func (c chunk) cmplScore() int {
	_, _, sc := c.meta()
	return sc
}

func (c chunk) meta() (chunk, int, int) {
	for _, s := range meta {
		if s.open == c || s.closed == c {
			return s.closed, s.ilglScore, s.cmplScore
		}
	}
	return 0, 0, 0
}

type stack []chunk

func (s stack) push(c chunk) stack {
	return append(s, c.closedWith())
}

func (s stack) drop(c chunk) (stack, bool) {
	if len(s) == 0 || s[len(s)-1] != c {
		return s, false
	}
	return s[:len(s)-1], true
}

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 319233
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 1118976874
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(input_test))
	want := 26397
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(input_test))
	want := 288957
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (sc int) {
	for _, l := range scan(r) {
		if n, _ := score(l); n != 0 {
			sc += n
		}
	}
	return sc
}

func PartTwo(r io.Reader) (sc int) {
	var ns []int
	for _, l := range scan(r) {
		if _, n := score(l); n != 0 {
			ns = append(ns, n)
		}
	}
	sort.Ints(ns)
	return median(ns)
}

func score(l []chunk) (int, int) {
	var st stack = make([]chunk, 0, len(l))
	for _, c := range l {
		if c.opened() {
			st = st.push(c.closedWith())
			continue
		}
		var ok bool
		if st, ok = st.drop(c); !ok {
			return c.ilglScore(), 0
		}
	}
	return 0, compl(st)
}

func compl(st stack) (sc int) {
	for i := len(st) - 1; i >= 0; i-- {
		sc = sc*5 + st[i].cmplScore()
	}
	return sc
}

func median(ns []int) int {
	return ns[len(ns)/2]
}

func scan(r io.Reader) (ls [][]chunk) {
	for s := bufio.NewScanner(r); s.Scan(); {
		cs := make([]chunk, len(s.Text()))
		for i, r := range s.Text() {
			cs[i] = chunk(r)
		}
		ls = append(ls, cs)
	}
	return ls
}

const input_test = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`
