// https://adventofcode.com/2021/day/1
package d01_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"testing"
)

const input_test = `199
200
208
210
200
207
240
269
260
263`

//go:embed testdata/input
var input string

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 1342
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 1378
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(input_test))
	want := 7
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(input_test))
	want := 5
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func PartOne(r io.Reader) int {
	return numberOfTimes(split(r), 1, func(i, j int) bool {
		return j > i
	})
}

func PartTwo(r io.Reader) int {
	return numberOfTimes(split(r), 3, func(i, j int) bool {
		return j > i
	})
}

func numberOfTimes(vs []int, w int, fn func(int, int) bool) (n int) {
	for i := w; i < len(vs); i++ {
		if fn(vs[i-w], vs[i]) {
			n++
		}
	}
	return n
}

func split(r io.Reader) (vs []int) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		n, err := strconv.Atoi(sc.Text())
		if err != nil {
			log.Fatal(err)
		}
		vs = append(vs, n)
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	return vs
}
