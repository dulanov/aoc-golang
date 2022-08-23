// https://adventofcode.com/2021/day/13
package d13_test

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

type dir string

const (
	dirH dir = "x"
	dirV dir = "y"
)

type line struct {
	d dir
	p int
}

type point struct {
	x, y int
}

func (p point) cmp(p2 point) bool {
	return (p.x != p2.x && p.x < p2.x) ||
		(p.x == p2.x && p.y < p2.y)
}

func (p point) rev() point {
	return point{p.y, p.x}
}

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 653
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 0
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(input_test))
	want := 17
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(input_test))
	want := 0
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) int {
	ps, ls := scan(r)
	return len(apply(ps, ls[0], 1))
}

func PartTwo(r io.Reader) int {
	return 0
}

func apply(ps []point, l line, n int) []point {
	rs, n := make([]point, 0, len(ps)), ((l.p+1)<<1)/((1<<n-1)+1)
	for _, p := range ps {
		if l.d == dirV {
			p = p.rev()
		}
		if m := p.x / n; m%2 == 0 {
			p = point{p.x - n*m, p.y}
		} else {
			p = point{n*(m+1) - p.x - 2, p.y}
		}
		if l.d == dirH {
			rs = append(rs, p)
		} else {
			rs = append(rs, p.rev())
		}
	}
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].cmp(rs[j])
	})
	pr := 0
	for i := 1; i < len(rs); i++ {
		if rs[i] != rs[i-1] {
			rs[pr+1], pr = rs[i], pr+1
		}
	}
	return rs[:pr+1]
}

func scan(r io.Reader) (ps []point, ls []line) {
	for s := bufio.NewScanner(r); s.Scan(); {
		if s.Text() == "" {
			continue
		}
		if strings.HasPrefix(s.Text(), "fold") {
			var d dir
			var p int
			fmt.Sscanf(s.Text(), "fold along %1s=%d", &d, &p)
			ls = append(ls, line{d, p})
			continue
		}
		var x, y int
		fmt.Sscanf(s.Text(), "%d,%d", &x, &y)
		ps = append(ps, point{x, y})
	}
	return ps, ls
}

const input_test = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`
