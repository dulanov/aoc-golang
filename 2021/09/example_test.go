// https://adventofcode.com/2021/day/09
package d09_test

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

type point struct {
	col, row int
}

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 591
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 1113424
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(input_test))
	want := 15
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(input_test))
	want := 1134
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (n int) {
	hs := scan(r)
	for _, p := range lowest(hs) {
		n += (int)(hs[p.col][p.row] + 1)
	}
	return n
}

func PartTwo(r io.Reader) (n int) {
	hs := scan(r)
	ns := sizes(hs, lowest(hs))
	sort.Sort(sort.Reverse(sort.IntSlice(ns)))
	return mul(ns[:3])
}

func lowest(hs [][]uint8) (ps []point) {
	rb, rc := hs[0], hs[1]
	for i, rn := range hs[2:] {
		nb, nc := rc[0], rc[1]
		for j, nn := range rc[2:] {
			if nc < nb && nc < nn &&
				nc < rb[j+1] && nc < rn[j+1] {
				ps = append(ps, point{i + 1, j + 1})
			}
			nb, nc = nc, nn
		}
		rb, rc = rc, rn
	}
	return ps
}

func sizes(hs [][]uint8, ps []point) (ns []int) {
	vs := make([][]bool, len(hs))
	for i := range vs {
		vs[i] = make([]bool, len(hs[0]))
	}
	for _, p := range ps {
		var n int
		for st := []point{p}; len(st) != 0; {
			p, st = st[len(st)-1], st[:len(st)-1]
			rb, rc, rn := hs[p.col-1], hs[p.col], hs[p.col+1]
			vb, vc, vn := vs[p.col-1], vs[p.col], vs[p.col+1]
			if vc[p.row] {
				continue
			}
			for p.row--; rc[p.row] != 9; p.row-- {
			}
			for i, js := p.row+1, [2]bool{}; rc[i] != 9; i++ {
				n, vc[i] = n+1, true
				if rb[i] == 9 {
					js[0] = false
				} else if !js[0] && !vb[i] {
					js[0], st = true, append(st, point{p.col - 1, i})
				}
				if rn[i] == 9 {
					js[1] = false
				} else if !js[1] && !vn[i] {
					js[1], st = true, append(st, point{p.col + 1, i})
				}
			}
		}
		ns = append(ns, n)
	}
	return ns
}

func mul(ns []int) int {
	rs := ns[0]
	for _, n := range ns[1:] {
		rs *= n
	}
	return rs
}

func scan(r io.Reader) (hs [][]uint8) {
	var ns []uint8
	for s := bufio.NewScanner(r); s.Scan(); {
		if len(hs) == 0 {
			ns = make([]uint8, len(s.Text())+2)
			for i := range ns {
				ns[i] = 9
			}
			hs = append(hs, ns)
		}
		ns := append(ns[:0:0], ns...)
		for i, r := range s.Text() {
			ns[i+1] = (uint8)(r - '0')
		}
		hs = append(hs, ns)
	}
	hs = append(hs, ns)
	return hs
}

const input_test = `2199943210
3987894921
9856789892
8767896789
9899965678`
