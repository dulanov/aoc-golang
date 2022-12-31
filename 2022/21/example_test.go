// https://adventofcode.com/2022/day/21
package d21_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

type stack[T any] []T

func (s stack[T]) empty() bool {
	return len(s) == 0
}

func (s stack[T]) peak() (T, bool) {
	if s.empty() {
		return *new(T), false
	}
	return s[len(s)-1], true
}

func (s stack[T]) pop() (stack[T], T, bool) {
	if s.empty() {
		return s, *new(T), false
	}
	return s[:len(s)-1], s[len(s)-1], true
}

type op byte

const (
	sum op = '+'
	sub op = '-'
	mul op = '*'
	div op = '/'
)

func (o op) calc(n1, n2 int) (n int) {
	switch o {
	case sum:
		return n1 + n2
	case sub:
		return n1 - n2
	case mul:
		return n1 * n2
	case div:
		return n1 / n2
	}
	return 0
}

func (o op) calcRev1(n, n1 int) (n2 int) {
	switch o {
	case sum:
		return n - n1
	case sub:
		return n1 - n
	case mul:
		return n / n1
	case div:
		return n1 / n
	}
	return 0
}

func (o op) calcRev2(n, n2 int) (n1 int) {
	switch o {
	case sum:
		return n - n2
	case sub:
		return n + n2
	case mul:
		return n / n2
	case div:
		return n * n2
	}
	return 0
}

type monkey struct {
	opn int
	ops [2]string
}

func (m monkey) num() (int, bool) {
	if m.ops[0] != "" {
		return 0, false
	}
	return m.opn, true
}

func (m monkey) job() (op, string, string, bool) {
	if m.ops[0] == "" {
		return 0, "", "", false
	}
	return op(m.opn), m.ops[0], m.ops[1], true
}

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 62386792426088
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 3876027196185
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := 152
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest))
	want := 301
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (n int) {
	return calc("root", scan(r))
}

func PartTwo(r io.Reader) (n int) {
	ms := scan(r)
	st := stack[string](prep("root", "humn", ms))
	for {
		var s, sn string
		st, s, _ = st.pop()
		if s == "humn" {
			break
		}
		sn, _ = st.peak()
		op, s1, s2, _ := ms[s].job()
		if s == "root" {
			if s1 == sn {
				n = calc(s2, ms)
			} else {
				n = calc(s1, ms)
			}
			continue
		}
		if s1 == sn {
			n = op.calcRev2(n, calc(s2, ms))
		} else {
			n = op.calcRev1(n, calc(s1, ms))
		}
	}
	return n
}

func prep(sf, st string, ms map[string]monkey) []string {
	if sf == st {
		return []string{st}
	}
	if _, ok := ms[sf].num(); ok {
		return []string{}
	}
	_, s1, s2, _ := ms[sf].job()
	if p := prep(s1, st, ms); len(p) > 0 {
		return append(p, sf)
	}
	if p := prep(s2, st, ms); len(p) > 0 {
		return append(p, sf)
	}
	return []string{}
}

func calc(s string, ms map[string]monkey) int {
	if n, ok := ms[s].num(); ok {
		return n
	}
	op, s1, s2, _ := ms[s].job()
	return op.calc(calc(s1, ms), calc(s2, ms))
}

func scan(r io.Reader) (ms map[string]monkey) {
	ms = make(map[string]monkey)
	for s := bufio.NewScanner(r); s.Scan(); {
		var it monkey
		fs := strings.Fields(s.Text())
		if len(fs) == 2 {
			n, _ := strconv.Atoi(fs[1])
			it.opn = n
		} else {
			it.opn, it.ops = int(fs[2][0]), [2]string{fs[1], fs[3]}
		}
		ms[fs[0][:len(fs[0])-1]] = it
	}
	return ms
}

const inputTest = `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`
