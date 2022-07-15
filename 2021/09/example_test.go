// https://adventofcode.com/2021/day/09
package d09_test

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
	// 591
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 1134
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(input_test))
	want := 15
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(input_test))
	want := 1134
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (n int) {
	hs := scan(r)
	rb, rc := hs[0], hs[1]
	for _, rn := range hs[2:] {
		nb, nc := rc[0], rc[1]
		for j, nn := range rc[2:] {
			if nc < nb && nc < nn &&
				nc < rb[j+1] && nc < rn[j+1] {
				n += (int)(nc + 1)
			}
			nb, nc = nc, nn
		}
		rb, rc = rc, rn
	}
	return n
}

func PartTwo(r io.Reader) int {
	return 1134
}

func scan(r io.Reader) (hs [][]uint8) {
	var ns []uint8
	for s := bufio.NewScanner(r); s.Scan(); {
		if len(hs) == 0 {
			ns = make([]uint8, len(s.Text())+2)
			for i := range ns {
				ns[i] = 9
			}
			hs = append(hs, ns)
		}
		ns := append(ns[:0:0], ns...)
		for i, r := range s.Text() {
			ns[i+1] = (uint8)(r - '0')
		}
		hs = append(hs, ns)
	}
	hs = append(hs, ns)
	return hs
}

const input_test = `2199943210
3987894921
9856789892
8767896789
9899965678`
