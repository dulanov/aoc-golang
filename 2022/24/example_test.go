// https://adventofcode.com/2022/day/24
package d24_test

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

func (d dir) blow(p, n, h, w int) int {
	switch l := h * w; d {
	case up:
		return (p - (w*n)%l + l) % l
	case dn:
		return (p + w*n) % l
	case rt:
		return (p/w)*w + (p+n)%w
	case lt:
		return (p/w)*w + (p-n%w+w)%w
	}
	return p
}

func (d dir) move(p, h, w int) (int, bool) {
	switch d {
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

func (d dir) opposite() dir {
	switch d {
	case up:
		return dn
	case dn:
		return up
	case rt:
		return lt
	case lt:
		return rt
	}
	return nn
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
	want := [3]int{18, 23, 13}
	if [3]int{n1, n2, n3} != want {
		t.Errorf("got %v; want %v", [3]int{n1, n2, n3}, want)
	}
}

func PartOne(r io.Reader) (n int) {
	h, w, bs := scan(r)
	return steps(-w, len(bs)-1, 0, h, w, bs) + 1
}

func PartTwo(r io.Reader) (n1, n2, n3 int) {
	h, w, bs := scan(r)
	n1 = steps(-w, len(bs)-1, 0, h, w, bs) + 1
	n2 = steps(len(bs)+w-1, 0, n1, h, w, bs) + 1
	n3 = steps(-w, len(bs)-1, n1+n2, h, w, bs) + 1
	return n1, n2, n3
}

func steps(p1, p2, m, h, w int, bs []dir) (n int) {
	vs := make(map[int]bool)
	bfs(p1, func(p int) bool {
		return p == p2
	}, func() {
		n, vs = n+1, make(map[int]bool)
	}, func(p int) (ps []int) {
		if p == p1 {
			ps = append(ps, p)
		}
	outer:
		for _, d := range []dir{nn, up, dn, rt, lt} {
			if p, ok := d.move(p, h, w); ok && !vs[p] {
				vs[p] = true
				for _, d2 := range []dir{up, dn, rt, lt} {
					if p2 := d2.opposite().blow(p, n+m+1, h, w); bs[p2] == d2 {
						continue outer
					}
				}
				ps = append(ps, p)
			}
		}
		return ps
	})
	return n
}

func bfs[T any](s T, ck func(s T) bool, ll func(), gn func(s T) []T) {
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

func scan(r io.Reader) (h, w int, bs []dir) {
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
	return len(bs) / w, w, bs
}

const inputTest = `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`
