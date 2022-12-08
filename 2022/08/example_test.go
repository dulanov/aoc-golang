// https://adventofcode.com/2022/day/08
package d08_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"math"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 1533
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 345744
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := 21
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest))
	want := 8
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (n int) {
	hs := scan(r)
	vs := make([]bool, len(hs))
	ln := int(math.Sqrt(float64(len(hs))))
	calc(hs, ln, func(h, w int) {
		vs[h*ln+w] = true
	}, func() {
		rev(vs)
	})
	return sum(vs) + (ln-1)*4
}

func PartTwo(r io.Reader) (n int) {
	hs := scan(r)
	ln := int(math.Sqrt(float64(len(hs))))
	calc(hs, ln, func(h, w int) {
		if n1 := score(hs, ln, h, w); n < n1 {
			n = n1
		}
	}, func() {})
	return n
}

func calc(hs []int, ln int, fn func(h, w int), frev func()) {
	for pass := 0; pass < 2; /* frw & bck */ pass++ {
		for i, j := 1, ln; i < ln-1; i, j = i+1, j+ln {
			for k := 1; k < ln-1; k++ {
				trg := false
				if hs[j+k] > hs[j] {
					trg, hs[j] = true, hs[j+k]
				}
				if hs[j+k] > hs[k] {
					trg, hs[k] = true, hs[j+k]
				}
				if trg {
					fn(i, k)
				}
			}
		}
		rev(hs)
		frev()
	}
}

func score(hs []int, ln, h, w int) (s int) {
	s = 1
	n, ps, vl := 1, h*ln+w, hs[h*ln+w]
	for _, l := range []struct{ dt, nm int }{
		{-ln, h - 1}, {ln, ln - h - 2}, {-1, w - 1}, {1, ln - w - 2}} {
		for i, j := 0, ps+l.dt; i < l.nm && vl > hs[j]; i, j = i+1, j+l.dt {
			n++
		}
		n, s = 1, s*n
	}
	return s
}

func sum(vs []bool) (n int) {
	for _, v := range vs {
		if v {
			n++
		}
	}
	return n
}

func rev[T any](vs []T) {
	for i, j := 0, len(vs)-1; i < j; i, j = i+1, j-1 {
		vs[i], vs[j] = vs[j], vs[i]
	}
}

func scan(r io.Reader) (hs []int) {
	for s := bufio.NewScanner(r); s.Scan(); {
		for _, r := range s.Text() {
			hs = append(hs, int(r-'0'))
		}
	}
	return hs
}

const inputTest = `30373
25512
65332
33549
35390`
