// https://adventofcode.com/2022/day/20
package d20_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"

	"golang.org/x/exp/constraints"
)

//go:embed testdata/input
var input string

func ExamplePartOne() {
	ns := PartOne(strings.NewReader(input))
	fmt.Println(sum(ns[:]...))
	// Output:
	// 4914
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 0
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := [...]int{4, -3, 2}
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest))
	want := 0
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (rs [3]int) {
	ns := scan(r)
	type item struct {
		vl, pi, ni int
	}
	vi, vs := 0, make([]item, len(ns))
	for i, n := range ns {
		vs[i] = item{n, i - 1, i + 1}
	}
	vs[0].pi, vs[len(ns)-1].ni = len(ns)-1, 0
	for i, v := range vs {
		if v.vl == 0 {
			vi = i
			continue
		}
		if v.vl > len(ns) {
			v.vl = v.vl % (len(ns) - 1)
		}
		if v.vl < 0 {
			v.vl += len(ns) - 1
		}
		vs[v.pi].ni, vs[v.ni].pi = v.ni, v.pi
		for j := v.vl; j != 0; j-- {
			v = vs[v.ni]
		}
		j := vs[v.ni].pi
		vs[i].pi, vs[i].ni = j, v.ni
		vs[j].ni, vs[v.ni].pi = i, i
	}
	for i, v := 0, vs[vi]; i <= 3000; i, v = i+1, vs[v.ni] {
		if i != 0 && i%1000 == 0 {
			rs[i/1000-1] = v.vl
		}
	}
	return rs
}

func PartTwo(r io.Reader) int {
	return 0
}

func sum[T constraints.Integer](ns ...T) (n T) {
	for _, v := range ns {
		n += v
	}
	return n
}

func scan(r io.Reader) (ns []int) {
	for s := bufio.NewScanner(r); s.Scan(); {
		n, _ := strconv.Atoi(s.Text())
		ns = append(ns, n)
	}
	return ns
}

const inputTest = `1
2
-3
3
-2
0
4`
