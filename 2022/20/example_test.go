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

type ring[T any] []struct {
	pi, ni int
	pl     T
}

func new[T any](vs []T) (c ring[T]) {
	c = make(ring[T], len(vs))
	for i, n := range vs {
		c[i] = struct {
			pi, ni int
			pl     T
		}{i - 1, i + 1, n}
	}
	c[0].pi, c[len(vs)-1].ni = len(vs)-1, 0
	return c
}

func (r *ring[T]) move(i, n int) {
	if n%len(*r) == 0 {
		return
	}
	j, _ := r.find(i, n)
	el, el2 := (*r)[i], (*r)[j]
	(*r)[el.pi].ni, (*r)[el.ni].pi = el.ni, el.pi
	(*r)[i].pi, (*r)[i].ni = j, el2.ni
	(*r)[j].ni, (*r)[el2.ni].pi = i, i
}

func (r *ring[T]) find(i, n int) (int, T) {
	el := (*r)[i]
	for j := n % len(*r); j != 0; j-- {
		i, el = el.ni, (*r)[el.ni]
	}
	return i, el.pl
}

func ExamplePartOne() {
	ns := PartOne(strings.NewReader(input))
	fmt.Println(sum(ns[:]...))
	// Output:
	// 4914
}

func ExamplePartTwo() {
	ns := PartTwo(strings.NewReader(input))
	fmt.Println(sum(ns[:]...))
	// Output:
	// 7973051839072
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
	want := [...]int{811589153, 2434767459, -1623178306}
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func PartOne(r io.Reader) (rs [3]int) {
	return decode(scan(r), 1, 1, 1e3)
}

func PartTwo(r io.Reader) (rs [3]int) {
	return decode(scan(r), 811589153, 10, 1e3)
}

func decode(ns []int, k, n, m int) (rs [3]int) {
	var idx int
	rg := new(conv(ns, func(i, l, n int) [2]int {
		if n == 0 {
			idx = i
		}
		if n *= k; n >= 0 {
			return [2]int{n, n % (l - 1)}
		}
		return [2]int{n, n%(l-1) + l - 1}
	}))
	for i := 0; i < n; i++ {
		for j := range rg {
			_, v := rg.find(j, 0)
			rg.move(j, v[1])
		}
	}
	for i, j := 0, idx; i < len(rs); i++ {
		var v [2]int
		j, v = rg.find(j, m)
		rs[i] = v[0]
	}
	return rs
}

func conv[T1, T2 any](vs []T1, f func(int, int, T1) T2) (rs []T2) {
	rs = make([]T2, len(vs))
	for i, v := range vs {
		rs[i] = f(i, len(vs), v)
	}
	return rs
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
