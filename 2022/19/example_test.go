// https://adventofcode.com/2022/day/19
package d19_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"reflect"
	"strings"
	"sync"
	"testing"

	"golang.org/x/exp/constraints"
)

//go:embed testdata/input
var input string

type stack[T any] []T

func (s stack[T]) empty() bool {
	return len(s) == 0
}

func (s stack[T]) push(v ...T) stack[T] {
	return append(s, v...)
}

func (s stack[T]) pop() (stack[T], T, bool) {
	if s.empty() {
		return s, *new(T), false
	}
	return s[:len(s)-1], s[len(s)-1], true
}

type res int

type robot = res

type factory [6]int

type blueprint [4][3]int

func (f factory) build(r robot, bl blueprint) (f2 factory, tm int, ok bool) {
	if (r == r2 && f[r1] == 0) ||
		(r == r3 && f[r2] == 0) {
		return f, 0, false
	}
	d1, d2, d3 := f[3]-bl[r][0], f[4]-bl[r][1], f[5]-bl[r][2]
	if d1 >= 0 && d2 >= 0 && d3 >= 0 {
		tm = 1
	} else if r == r0 || r == r1 {
		tm = (f[0]-d1-1)/f[0] + 1
	} else if r == r2 {
		tm = max((f[0]-d1-1)/f[0]+1, (f[1]-d2-1)/f[1]+1)
	} else if r == r3 {
		tm = max((f[0]-d1-1)/f[0]+1, (f[2]-d3-1)/f[2]+1)
	}
	f[3], f[4], f[5] = d1+f[0]*tm, d2+f[1]*tm, d3+f[2]*tm
	if r != r3 {
		f[r]++
	}
	return f, tm, true
}

const (
	r0 res = iota /* ore */
	r1            /* clay */
	r2            /* obsidian */
	r3            /* geode */
)

func ExamplePartOne() {
	ns := PartOne(strings.NewReader(input), 2)
	fmt.Println(fld(ns))
	// Output:
	// 1834
}

func ExamplePartTwo() {
	ns := PartTwo(strings.NewReader(input), 3, 2)
	fmt.Println(mul(ns...))
	// Output:
	// 2240
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest), 2)
	want := []int{9, 12}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest), 2, 2)
	want := []int{56, 62}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func PartOne(r io.Reader, j int) (ns []int) {
	return conv(scan(r), j, func(b blueprint) int {
		return opt(b, 24)
	})
}

func PartTwo(r io.Reader, n, j int) (ns []int) {
	return conv(scan(r)[:n], j, func(b blueprint) int {
		return opt(b, 32)
	})
}

func opt(bl blueprint, m int) (n int) {
	type st struct {
		nm, tm int
		fty    factory
	}
	lm := max(bl[r0][r0], bl[r1][r0], bl[r2][r0], bl[r3][r0])
	dfs(st{fty: factory{1}}, func(s st) (ss []st) {
		if s.nm > n {
			n = s.nm
		}
		if (m-s.tm)*(m-s.tm+1) <= (n-s.nm)*2 {
			return ss
		}
		for _, r := range []robot{r3, r2, r1, r0} {
			if (r == r0 && s.fty[r0] == lm) ||
				(r == r1 && s.fty[r1] == bl[r2][r1]) ||
				(r == r2 && s.fty[r2] == bl[r3][r2]) {
				continue
			}
			if f, t, ok := s.fty.build(r, bl); ok {
				if s.tm+t >= m {
					continue
				}
				if r != r3 {
					ss = append(ss, st{s.nm, s.tm + t, f})
				} else {
					ss = append(ss, st{s.nm + (m - s.tm - t), s.tm + t, f})
				}
			}
		}
		return ss
	})
	return n
}

func dfs[T any](s T, fn func(s T) []T) {
	for st := (stack[T]{s}); !st.empty(); {
		var s T
		st, s, _ = st.pop()
		for _, s = range fn(s) {
			st = st.push(s)
		}
	}
}

func conv[T1, T2 any](vs []T1, j int, fn func(v T1) T2) (rs []T2) {
	rs = make([]T2, len(vs))
	for i := 0; i < len(vs); i += j {
		var wg sync.WaitGroup
		for k := 0; k < j && i+k < len(vs); k++ {
			wg.Add(1)
			go func(k int) {
				defer wg.Done()
				rs[i+k] = fn(vs[i+k])
			}(k)
		}
		wg.Wait()
	}
	return rs
}

func max[T constraints.Ordered](ns ...T) (n T) {
	n = ns[0]
	for _, v := range ns[1:] {
		if v > n {
			n = v
		}
	}
	return n
}

func mul[T constraints.Integer](ns ...T) (n T) {
	n = ns[0]
	for _, v := range ns[1:] {
		n *= v
	}
	return n
}

func fld[T constraints.Integer](ns []T) (n T) {
	for i, v := range ns {
		n += T(i+1) * v
	}
	return n
}

func scan(r io.Reader) (bs []blueprint) {
	for s := bufio.NewScanner(r); s.Scan(); {
		var i, n1, n2, n3, n4, n5, n6 int
		fmt.Sscanf(s.Text(), "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&i, &n1, &n2, &n3, &n4, &n5, &n6)
		bs = append(bs, blueprint{{n1, 0, 0}, {n2, 0, 0}, {n3, n4, 0}, {n5, 0, n6}})
	}
	return bs
}

const inputTest = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`
