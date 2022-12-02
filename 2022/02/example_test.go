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
	scissors
	paper

	lost game = iota
	draw
	won
)

func (s shape) defeats(s2 shape) bool {
	if (s == rock && s2 == scissors) ||
		(s == scissors && s2 == paper) ||
		(s == paper && s2 == rock) {
		return true
	}
	return false
}

func (s shape) score() int {
	switch s {
	case rock:
		return 1
	case scissors:
		return 3
	case paper:
		return 2
	}
	return 0
}

func (s shape) adjust(g game) shape {
	if g == draw {
		return s
	}
	for _, s2 := range []shape{rock, paper, scissors} {
		if s != s2 && s2.play(s) == g {
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
	if g == draw {
		return 3
	}
	if g == won {
		return 6
	}
	return 0 /* lost */
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
		var s1, s2 shape
		var g game
		switch fs[0] {
		case "A":
			s1 = rock
		case "B":
			s1 = paper
		case "C":
			s1 = scissors
		}
		switch fs[1] {
		case "X":
			s2, g = rock, lost
		case "Y":
			s2, g = paper, draw
		case "Z":
			s2, g = scissors, won
		}
		guide = append(guide, struct {
			s1, s2 shape
			g      game
		}{s1, s2, g})
	}
	return guide
}

const inputTest = `A Y
B X
C Z`
