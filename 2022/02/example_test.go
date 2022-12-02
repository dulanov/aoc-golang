// https://adventofcode.com/2022/day/02
package d02_test

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

type shape int
type game int

const (
	rock shape = iota
	paper
	scissors
)

const (
	lost game = iota
	draw
	won
)

func (s shape) defeats(s2 shape) bool {
	return ((int(s) - int(s2) + 2) % 3) == 0
}

func (s shape) score() int {
	return int(s) + 1
}

func (s shape) adjust(g game) shape {
	for _, s2 := range []shape{rock, paper, scissors} {
		if s2.play(s) == g {
			return s2
		}
	}
	return s
}

func (s shape) play(s2 shape) game {
	if s == s2 {
		return draw
	}
	if s.defeats(s2) {
		return won
	}
	return lost
}

func (g game) score() int {
	return int(g) * 3
}

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 13675
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 14184
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := 15
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest))
	want := 12
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (n int) {
	for _, m := range scan(r) {
		n += m.s2.score() + m.s2.play(m.s1).score()
	}
	return n
}

func PartTwo(r io.Reader) (n int) {
	for _, m := range scan(r) {
		n += m.s1.adjust(m.g).score() + m.g.score()
	}
	return n
}

func scan(r io.Reader) (guide []struct {
	s1, s2 shape
	g      game
}) {
	for s := bufio.NewScanner(r); s.Scan(); {
		fs := strings.Fields(s.Text())
		guide = append(guide, struct {
			s1, s2 shape
			g      game
		}{(shape)(fs[0][0] - 'A'), (shape)(fs[1][0] - 'X'), (game)(fs[1][0] - 'X')})
	}
	return guide
}

const inputTest = `A Y
B X
C Z`
