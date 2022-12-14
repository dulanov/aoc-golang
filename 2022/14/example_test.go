// https://adventofcode.com/2022/day/14
package d14_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

type stack[T any] []T

func (s stack[T]) empty() bool {
	return len(s) == 0
}

func (s stack[T]) push(v ...T) stack[T] {
	return append(s, v...)
}

func (s stack[T]) pop() (stack[T], T, bool) {
	if s.empty() {
		return s, *new(T), false
	}
	return s[:len(s)-1], s[len(s)-1], true
}

type pos [2]int

func (p pos) less(o pos) bool {
	return (p[0] != o[0] && p[0] < o[0]) ||
		(p[0] == o[0] && p[1] < o[1])
}

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 655
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 26484
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := 24
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest))
	want := 93
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (n int) {
	ps, lm := scan(r)
	return sim(ps, lm+1, false)
}

func PartTwo(r io.Reader) (n int) {
	ps, lm := scan(r)
	return sim(ps, lm+2, true)
}

func sim(ps []pos, lm int, floor bool) (n int) {
	vs := make(map[pos]bool, len(ps))
	for _, p := range ps {
		vs[p] = true
	}
	for st := (stack[pos]{{500, 0}}); !st.empty(); n++ {
		var p pos
		st, p, _ = st.pop()
	unit:
		for {
			for _, d := range [][2]int{{0, 1}, {-1, 1}, {1, 1}, {0, 0}} {
				if !floor && p[1]+d[1] == lm /* abyss below */ {
					return n
				}
				if d[0] == 0 && d[1] == 0 /* next unit marker */ {
					break unit
				}
				if p2 := (pos{p[0] + d[0], p[1] + d[1]}); !vs[p2] && (!floor || p2[1] != lm) {
					st, p, vs[p2] = st.push(p), p2, true
					break
				}
			}
		}
	}
	return n
}

func scan(r io.Reader) (ps []pos, lm int) {
	for s := bufio.NewScanner(r); s.Scan(); {
		rs := strings.Split(s.Text(), " -> ")
		for i := 0; i < len(rs)-1; i++ {
			var p1, p2 pos
			fmt.Sscanf(rs[i], "%d,%d", &p1[0], &p1[1])
			fmt.Sscanf(rs[i+1], "%d,%d", &p2[0], &p2[1])
			if p2.less(p1) {
				p1, p2 = p2, p1
			}
			for x := p1[0]; x <= p2[0]; x++ {
				for y := p1[1]; y <= p2[1]; y++ {
					if lm < p2[1] {
						lm = p2[1]
					}
					ps = append(ps, pos{x, y})
				}
			}
		}
	}
	return ps, lm
}

const inputTest = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`
