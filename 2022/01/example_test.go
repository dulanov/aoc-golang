// https://adventofcode.com/2022/day/01
package d01_test

import (
	"bufio"
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
	// 67016
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 200116
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
	want := 45000
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) int {
	ns := flt(scan(r), sum)
	sort.Sort(sort.Reverse(sort.IntSlice(ns)))
	return ns[0]
}

func PartTwo(r io.Reader) int {
	ns := flt(scan(r), sum)
	sort.Sort(sort.Reverse(sort.IntSlice(ns)))
	return sum(ns[:3])
}

func flt(ls [][]int, fn func([]int) int) (ns []int) {
	ns = make([]int, len(ls))
	for _, l := range ls {
		ns = append(ns, fn(l))
	}
	return ns
}

func sum(ns []int) (n int) {
	for _, v := range ns {
		n += v
	}
	return n
}

func scan(r io.Reader) (items [][]int) {
	for i, s := 0, bufio.NewScanner(r); s.Scan(); {
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
