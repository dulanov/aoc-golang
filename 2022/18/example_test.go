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
	got := PartTwo(strings.NewReader(inputTest0))
	want := 0
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (n int) {
	cs := scan(r)
	for _, c1 := range cs {
		for _, c2 := range cs {
			if abs(c1[0]-c2[0])+abs(c1[1]-c2[1])+abs(c1[2]-c2[2]) == 1 {
				n++
			}
		}
	}
	return len(cs)*6 - n
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

func scan(r io.Reader) (cs [][3]int) {
	for s := bufio.NewScanner(r); s.Scan(); {
		var x, y, z int
		fmt.Sscanf(s.Text(), "%d,%d,%d", &x, &y, &z)
		cs = append(cs, [3]int{x, y, z})
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
