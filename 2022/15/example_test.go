// https://adventofcode.com/2022/day/15
package d15_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"sort"
	"strings"
	"testing"

	"golang.org/x/exp/constraints"
)

//go:embed testdata/input
var input string

type pos [2]int

func (p pos) less(o pos) bool {
	return (p[0] != o[0] && p[0] < o[0]) ||
		(p[0] == o[0] && p[1] < o[1])
}

func (p pos) manhattan(o pos) int {
	return abs(p[0]-o[0]) + abs(p[1]-o[1])
}

type sens struct {
	p pos
	d int
}

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input), 2_000_000))
	// Output:
	// 4873353
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 11600823139120
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest), 10)
	want := 26
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest))
	want := 56_000_011
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader, y int) (n int) {
	ss, bs := scan(r)
	for _, b := range bs {
		if b[1] == y {
			n--
		}
	}
	rs := []pos{} /* ranges */
	for _, s := range ss {
		if l := s.d - abs(s.p[1]-y); l >= 0 {
			rs = append(rs, pos{s.p[0] - l, s.p[0] + l})
		}
	}
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].less(rs[j])
	})
	for i, l := 0, rs[0][0]-1; i < len(rs); i++ {
		if rs[i][0] > l {
			n, l = n+rs[i][1]-rs[i][0]+1, rs[i][1]
		} else if rs[i][1] > l {
			n, l = n+rs[i][1]-l, rs[i][1]
		}
	}
	return n
}

func PartTwo(r io.Reader) (n int) {
	ss, _ := scan(r)
	ls := [4][]int{} /* lines (4x90°) */
	for _, s := range ss {
		for i, d := range [][2]int{{1, -1}, {1, 1}, {-1, -1}, {-1, 1}} {
			ls[i] = append(ls[i], s.p[0]+s.p[1]*d[0]+(s.d+1)*d[1])
		}
	}
	for i := range ls {
		sort.Ints(ls[i])
	}
	ls = [4][]int{intersect(unq(ls[0]), unq(ls[1])),
		intersect(unq(ls[2]), unq(ls[3]))}
	for _, l1 := range ls[0] {
		for _, l2 := range ls[1] {
			if p := (pos{(l1 + l2) / 2, (l1 - l2) / 2}); !reached(ss, p) {
				return p[0]*4_000_000 + p[1]
			}
		}
	}
	return 0
}

func reached(ss []sens, p pos) bool {
	for _, s := range ss {
		if s.p.manhattan(p) <= s.d {
			return true
		}
	}
	return false
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func intersect[T constraints.Ordered](vs1, vs2 []T) (rs []T) {
	for i, j := 0, 0; i < len(vs1) && j < len(vs2); i, j = i+1, j+1 {
		if vs1[i] < vs2[j] {
			j--
		} else if vs1[i] > vs2[j] {
			i--
		} else {
			rs = append(rs, vs1[i])
		}
	}
	return rs
}

func unq[T comparable](vs []T) (rs []T) {
	pr := 0
	for i := 1; i < len(vs); i++ {
		if vs[i] != vs[i-1] {
			vs[pr+1], pr = vs[i], pr+1
		}
	}
	return vs[:pr+1]
}

func scan(r io.Reader) (ss []sens, bs []pos) {
	for s := bufio.NewScanner(r); s.Scan(); {
		var p1, p2 pos
		fmt.Sscanf(s.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&p1[0], &p1[1], &p2[0], &p2[1])
		ss, bs = append(ss, sens{p1, p1.manhattan(p2)}), append(bs, p2)
	}
	sort.Slice(bs, func(i, j int) bool {
		return bs[i].less(bs[j])
	})
	return ss, unq(bs)
}

const inputTest = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`
