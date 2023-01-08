// https://adventofcode.com/2022/day/23
package d23_test

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

type dir int

func (d dir) adj(p pos) pos {
	switch d {
	case north:
		return pos{p[0], p[1] - 1}
	case south:
		return pos{p[0], p[1] + 1}
	case west:
		return pos{p[0] - 1, p[1]}
	case east:
		return pos{p[0] + 1, p[1]}
	case northeast:
		return pos{p[0] + 1, p[1] - 1}
	case northwest:
		return pos{p[0] - 1, p[1] - 1}
	case southeast:
		return pos{p[0] + 1, p[1] + 1}
	case southwest:
		return pos{p[0] - 1, p[1] + 1}
	}
	return p
}

type pos [2]int

func (p pos) area(p2 pos) int {
	return (abs(p2[0]-p[0]) + 1) * (abs(p2[1]-p[1]) + 1)
}

const (
	north dir = iota
	south
	west
	east

	northeast
	northwest
	southeast
	southwest
)

const (
	empty = '.'
	busy  = '#'
)

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 3947
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 1012
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := 110
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest))
	want := 20
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (n int) {
	ps := scan(r)
	proc(ps, 10)
	p1, p2 := rect(ps)
	return p1.area(p2) - len(ps)
}

func PartTwo(r io.Reader) (n int) {
	return proc(scan(r), -1)
}

func proc(ps map[pos]struct{}, n int) (m int) {
	ds := [4][3]dir{
		{north, northeast, northwest},
		{south, southeast, southwest},
		{west, northwest, southwest},
		{east, northeast, southeast}}
	for i := 0; ; i++ {
		ms := map[pos]pos{}
		for p := range ps {
			for _, d := range []dir{north, south, west, east,
				northeast, northwest, southeast, southwest} {
				if _, ok := ps[d.adj(p)]; !ok {
					continue
				}
			outer:
				for j := 0; j < len(ds); j++ {
					ds := ds[(j+i)%len(ds)]
					for _, d := range ds {
						if _, ok := ps[d.adj(p)]; ok {
							break
						}
						if d == ds[len(ds)-1] {
							p2 := ds[0].adj(p)
							if _, ok := ms[p2]; ok {
								delete(ms, p2)
							} else {
								ms[p2] = p
							}
							break outer
						}
					}
				}
				break
			}
		}
		for p2, p := range ms {
			delete(ps, p)
			ps[p2] = struct{}{}
		}
		if i == n - 1 || len(ms) == 0 {
			return i + 1
		}
	}
}

func rect(ps map[pos]struct{}) (p1, p2 pos) {
	for p := range ps {
		if p[0] < p1[0] {
			p1[0] = p[0]
		}
		if p[1] < p1[1] {
			p1[1] = p[1]
		}
		if p[0] > p2[0] {
			p2[0] = p[0]
		}
		if p[1] > p2[1] {
			p2[1] = p[1]
		}
	}
	return p1, p2
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func scan(r io.Reader) (ps map[pos]struct{}) {
	ps = make(map[pos]struct{})
	for i, s := 0, bufio.NewScanner(r); s.Scan(); i++ {
		for j, c := range s.Text() {
			if c == busy {
				ps[pos{j, i}] = struct{}{}
			}
		}
	}
	return ps
}

const inputTest = `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`
