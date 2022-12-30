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

type cycle[T any] []struct {
	pi, ni int
	pl     T
}

func new[T any](vs []T) (c cycle[T]) {
	c = make(cycle[T], len(vs))
	for i, n := range vs {
		c[i] = struct {
			pi, ni int
			pl     T
		}{i - 1, i + 1, n}
	}
	c[0].pi, c[len(vs)-1].ni = len(vs)-1, 0
	return c
}

func (c *cycle[T]) move(i, n int) {
	el := (*c)[i]
	(*c)[el.pi].ni, (*c)[el.ni].pi = el.ni, el.pi
	for j := 0; j < n; j++ {
		el = (*c)[el.ni]
	}
	j := (*c)[el.ni].pi
	(*c)[i].pi, (*c)[i].ni = j, el.ni
	(*c)[j].ni, (*c)[el.ni].pi = i, i
}

func (c *cycle[T]) value(i, n int) (int, T) {
	el := (*c)[i]
	for j := 0; j < n; j++ {
		el = (*c)[el.ni]
	}
	return (*c)[el.ni].pi, el.pl
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
	var idx int
	cy := new(conv(scan(r), func(i, l, n int) [2]int {
		if n == 0 {
			idx = i
		}
		if n >= 0 {
			return [2]int{n, n % (l - 1)}
		}
		return [2]int{n, n%(l-1) + l - 1}
	}))
	for i := range cy {
		_, v := cy.value(i, 0)
		cy.move(i, v[1])
	}
	for i, j := 0, idx; i < 3; i++ {
		var v [2]int
		j, v = cy.value(j, 1e3)
		rs[i] = v[0]
	}
	return rs
}

func PartTwo(r io.Reader) (rs [3]int) {
	var idx int
	cy := new(conv(scan(r), func(i, l, n int) [2]int {
		if n == 0 {
			idx = i
		}
		if n *= 811589153; n >= 0 {
			return [2]int{n, n % (l - 1)}
		}
		return [2]int{n, n%(l-1) + l - 1}
	}))
	for i := 0; i < 10; i++ {
		for j := range cy {
			_, v := cy.value(j, 0)
			cy.move(j, v[1])
		}
	}
	for i, j := 0, idx; i < 3; i++ {
		var v [2]int
		j, v = cy.value(j, 1e3)
		rs[i] = v[0]
	}
	return rs
}

func conv[T1, T2 any](vs []T1, f func(int, int, T1) T2) (rs []T2) {
	for i, v := range vs {
		rs = append(rs, f(i, len(vs), v))
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
