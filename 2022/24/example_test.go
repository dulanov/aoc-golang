// https://adventofcode.com/2022/day/24
package d24_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

type queue[T any] []T

func (q queue[T]) empty() bool {
	return len(q) == 0
}

func (q queue[T]) enq(v ...T) queue[T] {
	return append(q, v...)
}

func (q queue[T]) deq() (queue[T], T, bool) {
	if q.empty() {
		return q, *new(T), false
	}
	return q[1:], q[0], true
}

type dir byte

func (d dir) blow(p int, h, w int) int {
	switch d {
	case up:
		return (p - w + h*w) % (h * w)
	case dn:
		return (p + w) % (h * w)
	case rt:
		return (p/w)*w + (p+1)%w
	case lt:
		return (p/w)*w + (p-1+w)%w
	}
	return p
}

func (d dir) move(p int, h, w int) (int, bool) {
	switch d {
	case nn:
	case up:
		p -= w
	case dn:
		p += w
	case rt:
		if (p+1)%w == 0 {
			return p, false
		}
		p++
	case lt:
		if p%w == 0 {
			return p, false
		}
		p--
	}
	return p, p >= 0 && p < h*w
}

const (
	nn dir = '.'
	up dir = '^'
	dn dir = 'v'
	rt dir = '>'
	lt dir = '<'
)

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 277
}

func ExamplePartTwo() {
	n1, n2, n3 := PartTwo(strings.NewReader(input))
	fmt.Println(n1 + n2 + n3)
	// Output:
	// 877
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := 18
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	n1, n2, n3 := PartTwo(strings.NewReader(inputTest))
	want := []int{18, 23, 13}
	if !reflect.DeepEqual([]int{n1, n2, n3}, want) {
		t.Errorf("got %v; want %v", []int{n1, n2, n3}, want)
	}
}

func PartOne(r io.Reader) (n int) {
	w, bs := scan(r)
	p, h := -w, len(bs)/w
	return steps(p, len(bs)-1, 0, h, w, bs) + 1
}

func PartTwo(r io.Reader) (n1, n2, n3 int) {
	w, bs := scan(r)
	p1, p2, h := -w, len(bs)+w-1, len(bs)/w
	n1 = steps(p1, len(bs)-1, 0, h, w, bs) + 1
	n2 = steps(p2, 0, n1, h, w, bs) + 1
	n3 = steps(p1, len(bs)-1, n1+n2, h, w, bs) + 1
	return n1, n2, n3
}

func steps(fr, to, of, h, w int, bs []dir) (n int) {
	fs, vs := make([][]bool, lcm(len(bs)/w, w)), make(map[[2]int]bool)
	for i := range fs {
		fs[i] = make([]bool, len(bs))
	}
	for i, d := range bs {
		if d == nn {
			continue
		}
		for j, p := 0, i; j < len(fs); j++ {
			fs[j][p], p = true, d.blow(p, h, w)
		}
	}
	bfs(fr, func(p int) bool {
		return p == to
	}, func(p int) (ps []int) {
		if p == fr {
			ps = append(ps, p)
		}
		for _, d := range []dir{nn, dn, rt, lt, up} {
			if p, ok := d.move(p, h, w); ok && !vs[[2]int{n, p}] && !fs[(n+of+1)%len(fs)][p] {
				ps, vs[[2]int{n, p}] = append(ps, p), true
			}
		}
		return ps
	}, func() {
		n++
	})
	return n
}

func bfs[T comparable](s T, ck func(s T) bool, gn func(s T) []T, ll func()) {
	for q, n := (queue[T]{s}), 1; !q.empty(); n-- {
		if n == 0 {
			n = len(q)
			ll()
		}
		var s T
		q, s, _ = q.deq()
		if ck(s) {
			return
		}
		q = q.enq(gn(s)...)
	}
}

func lcm(n1, n2 int) int {
	return n1 * n2 / gcd(n1, n2)
}

func gcd(n1, n2 int) int {
	for n2 != 0 {
		n1, n2 = n2, n1%n2
	}
	return n1
}

func scan(r io.Reader) (w int, bs []dir) {
	for s := bufio.NewScanner(r); s.Scan(); {
		if w == 0 {
			w = len(s.Text()) - 2
			continue
		}
		if s.Text()[1] == '#' {
			continue
		}
		bs = append(bs, []dir(s.Text())[1:w+1]...)
	}
	return w, bs
}

const inputTest = `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`
