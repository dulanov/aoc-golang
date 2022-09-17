// https://adventofcode.com/2021/day/15
package d15_test

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
	if len(s) == 0 {
		return s, *new(T), false
	}
	return s[:len(s)-1], s[len(s)-1], true
}

type point struct {
	x, y int
}

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 363
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 2835
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(input_test))
	want := 40
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(input_test))
	want := 315
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (n int) {
	return lowest(scan(r))
}

func PartTwo(r io.Reader) (n int) {
	return lowest(expand(scan(r), 5))
}

func expand(ls [][]uint8, n int) [][]uint8 {
	rs := make([][]uint8, len(ls)*n)
	for i := range rs {
		rs[i] = make([]uint8, len(ls[0])*n)
		for j := range rs[i] {
			rs[i][j] = (ls[i%len(ls)][j%len(ls[0])]-1+
				uint8(i/len(ls))+uint8(j/len(ls[0])))%9 + 1
		}
	}
	return rs
}

func lowest(ls [][]uint8) (n int) {
	vs := make([][]bool, len(ls))
	for i := range vs {
		vs[i] = make([]bool, len(ls[0]))
	}
	vs[0][0] = true
	for ss := [10]stack[point]{{point{0, 0}}}; ; n++ {
		for !ss[n%10].empty() {
			var p point
			ss[n%10], p, _ = ss[n%10].pop()
			for _, d := range []point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
				if p.x+d.x < 0 || p.y+d.y < 0 ||
					p.x+d.x >= len(ls[0]) || p.y+d.y >= len(ls) ||
					vs[p.y+d.y][p.x+d.x] {
					continue
				}
				lv := int(ls[p.y+d.y][p.x+d.x])
				if p.x+d.x == len(ls[0])-1 && p.y+d.y == len(ls)-1 {
					return n + lv
				}
				ss[(n+lv)%10], vs[p.y+d.y][p.x+d.x] =
					ss[(n+lv)%10].push(point{p.x + d.x, p.y + d.y}), true
			}
		}
	}
}

func scan(r io.Reader) (ls [][]uint8) {
	for s := bufio.NewScanner(r); s.Scan(); {
		ns := make([]uint8, len(s.Text()))
		for i, r := range s.Text() {
			ns[i] = (uint8)(r - '0')
		}
		ls = append(ls, ns)
	}
	return ls
}

const input_test = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`
