// https://adventofcode.com/2021/day/11
package d11_test

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

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 1546
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 0
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(input_test))
	want := 1656
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(input_test))
	want := 0
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (n int) {
	rs := scan(r)
	for i := 0; i < 100; i++ {
		n += step(rs)
	}
	return n
}

func PartTwo(r io.Reader) int {
	return 0
}

func step(rs [][]uint8) (n int) {
	var q []struct{ x, y int }
	for i, r := range rs {
		for j := range r {
			if r[j]++; r[j] == 10 {
				r[j], q = 0, append(q, struct{ x, y int }{j, i})
			}
		}
	}
	for ; len(q) != 0; n++ {
		x, y := q[len(q)-1].x, q[len(q)-1].y
		q = q[:len(q)-1]
		for _, d := range []struct{ dx, dy int }{
			{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
			if x+d.dx < 0 || x+d.dx >= len(rs[y]) || y+d.dy < 0 || y+d.dy >= len(rs) {
				continue
			}
			if rs[y+d.dy][x+d.dx] == 0 {
				continue
			}
			if rs[y+d.dy][x+d.dx]++; rs[y+d.dy][x+d.dx] == 10 {
				rs[y+d.dy][x+d.dx], q = 0, append(q, struct{ x, y int }{x + d.dx, y + d.dy})
			}
		}
	}
	return n
}

func scan(r io.Reader) (rs [][]uint8) {
	for s := bufio.NewScanner(r); s.Scan(); {
		ns := make([]uint8, len(s.Text()))
		for i, r := range s.Text() {
			ns[i] = uint8(r - '0')
		}
		rs = append(rs, ns)
	}
	return rs
}

const input_test = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`
