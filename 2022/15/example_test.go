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

type reg = pos

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input), 2000000))
	// Output:
	// 4873353
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 0
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
	want := 0
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader, y int) (n int) {
	bs, ss := scan(r)
	rs := []reg{}
	for _, s := range ss {
		if l := s.d - abs(s.p[1]-y); l >= 0 {
			rs = append(rs, reg{s.p[0] - l, s.p[0] + l})
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
	for _, b := range bs {
		if b[1] == y {
			n--
		}
	}
	return n
}

func PartTwo(r io.Reader) int {
	return 0
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
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

func scan(r io.Reader) (bs []pos, ss []struct {
	p pos
	d int
}) {
	for s := bufio.NewScanner(r); s.Scan(); {
		var p1, p2 pos
		fmt.Sscanf(s.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&p1[0], &p1[1], &p2[0], &p2[1])
		bs, ss = append(bs, p2), append(ss, struct {
			p pos
			d int
		}{p1, p1.manhattan(p2)})
	}
	sort.Slice(bs, func(i, j int) bool {
		return bs[i].less(bs[j])
	})
	return unq(bs), ss
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
