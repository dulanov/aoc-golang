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
	dir dir
	num int
}

type pos [2]int

func (p pos) less(o pos) bool {
	return (p[0] != o[0] && p[0] < o[0]) ||
		(p[0] == o[0] && p[1] < o[1])
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
		for i := 0; i < op.num; i++ {
			ps[0] = move(ps[0], op.dir)
			for j, p := range ps[1:] {
				if abs(p[0]-ps[j][0]) > 1 ||
					abs(p[1]-ps[j][1]) > 1 {
					ps[j+1] = pos{
						p[0] + sgn(ps[j][0]-p[0]),
						p[1] + sgn(ps[j][1]-p[1])}
				}
			}
			if vs[len(vs)-1] != ps[len(ps)-1] {
				vs = append(vs, ps[len(ps)-1])
			}
		}
	}
	sort.Slice(vs, func(i, j int) bool {
		return vs[i].less(vs[j])
	})
	return len(unq(vs))
}

func move(p pos, dir dir) pos {
	switch dir {
	case dirUp:
		p[1] += 1
	case dirDown:
		p[1] -= 1
	case dirLeft:
		p[0] -= 1
	case dirRight:
		p[0] += 1
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
		fmt.Sscanf(s.Text(), "%c %d", &op.dir, &op.num)
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
