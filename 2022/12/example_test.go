// https://adventofcode.com/2022/day/12
package d12_test

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

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 440
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 439
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := 31
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest))
	want := 29
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) int {
	hs, w, fr, to := scan(r)
	return bfs(to, func(i int) (rs []int) {
		for _, d := range []int{-1, 1, -w, w} {
			if j := i + d; j >= 0 && j < len(hs) && abs(hs[i])-1 <= hs[j] {
				rs, hs[j] = append(rs, j), -hs[j]
			}
		}
		return rs
	}, func(i int) bool {
		return i == fr
	})
}

func PartTwo(r io.Reader) int {
	hs, w, _, to := scan(r)
	return bfs(to, func(i int) (rs []int) {
		for _, d := range []int{-1, 1, -w, w} {
			if j := i + d; j >= 0 && j < len(hs) && abs(hs[i])-1 <= hs[j] {
				rs, hs[j] = append(rs, j), -hs[j]
			}
		}
		return rs
	}, func(i int) bool {
		return hs[i] == -1
	})
}

func bfs(fr int, ps func(p int) []int, ck func(p int) bool) (n int) {
	for q := (queue[int]{fr, -1}); ; {
		var p int
		q, p, _ = q.deq()
		if p == -1 {
			n, q = n+1, q.enq(-1)
			continue
		}
		if ck(p) {
			return n
		}
		q = q.enq(ps(p)...)
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func scan(r io.Reader) (hs []int, w, fr, to int) {
	for s := bufio.NewScanner(r); s.Scan(); {
		if w == 0 {
			w = len(s.Text())
		}
		ns := make([]int, len(s.Text()))
		for i, r := range s.Text() {
			if r == 'S' {
				r, fr = 'a', len(hs)+i
			} else if r == 'E' {
				r, to = 'z', len(hs)+i
			}
			ns[i] = int(r - 'a' + 1)
		}
		hs = append(hs, ns...)
	}
	return hs, w, fr, to
}

const inputTest = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
