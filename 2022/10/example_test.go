// https://adventofcode.com/2022/day/10
package d10_test

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

type inr string

const (
	nop inr = "noop"
	adx inr = "addx"
)

type ir struct {
	inr inr
	arg int
}

func (o ir) execute(pc, rx int) (int, int) {
	switch o.inr {
	case nop:
		return pc + 1, rx
	case adx:
		return pc + 2, rx + o.arg
	}
	return pc, rx
}

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 13720
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// ####.###..#..#.###..#..#.####..##..#..#.
	// #....#..#.#..#.#..#.#..#....#.#..#.#..#.
	// ###..###..#..#.#..#.####...#..#....####.
	// #....#..#.#..#.###..#..#..#...#....#..#.
	// #....#..#.#..#.#.#..#..#.#....#..#.#..#.
	// #....###...##..#..#.#..#.####..##..#..#.
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := 13140
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest))
	want := `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`
	if got != want {
		t.Errorf("got %s; want %s", got, want)
	}
}

func PartOne(r io.Reader) (n int) {
	lm := 20
	exec(scan(r), func(pc, pn, rx int) {
		if pc <= lm && pc+pn > lm {
			n, lm = n+lm*rx, lm+40
		}
	})
	return n
}

func PartTwo(r io.Reader) string {
	bh, bw := 6, 40
	bs := append([]byte{}, strings.Repeat(".", bh*bw)...)
	exec(scan(r), func(pc, pn, rx int) {
		for i := 0; i < pn; i++ {
			if abs((pc+i-1)%bw-rx) <= 1 {
				/* sprite collision: 3 pixels */
				bs[pc+i-1] = '#'
			}
		}
	})
	return strings.Join(splitBy(bs, bw, func(bs []byte) string {
		return string(bs)
	}), "\n")
}

func exec(vs []ir, fn func(int, int, int)) {
	pc, rx := 1, 1
	for _, ir := range vs {
		pc2, rx2 := ir.execute(pc, rx)
		fn(pc, pc2-pc, rx)
		pc, rx = pc2, rx2
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func splitBy[T1, T2 any](vs []T1, n int, fn func([]T1) T2) (rs []T2) {
	for n < len(vs) {
		vs, rs = vs[n:], append(rs, fn(vs[0:n:n]))
	}
	return append(rs, fn(vs))
}

func scan(r io.Reader) (rs []ir) {
	for s := bufio.NewScanner(r); s.Scan(); {
		var ir ir
		fmt.Sscanf(s.Text(), "%s %d", &ir.inr, &ir.arg)
		rs = append(rs, ir)
	}
	return rs
}

const inputTest = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`
