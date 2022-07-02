// https://adventofcode.com/2021/day/05
package d05_test

import (
	_ "embed"
	"fmt"
	"io"
	"sort"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

type line struct {
	a, b point
}

type point struct {
	x, y int
}

func (p point) cmp(p2 point) bool {
	return (p.x != p2.x && p.x < p2.x) ||
		(p.x == p2.x && p.y < p2.y)
}

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 5167
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 17604
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(input_test))
	want := 5
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(input_test))
	want := 12
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) int {
	return overlaps(scan(r), func(l line) bool {
		/* horizontal and vertical lines only */
		return l.a.x == l.b.x || l.a.y == l.b.y
	}, func(ps []point) int {
		if len(ps) == 1 {
			return 0
		} else {
			return 1 /* at least two lines overlap */
		}
	})
}

func PartTwo(r io.ReadSeeker) int {
	return overlaps(scan(r), func(l line) bool {
		/* all lines */
		return true
	}, func(ps []point) int {
		if len(ps) == 1 {
			return 0
		} else {
			return 1 /* at least two lines overlap */
		}
	})
}

func overlaps(ls []line, fn func(line) bool, gn func([]point) int) int {
	ps := points(ls, fn)
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].cmp(ps[j])
	})
	return sum(group2(ps, gn))
}

func points(ls []line, fn func(line) bool) (ps []point) {
	for _, l := range ls {
		if !fn(l) {
			continue
		}
		var dx, dy int
		if l.a.x < l.b.x {
			dx = 1
		} else if l.a.x > l.b.x {
			dx = -1
		}
		if l.a.y < l.b.y {
			dy = 1
		} else if l.a.y > l.b.y {
			dy = -1
		}
		for i := 0; i <= ((l.b.x-l.a.x)*dx+(l.b.y-l.a.y)*dy)/(dx*dx+dy*dy); i++ {
			ps = append(ps, point{l.a.x + dx*i, l.a.y + dy*i})
		}
	}
	return ps
}

func group2(ps []point, gn func([]point) int) (ns []int) {
	for i := 0; i < len(ps); i++ {
		s := []point{ps[i]}
		for j := i + 1; j < len(ps); j++ {
			if ps[j] != ps[i] {
				i += j - i - 1
				break
			}
			s = append(s, ps[j])
		}
		ns = append(ns, gn(s))
	}
	return ns
}

func sum(ns []int) (rs int) {
	for _, n := range ns {
		rs += n
	}
	return rs
}

func scan(r io.Reader) (ls []line) {
	for {
		var l line
		if _, err := fmt.Fscanf(r, "%d,%d -> %d,%d\n", &l.a.x, &l.a.y, &l.b.x, &l.b.y); err == io.EOF {
			return ls
		}
		ls = append(ls, l)
	}
}

const input_test = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
