// https://adventofcode.com/2022/day/04
package d04_test

import (
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
	// 424
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 804
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := 2
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest))
	want := 4
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (n int) {
	for _, pr := range scan(r) {
		if (pr[0] <= pr[2] && pr[1] >= pr[3]) ||
			(pr[0] >= pr[2] && pr[1] <= pr[3]) {
			n++
		}
	}
	return n
}

func PartTwo(r io.Reader) (n int) {
	for _, pr := range scan(r) {
		if pr[0] <= pr[3] && pr[1] >= pr[2] {
			n++
		}
	}
	return n
}

func scan(r io.Reader) (ps [][4]int) {
	for {
		var n1, n2, n3, n4 int
		if _, err := fmt.Fscanf(r, "%d-%d,%d-%d\n", &n1, &n2, &n3, &n4); err == io.EOF {
			break
		}
		ps = append(ps, [4]int{n1, n2, n3, n4})
	}
	return ps
}

const inputTest = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
