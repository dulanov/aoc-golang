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

const (
	up dir = iota
	fwd
	rgt
)

type dir byte

func (d dir) opposite() dir {
	return 8 - d
}

type side [4]byte

func (s side) opposite() side {
	switch d := dir(s[3]); d {
	case up:
		return side{s[0], s[1], s[2] + 1, byte(d.opposite())}
	case up.opposite():
		return side{s[0], s[1], s[2] - 1, byte(d.opposite())}
	case fwd:
		return side{s[0], s[1] + 1, s[2], byte(d.opposite())}
	case fwd.opposite():
		return side{s[0], s[1] - 1, s[2], byte(d.opposite())}
	case rgt:
		return side{s[0] + 1, s[1], s[2], byte(d.opposite())}
	case rgt.opposite():
		return side{s[0] - 1, s[1], s[2], byte(d.opposite())}
	}
	return s
}

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 4314
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 0
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
			want: 0, //10,
		},
		{
			in:   inputTest1,
			want: 0, //58,
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
	cs, ss := scan(r), map[side]struct{}{}
	for _, c := range cs {
		for _, d := range []dir{up, up.opposite(), fwd, fwd.opposite(), rgt, rgt.opposite()} {
			s := side{c[0], c[1], c[2], byte(d)}
			if _, ok := ss[s.opposite()]; ok {
				delete(ss, s.opposite())
				continue
			}
			ss[s] = struct{}{}
		}
	}
	return len(ss)
}

func PartTwo(r io.Reader) int {
	return 0
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
