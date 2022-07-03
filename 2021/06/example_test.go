// https://adventofcode.com/2021/day/06
package d06_test

import (
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

const (
	DaysFirstCycle = 8
	DaysNextCycles = 6
)

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input), 80))
	// Output:
	// 360268
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input), 256))
	// Output:
	// 1632146183902
}

func TestPartOne(t *testing.T) {
	for i, tt := range []struct {
		days, want int
	}{
		{18, 26},
		{80, 5934},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := PartOne(strings.NewReader(input_test), tt.days)
			if got != tt.want {
				t.Errorf("got %d; want %d", got, tt.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(input_test), 256)
	want := 26984457539
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader, days int) int {
	return sim(scan(r), days)
}

func PartTwo(r io.ReadSeeker, days int) int {
	return sim(scan(r), days)
}

func sim(ns []int, days int) int {
	var ps [DaysFirstCycle + 1]int
	for _, n := range ns {
		ps[n]++
	}
	for i := 0; i < days; i++ {
		n := ps[0]
		copy(ps[:len(ps)-1], ps[1:])
		ps[DaysNextCycles] += n
		ps[DaysFirstCycle] = n
	}
	return sum(ps[:])
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

const input_test = `3,4,3,1,2`
