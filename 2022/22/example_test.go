// https://adventofcode.com/2022/day/22
package d22_test

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

type cube [6]struct {
	i, j, l int
	bs      []byte
}

func (c cube) rotate(p pos) pos {
	switch l := c[p.side].l; p.dir {
	case right: /* -> down */
		p.row, p.col = 0, l-1-p.row
	case down: /* -> left */
		p.row, p.col = p.col, l-1
	case left: /* -> up */
		p.row, p.col = l-1, l-1-p.row
	case up: /* -> right */
		p.row, p.col = p.col, 0
	}
	return pos{p.side, p.row, p.col, (p.dir + 1) % 4}
}

func (c cube) find(i, j int) (int, bool) {
	for n, s := range c {
		if s.i == i && s.j == j {
			return n, true
		}
	}
	return 0, false
}

func (c cube) at(p pos) (byte, bool) {
	s := c[p.side]
	if p.row < 0 || p.row >= s.l ||
		p.col < 0 || p.col >= s.l {
		return empty, false
	}
	return s.bs[p.row*s.l+p.col], true
}

type pos struct {
	side, row, col, dir int
}

func (p pos) next() pos {
	switch p.dir {
	case right:
		p.col++
	case down:
		p.row++
	case left:
		p.col--
	case up:
		p.row--
	}
	return p
}

type instr struct {
	rot, num int
}

const (
	empty = ' '
	open  = '.'
	wall  = '#'
)

const (
	right = iota
	down
	left
	up
)

func ExamplePartOne() {
	ns := PartOne(strings.NewReader(input), 50)
	fmt.Println(ns[0]*1000 + ns[1]*4 + ns[2])
	// Output:
	// 126350
}

func ExamplePartTwo() {
	ns := PartTwo(strings.NewReader(input), 50)
	fmt.Println(ns[0]*1000 + ns[1]*4 + ns[2])
	// Output:
	// 129339
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest), 4)
	want := [3]int{6, 8, right}
	if got != want {
		t.Errorf("got %+v; want %+v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest), 4)
	want := [3]int{5, 7, up}
	if got != want {
		t.Errorf("got %+v; want %+v", got, want)
	}
}

func PartOne(r io.Reader, w int) (ns [3]int) {
	c, irs := scan(r, w)
	p := exec(c, irs, func(n, d int) (int, int) {
		p := pos{col: c[n].i, row: c[n].j, dir: d}
		for p = p.next(); ; p = p.next() {
			if b, ok := c.find((p.col+4)%4, (p.row+4)%4); ok {
				return b, 0
			}
		}
	})
	return [3]int{p.row + c[p.side].j*c[p.side].l + 1,
		p.col + c[p.side].i*c[p.side].l + 1, p.dir}
}

func PartTwo(r io.Reader, w int) (ns [3]int) {
	c, irs := scan(r, w)
	m := make(map[[2]int][2]int, 24)
	if w == 4 {
		m = map[[2]int][2]int{
			{0, right}: {5, 2}, {0, down}: {3, 0}, {0, left}: {2, 3}, {0, up}: {1, 2},
			{1, right}: {2, 0}, {1, down}: {4, 2}, {1, left}: {5, 1}, {1, up}: {0, 2},
			{2, right}: {3, 0}, {2, down}: {4, 3}, {2, left}: {1, 0}, {2, up}: {0, 1},
			{3, right}: {5, 1}, {3, down}: {4, 0}, {3, left}: {2, 0}, {3, up}: {0, 0},
			{4, right}: {5, 0}, {4, down}: {1, 2}, {4, left}: {2, 1}, {4, up}: {3, 0},
			{5, right}: {0, 2}, {5, down}: {1, 3}, {5, left}: {4, 0}, {5, up}: {3, 3},
		}
	} else /* w == 50 */ {
		m = map[[2]int][2]int{
			{0, right}: {1, 0}, {0, down}: {2, 0}, {0, left}: {3, 2}, {0, up}: {5, 1},
			{1, right}: {4, 2}, {1, down}: {2, 1}, {1, left}: {0, 0}, {1, up}: {5, 0},
			{2, right}: {1, 3}, {2, down}: {4, 0}, {2, left}: {3, 3}, {2, up}: {0, 0},
			{3, right}: {4, 0}, {3, down}: {5, 0}, {3, left}: {0, 2}, {3, up}: {2, 1},
			{4, right}: {1, 2}, {4, down}: {5, 1}, {4, left}: {3, 0}, {4, up}: {2, 0},
			{5, right}: {4, 3}, {5, down}: {1, 0}, {5, left}: {0, 3}, {5, up}: {3, 0},
		}
	}
	p := exec(c, irs, func(n, d int) (int, int) {
		ns := m[[2]int{n, d}]
		return ns[0], ns[1]
	})
	return [3]int{p.row + c[p.side].j*c[p.side].l + 1,
		p.col + c[p.side].i*c[p.side].l + 1, p.dir}
}

func exec(c cube, irs []instr, tr func(s, d int) (int, int)) (p pos) {
	p.col = bytes.IndexByte(c[0].bs, open)
	for _, ir := range irs {
		p.dir = (p.dir + ir.rot) % 4
		for i, p2 := 0, p.next(); i < ir.num; i, p, p2 = i+1, p2, p2.next() {
			if b, ok := c.at(p2); ok {
				if b == wall {
					break
				}
				continue
			}
			l := c[p.side].l
			n1, n2 := tr(p.side, p.dir)
			p2.side, p2.col, p2.row = n1, (p2.col+l)%l, (p2.row+l)%l
			for i := 0; i < n2; i++ {
				p2 = c.rotate(p2)
			}
			if b, _ := c.at(p2); b == wall {
				break
			}
		}
	}
	return p
}

func scan(r io.Reader, w int) (c cube, irs []instr) {
	re := strings.NewReplacer("R", "|R|", "L", "|L|")
	for i, n, s := 0, 0, bufio.NewScanner(r); s.Scan(); i++ {
		if s.Text() == "" /* instructions */ {
			s.Scan()
			ss := strings.Split(re.Replace(s.Text()), "|")
			for i, ir := -1, (instr{rot: 0}); i < len(ss)-1; i += 2 {
				if i != -1 {
					if ss[i] == "R" {
						ir.rot = 1 /* clockwise */
					} else /* L */ {
						ir.rot = 3 /* counter-clockwise */
					}
				}
				ir.num, _ = strconv.Atoi(ss[i+1])
				irs = append(irs, ir)
			}
			continue
		}
		if i%w == 0 /* new sides */ {
			for j := 0; j < len(s.Text()); j += w {
				if s.Text()[j] == empty {
					continue
				}
				c[n].i, c[n].j, c[n].l, n = j/w, i/w, w, n+1
			}
		}
		bs := strings.ReplaceAll(s.Text(), string(empty), "")
		for j := 0; j < len(bs); j += w {
			sd := &c[n-(len(bs)-j)/w]
			sd.bs = append(sd.bs, bs[j:j+w]...)
		}
	}
	return c, irs
}

const inputTest = `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`
