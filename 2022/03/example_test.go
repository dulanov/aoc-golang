// https://adventofcode.com/2022/day/03
package d03_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"math"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 7997
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 2545
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := 157
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest))
	want := 70
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (n int) {
	for _, ns := range scan(r) {
		n += int(math.Log2(float64(
			and(mask(ns[:len(ns)/2]), mask(ns[len(ns)/2:])))))
	}
	return n
}

func PartTwo(r io.Reader) (n int) {
	rs := scan(r)
	for i := 0; i < len(rs); i += 3 {
		n += int(math.Log2(float64(
			and(mask(rs[i]), mask(rs[i+1]), mask(rs[i+2])))))
	}
	return n
}

func and(ms ...uint64) (m uint64) {
	m = ms[0]
	for _, m1 := range ms[1:] {
		m &= m1
	}
	return m
}

func mask(ns []int) (m uint64) {
	for _, n := range ns {
		m |= 1 << n
	}
	return m
}

func scan(r io.Reader) (rs [][]int) {
	for s := bufio.NewScanner(r); s.Scan(); {
		bs := s.Bytes()
		ns := make([]int, len(bs))
		for i, b := range bs {
			if ns[i] = int(b) - 'a' + 1; ns[i] < 0 {
				ns[i] += 58
			}
		}
		rs = append(rs, ns)
	}
	return rs
}

const inputTest = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
