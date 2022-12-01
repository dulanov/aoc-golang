// https://adventofcode.com/2022/day/01
package d01_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 67016
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 0
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := 24000
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
	for _, cs := range scan(r) {
		if s := sum(cs); n < s {
			n = s
		}
	}
	return n
}

func PartTwo(r io.Reader) int {
	return 0
}

func sum(ns []int) (n int) {
	for _, v := range ns {
		n += v
	}
	return n
}

func scan(r io.Reader) (items [][]int) {
	for s, i := bufio.NewScanner(r), 0; s.Scan(); {
		if len(s.Text()) == 0 {
			i++
			continue
		}
		if n, _ := strconv.Atoi(s.Text()); len(items) == i {
			items = append(items, []int{n})
		} else {
			items[i] = append(items[i], n)
		}
	}
	return items
}

const inputTest = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
