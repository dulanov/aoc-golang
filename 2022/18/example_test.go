// https://adventofcode.com/2022/day/18
package d18_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

type stack[T any] []T

func (s stack[T]) empty() bool {
	return len(s) == 0
}

func (s stack[T]) push(v ...T) stack[T] {
	return append(s, v...)
}

func (s stack[T]) pop() (stack[T], T, bool) {
	if s.empty() {
		return s, *new(T), false
	}
	return s[:len(s)-1], s[len(s)-1], true
}

const (
	up dir = iota
	fwd
	rgt
)

type dir byte

func (d dir) adjacent() (dx, dy, dz byte) {
	switch d {
	case up:
		return 0, 0, 1
	case up.opposite():
		return 0, 0, 255
	case fwd:
		return 0, 1, 0
	case fwd.opposite():
		return 0, 255, 0
	case rgt:
		return 1, 0, 0
	case rgt.opposite():
		return 255, 0, 0
	}
	return 0, 0, 0
}

func (d dir) neighbors() [4]dir {
	switch d {
	case up, up.opposite():
		return [...]dir{fwd, fwd.opposite(), rgt, rgt.opposite()}
	case fwd, fwd.opposite():
		return [...]dir{up, up.opposite(), rgt, rgt.opposite()}
	case rgt, rgt.opposite():
		return [...]dir{up, up.opposite(), fwd, fwd.opposite()}
	}
	return [4]dir{}
}

func (d dir) opposite() dir {
	return 8 - d
}

type side [4]byte

func (s side) neighbors() (ns [4][3]side) {
	/* first attempt - 90° */
	dx, dy, dz := dir(s[3]).adjacent()
	for i, d := range dir(s[3]).neighbors() {
		dx2, dy2, dz2 := d.opposite().adjacent()
		ns[i][0] = side{s[0] + dx + dx2, s[1] + dy + dy2, s[2] + dz + dz2, byte(d)}
	}
	/* second attempt - 0° */
	for i, d := range dir(s[3]).neighbors() {
		dx2, dy2, dz2 := d.opposite().adjacent()
		ns[i][1] = side{s[0] + dx2, s[1] + dy2, s[2] + dz2, s[3]}
	}
	/* third attempt - 270° */
	for i, d := range dir(s[3]).neighbors() {
		ns[i][2] = side{s[0], s[1], s[2], byte(d.opposite())}
	}
	return ns
}

func (s side) opposite() side {
	dx, dy, dz := dir(s[3]).adjacent()
	return side{s[0] + dx, s[1] + dy, s[2] + dz, byte(dir(s[3]).opposite())}
}

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 4314
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 2444
}

func TestPartOne(t *testing.T) {
	for i, tc := range []struct {
		in   string
		want int
	}{
		{
			in:   inputTest0,
			want: 10,
		},
		{
			in:   inputTest1,
			want: 64,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := PartOne(strings.NewReader(tc.in)); got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	for i, tc := range []struct {
		in   string
		want int
	}{
		{
			in:   inputTest0,
			want: 10,
		},
		{
			in:   inputTest1,
			want: 58,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := PartTwo(strings.NewReader(tc.in)); got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}
}

func PartOne(r io.Reader) (n int) {
	return len(prep(scan(r)))
}

func PartTwo(r io.Reader) (n int) {
	var fr side
	ss := prep(scan(r))
	for s := range ss {
		if s[0] > fr[0] && s[3] == byte(rgt) {
			fr = s
		}
	}
	ss[fr] = true
	return dfs(fr, func(s side) (ns []side) {
		for _, fs := range s.neighbors() {
			for _, fr := range fs {
				if v, ok := ss[fr]; ok {
					if !v {
						ns, ss[fr] = append(ns, fr), true
					}
					break
				}
			}
		}
		return ns
	})
}

func dfs(fr side, fn func(s side) []side) (n int) {
	for st := (stack[side]{fr}); !st.empty(); n++ {
		var s side
		st, s, _ = st.pop()
		for _, s = range fn(s) {
			st = st.push(s)
		}
	}
	return n
}

func prep(cs [][3]byte) (ss map[side]bool) {
	ss = map[side]bool{}
	for _, c := range cs {
		for _, d := range []dir{up, up.opposite(), fwd, fwd.opposite(), rgt, rgt.opposite()} {
			s := side{c[0], c[1], c[2], byte(d)}
			if _, ok := ss[s.opposite()]; ok {
				delete(ss, s.opposite())
				continue
			}
			ss[s] = false
		}
	}
	return ss
}

func scan(r io.Reader) (cs [][3]byte) {
	for s := bufio.NewScanner(r); s.Scan(); {
		var x, y, z byte
		fmt.Sscanf(s.Text(), "%d,%d,%d", &x, &y, &z)
		cs = append(cs, [3]byte{x, y, z})
	}
	return cs
}

const (
	inputTest0 = `1,1,1
2,1,1`

	inputTest1 = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`
)
