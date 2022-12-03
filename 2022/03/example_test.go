// https://adventofcode.com/2022/day/03
package d03_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"sort"
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
		ns1, ns2 := ns[:len(ns)/2], ns[len(ns)/2:]
		sort.Ints(ns1)
		sort.Ints(ns2)
		n += intersect(ns1, ns2)[0]
	}
	return n
}

func PartTwo(r io.Reader) (n int) {
	rs := scan(r)
	for i := 0; i < len(rs); i += 3 {
		ns1, ns2, ns3 := rs[i], rs[i+1], rs[i+2]
		sort.Ints(ns1)
		sort.Ints(ns2)
		sort.Ints(ns3)
		n += intersect(intersect(ns1, ns2), ns3)[0]
	}
	return n
}

func intersect(ns1, ns2 []int) (ns []int) {
	for i1, i2 := 0, 0; i1 < len(ns1) && i2 < len(ns2); {
		if ns1[i1] == ns2[i2] {
			ns, i1, i2 = append(ns, ns1[i1]), i1+1, i2+1
			continue
		}
		if ns1[i1] < ns2[i2] {
			i1++
		} else {
			i2++
		}
	}
	return ns
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
