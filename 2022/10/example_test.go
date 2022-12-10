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
	lm, pc, rx := 20, 1, 1
	for _, ir := range scan(r) {
		pc2, rx2 := ir.execute(pc, rx)
		if pc <= lm && pc2 > lm {
			n, lm = n+lm*rx, lm+40
		}
		pc, rx = pc2, rx2
	}
	return n
}

func PartTwo(r io.Reader) string {
	bh, bw, pc, rx := 6, 40, 1, 1
	bs := append([]byte{}, strings.Repeat(".", bh*bw)...)
	for _, ir := range scan(r) {
		pc2, rx2 := ir.execute(pc, rx)
		for i := 0; i < pc2-pc; i++ {
			if n := (pc + i - 1) % bw; n >= rx-1 && n <= rx+1 {
				bs[pc+i-1] = '#'
			}
		}
		pc, rx = pc2, rx2
	}
	return strings.Join(splitBy(bs, bw, func(bs []byte) string {
		return string(bs)
	}), "\n")
}

func splitBy[T1, T2 any](vs []T1, n int, fn func([]T1) T2) (rs []T2) {
	for n < len(vs) {
		vs, rs = vs[n:], append(rs, fn(vs[0:n:n]))
	}
	return append(rs, fn(vs))
}

func scan(r io.Reader) (irs []ir) {
	for s := bufio.NewScanner(r); s.Scan(); {
		var ir ir
		fmt.Sscanf(s.Text(), "%s %d", &ir.inr, &ir.arg)
		irs = append(irs, ir)
	}
	return irs
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
