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
	// 0
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
	want := 0
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (n int) {
	for _, m := range scan(r) {
		n += m[1].score() + m[1].play(m[0]).score()
	}
	return n
}

func PartTwo(r io.Reader) int {
	return 0
}

func scan(r io.Reader) (guide [][2]shape) {
	for s := bufio.NewScanner(r); s.Scan(); {
		fs := strings.Fields(s.Text())
		var s1, s2 shape
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
			s2 = rock
		case "Y":
			s2 = paper
		case "Z":
			s2 = scissors
		}
		guide = append(guide, [2]shape{s1, s2})
	}
	return guide
}

const inputTest = `A Y
B X
C Z`
