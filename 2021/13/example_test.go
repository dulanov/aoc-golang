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

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 653
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// #....#..#.###..####.###..###..###..#..#.
	// #....#.#..#..#.#....#..#.#..#.#..#.#.#..
	// #....##...#..#.###..###..#..#.#..#.##...
	// #....#.#..###..#....#..#.###..###..#.#..
	// #....#.#..#.#..#....#..#.#....#.#..#.#..
	// ####.#..#.#..#.####.###..#....#..#.#..#.
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
	want := `#####
#...#
#...#
#...#
#####
.....
.....`
	if got != want {
		t.Errorf("got \n%v; want \n%v", got, want)
	}
}

func PartOne(r io.Reader) int {
	ps, ls := scan(r)
	fold(ps, ls[0], 1)
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].cmp(ps[j])
	})
	return len(uniq(ps))
}

func PartTwo(r io.Reader) string {
	ps, ls := scan(r)
	l1, l2, n1, n2 := ls[0], line{}, 1, 0
	for _, l := range ls[1:] {
		if l1.d != l.d && l2.p == 0 {
			l2 = l
		}
		if l1.d == l.d {
			n1++
		} else {
			n2++
		}
	}
	if l1.d == dirH {
		l1, l2, n1, n2 = l2, l1, n2, n1
	}
	h := fold(ps, l1, n1)
	w := fold(ps, l2, n2)
	ss := make([]string, h)
	for i := range ss {
		ss[i] = strings.Repeat(".", w)
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].cmp(ps[j])
	})
	for _, p := range uniq(ps) {
		ss[p.y] = string(ss[p.y][:p.x]) + "#" + string(ss[p.y][p.x+1:])
	}
	return strings.Join(ss, "\n")
}

func fold(ps []point, l line, n int) int {
	n = ((l.p + 1) << 1) / ((1<<n - 1) + 1)
	for i := range ps {
		if l.d == dirH {
			ps[i].x = abs(ps[i].x - (ps[i].x/n)*n - ((ps[i].x/n)%2)*(n-2))
		} else {
			ps[i].y = abs(ps[i].y - (ps[i].y/n)*n - ((ps[i].y/n)%2)*(n-2))
		}
	}
	return n - 1
}

func uniq(ps []point) []point {
	pr := 0
	for i := 1; i < len(ps); i++ {
		if ps[i] != ps[i-1] {
			ps[pr+1], pr = ps[i], pr+1
		}
	}
	return ps[:pr+1]
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
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
