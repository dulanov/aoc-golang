// https://adventofcode.com/2022/day/16
package d16_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"math"
	"math/bits"
	"strings"
	"testing"
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

type valve struct {
	id, rt int
	cs     []int
}

func BenchmarkExamplePartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartOne(strings.NewReader(input), true)
	}
}

func BenchmarkExamplePartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartTwo(strings.NewReader(input), true)
	}
}

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input), true))
	// Output:
	// 1796
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input), true))
	// Output:
	// 1999
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest), false)
	want := 1651
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	t.Skip()
	got := PartTwo(strings.NewReader(inputTest), false)
	want := 1707
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader, opt bool) (n int) {
	vs := scan(r)
	ds, vs := fwa(vs), flt(vs, func(v valve) bool {
		return v.rt > 0
	})
	dfs(30, (1<<len(vs))-1, ds, vs, func(pt, rt int) {
		if rt > n {
			n = rt
		}
	}, opt)
	return n
}

func PartTwo(r io.Reader, opt bool) (n int) {
	ps, vs := make(map[int]int), scan(r)
	ds, vs := fwa(vs), flt(vs, func(v valve) bool {
		return v.rt > 0
	})
	dfs(26, ((1 << len(vs)) - 1), ds, vs, func(pt, rt int) {
		if pt = ((1 << len(vs)) - 1) &^ pt; bits.OnesCount16(uint16(pt)) >= 4 {
			if rt > ps[pt] {
				ps[pt] = rt
			}
		}
	}, opt)
	for p1, n1 := range ps {
		for p2, n2 := range ps {
			if p1&p2 == 0 && n1+n2 > n {
				n = n1 + n2
			}
		}
	}
	return n
}

// https://en.wikipedia.org/wiki/Depth-first_search#Pseudocode
func dfs(lm, pt int, ds [][]int, vs []valve, fn func(pt, rt int), opt bool) {
	type state struct {
		tm, pt, rt, id int
	}
	for st := (stack[state]{{lm, pt, 0, 0}}); !st.empty(); {
		var s state
		st, s, _ = st.pop()
		fn(s.pt, s.rt)
		for i, j := 0, s.pt; j != 0; i, j = i+1, j>>1 {
			if v := vs[i]; j&1 != 0 && (!opt || s.tm < 20 || ds[s.id][v.id] <= 3) {
				if t := s.tm - ds[s.id][v.id] - 1; t > 0 {
					st = st.push(state{t, s.pt &^ (1 << i), s.rt + v.rt*t, v.id})
				}
			}
		}
	}
}

// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm#Path_reconstruction
func fwa(vs []valve) (ds [][]int) {
	ds = make([][]int, len(vs))
	for i := 0; i < len(vs); i++ {
		ds[i] = make([]int, len(vs))
		for j := 0; j < len(vs); j++ {
			ds[i][j] = math.MaxUint8
		}
		for _, c := range vs[i].cs {
			ds[i][c] = 1
		}
		ds[i][i] = 0
	}
	for k := 0; k < len(vs); k++ {
		for j := 0; j < len(vs); j++ {
			for i := 0; i < len(vs); i++ {
				if d := ds[i][k] + ds[k][j]; ds[i][j] > d {
					ds[i][j] = d
				}
			}
		}
	}
	return ds
}

func flt[T any](vs []T, fn func(v T) bool) (rs []T) {
	for _, v := range vs {
		if fn(v) {
			rs = append(rs, v)
		}
	}
	return rs
}

func scan(r io.Reader) (vs []valve) {
	vs = []valve{{}}
	ids, tls := map[string]int{"AA": 0}, [][]string{{""}}
	rpl := strings.NewReplacer(", ", ",", "tunnel leads to valve", "tunnels lead to valves")
	for i, s := 1, bufio.NewScanner(r); s.Scan(); i++ {
		var n int
		var s1, s2 string
		fmt.Sscanf(rpl.Replace(s.Text()),
			"Valve %s has flow rate=%d; tunnels lead to valves %s", &s1, &n, &s2)
		if s1 == "AA" {
			i, tls[0], vs[0] = i-1, strings.Split(s2, ","), valve{id: 0, rt: 0}
			continue
		}
		ids[s1], tls, vs = i, append(tls, strings.Split(s2, ",")), append(vs, valve{id: i, rt: n})
	}
	for i, t := range tls {
		vs[i].cs = make([]int, len(t))
		for j, s := range t {
			vs[i].cs[j] = ids[s]
		}
	}
	return vs
}

const inputTest = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`
