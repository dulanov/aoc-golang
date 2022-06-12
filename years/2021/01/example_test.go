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

//go:embed testdata/input
var input string

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 1342
}

func TestPartOne(t *testing.T) {
	const in = "199\n200\n208\n210\n200\n207\n240\n269\n260\n263"
	got := PartOne(strings.NewReader(in))
	want := 7
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func PartOne(r io.Reader) int {
	return numberOfTimes(r, func(i, j int) bool {
		return j > i
	})
}

func numberOfTimes(r io.Reader, fn func(int, int) bool) (n int) {
	vs := conv(split(r))
	for i := 1; i < len(vs); i++ {
		if fn(vs[i-1], vs[i]) {
			n++
		}
	}
	return n
}

func split(r io.Reader) (vs []string) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		vs = append(vs, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	return vs
}

func conv(vs []string) []int {
	ns := make([]int, len(vs))
	for i, s := range vs {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		ns[i] = n
	}
	return ns
}
