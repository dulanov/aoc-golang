// https://adventofcode.com/2022/day/09
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

type dir byte

const (
	dirUp    dir = 'U'
	dirDown  dir = 'D'
	dirLeft  dir = 'L'
	dirRight dir = 'R'
)

type op struct {
	dir   dir
	steps int
}

type pos struct {
	x, y int
}

func (p pos) cmp(p2 pos) bool {
	return (p.x != p2.x && p.x < p2.x) ||
		(p.x == p2.x && p.y < p2.y)
}

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 5960
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 2327
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest0))
	want := 13
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	for i, tc := range []struct {
		in   string
		want int
	}{
		{
			in:   inputTest0,
			want: 1,
		},
		{
			in:   inputTest1,
			want: 36,
		},
	} {
		t.Run(fmt.Sprintf("inputTest%d", i), func(t *testing.T) {
			got := PartTwo(strings.NewReader(tc.in))
			if got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}
}

func PartOne(r io.Reader) int {
	return sim(scan(r), 1)
}

func PartTwo(r io.Reader) int {
	return sim(scan(r), 9)
}

func sim(ops []op, n int) int {
	ps, vs := make([]pos, n+1), []pos{{}}
	for _, op := range ops {
		for i := 0; i < op.steps; i++ {
			ps[0] = move(ps[0], op.dir)
			for j, p := range ps[1:] {
				if abs(p.x-ps[j].x) > 1 ||
					abs(p.y-ps[j].y) > 1 {
					ps[j+1] = pos{
						p.x + sgn(ps[j].x-p.x),
						p.y + sgn(ps[j].y-p.y)}
				}
			}
			vs = append(vs, ps[len(ps)-1])
		}
	}
	sort.Slice(vs, func(i, j int) bool {
		return vs[i].cmp(vs[j])
	})
	return len(unq(vs))
}

func move(p pos, dir dir) pos {
	switch dir {
	case dirUp:
		p.y += 1
	case dirDown:
		p.y -= 1
	case dirLeft:
		p.x -= 1
	case dirRight:
		p.x += 1
	}
	return p
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sgn(n int) int {
	if n < 0 {
		return -1
	}
	if n > 0 {
		return 1
	}
	return 0
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

func scan(r io.Reader) (ops []op) {
	for s := bufio.NewScanner(r); s.Scan(); {
		var op op
		fmt.Sscanf(s.Text(), "%c %d", &op.dir, &op.steps)
		ops = append(ops, op)
	}
	return ops
}

const (
	inputTest0 = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	inputTest1 = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
)
