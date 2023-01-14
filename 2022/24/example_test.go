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
	fmt.Println(PartTwo(strings.NewReader(input)))
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
	got := PartTwo(strings.NewReader(inputTest))
	want := 54
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (n int) {
	w, bs := scan(r)
	return steps(-w, len(bs)-1, 0, w, bs)
}

func PartTwo(r io.Reader) (n int) {
	w, bs := scan(r)
	return steps(-w, len(bs)-1, steps(len(bs)+w-1, 0,
		steps(-w, len(bs)-1, 0, w, bs), w, bs), w, bs)
}

func steps(fr, to, ofst, w int, bs []dir) int {
	h := len(bs) / w
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
	return bfs(fr, -1, func(n, p int) (ps []int) {
		if p == fr {
			ps = append(ps, p)
		}
		for _, d := range []dir{nn, dn, rt, lt, up} {
			if p, ok := d.move(p, h, w); ok && !fs[(n+ofst+1)%len(fs)][p] && !vs[[2]int{n, p}] {
				ps, vs[[2]int{n, p}] = append(ps, p), true
			}
		}
		return ps
	}, func(p int) bool {
		return p == to
	}) + ofst + 1
}

func bfs[T comparable](s, sp T, ps func(n int, s T) []T, ck func(s T) bool) (n int) {
	for q := (queue[T]{s, sp}); ; {
		var s T
		q, s, _ = q.deq()
		if s == sp {
			n, q = n+1, q.enq(sp)
			continue
		}
		if ck(s) {
			return n
		}
		q = q.enq(ps(n, s)...)
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
