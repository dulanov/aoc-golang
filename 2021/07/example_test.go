// https://adventofcode.com/2021/day/7
package d07_test

import (
	_ "embed"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 344297
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 97164301
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := 37
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest))
	want := 168
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (cs int) {
	ns := scan(r)
	sort.Ints(ns)
	return calc(ns, ns[len(ns)>>1], func(n int) int {
		return n
	})
}

func PartTwo(r io.Reader) (cs int) {
	ns := scan(r)
	return calc(ns, sum(ns)/len(ns), func(n int) int {
		return (n * (n + 1)) >> 1
	})
}

func calc(ns []int, pos int, fn func(int) int) (cs int) {
	ch := make(chan int)
	go func() { ch <- cost(ns, pos, fn) }()
	go func() { ch <- cost(ns, pos+1, fn) }()
	return min(<-ch, <-ch)
}

func cost(ns []int, pos int, fn func(int) int) (cs int) {
	for _, n := range ns {
		cs += fn(abs(n - pos))
	}
	return cs
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func sum(ns []int) (rs int) {
	for _, n := range ns {
		rs += n
	}
	return rs
}

func scan(r io.Reader) (ns []int) {
	b, _ := io.ReadAll(r)
	for _, s := range strings.Split(string(b), ",") {
		n, _ := strconv.Atoi(s)
		ns = append(ns, n)
	}
	return ns
}

const inputTest = `16,1,2,0,4,2,7,1,2,14`
