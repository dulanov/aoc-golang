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
	return overlaps(scan(r), func(l [2]point) bool {
		/* horizontal and vertical lines only */
		return l[0].x == l[1].x || l[0].y == l[1].y
	}, func(ps []point) int {
		if len(ps) == 1 {
			return 0
		} else {
			return 1 /* at least two lines overlap */
		}
	})
}

func PartTwo(r io.ReadSeeker) int {
	return overlaps(scan(r), func(l [2]point) bool {
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

func overlaps(ls [][2]point, fn func([2]point) bool, gn func([]point) int) int {
	ps := points(ls, fn)
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].cmp(ps[j])
	})
	return sum(group2(ps, gn))
}

func points(ls [][2]point, fn func([2]point) bool) (ps []point) {
	for _, l := range ls {
		if !fn(l) {
			continue
		}
		dp := point{
			dlt(l[0].x, l[1].x),
			dlt(l[0].y, l[1].y)}
		ps = append(ps, l[0])
		for l[0] != l[1] {
			l[0].x += dp.x
			l[0].y += dp.y
			ps = append(ps, l[0])
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

func dlt(a, b int) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

func sum(ns []int) (rs int) {
	for _, n := range ns {
		rs += n
	}
	return rs
}

func scan(r io.Reader) (ls [][2]point) {
	for {
		var l [2]point
		if _, err := fmt.Fscanf(r, "%d,%d -> %d,%d\n", &l[0].x, &l[0].y, &l[1].x, &l[1].y); err == io.EOF {
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
