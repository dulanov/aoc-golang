// https://adventofcode.com/2022/day/25
package d25_test

import (
	"bufio"
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
	// 20-==01-2-=1-2---1-0
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := "2=-1=0"
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func PartOne(r io.Reader) string {
	var n int
	for _, s := range scan(r) {
		n += snf2Dec(s)
	}
	return dec2Snf(n)
}

func dec2Snf(n int) (s string) {
	for ; n != 0; n = (n + 2) / 5 {
		s = string("012=-"[n%5]) + s
	}
	return s
}

func snf2Dec(s string) (n int) {
	for _, r := range s {
		n = 5*n + 2 - strings.IndexRune("210-=", r)
	}
	return n
}

func scan(r io.Reader) (ss []string) {
	for s := bufio.NewScanner(r); s.Scan(); {
		ss = append(ss, s.Text())
	}
	return ss
}

const inputTest = `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`
