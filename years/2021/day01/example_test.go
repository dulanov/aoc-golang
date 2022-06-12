// https://adventofcode.com/2021/day/1
package day01_test

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

func Example() {
	fmt.Println(numberOfTimes(strings.NewReader(input), byAsc))
	// Output:
	// 1342
}

func TestNumberOfTimes(t *testing.T) {
	const in = "199\n200\n208\n210\n200\n207\n240\n269\n260\n263"
	got := numberOfTimes(strings.NewReader(in), byAsc)
	want := 7
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func numberOfTimes(r io.Reader, fn func(int, int) bool) int {
	var n int
	vs := conv(split(r))
	if len(vs) <= 1 {
		return n
	}
	for i := 1; i < len(vs); i++ {
		if fn(vs[i-1], vs[i]) {
			n++
		}
	}
	return n
}

func split(r io.Reader) []string {
	var vs []string
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
	rs := make([]int, len(vs))
	for i, s := range vs {
		j, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		rs[i] = j
	}
	return rs
}

func byAsc(i, j int) bool {
	return j > i
}
