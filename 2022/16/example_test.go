// https://adventofcode.com/2022/day/16
package d16_test

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
	conns  []int
}

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 1796
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 0
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := 1651
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest))
	want := 1707
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (n int) {
	vs := scan(r)
	ds := make([][]int, len(vs))
	for i := 0; i < len(vs); i++ {
		ds[i] = make([]int, len(vs))
		for j := 0; j < len(vs); j++ {
			ds[i][j] = math.MaxUint8
		}
		for _, c := range vs[i].conns {
			ds[i][c] = 1
		}
		ds[i][i] = 0
	}
	// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm#Path_reconstruction
	for k := 0; k < len(vs); k++ {
		for j := 0; j < len(vs); j++ {
			for i := 0; i < len(vs); i++ {
				if d := ds[i][k] + ds[k][j]; ds[i][j] > d {
					ds[i][j] = d
				}
			}
		}
	}
	vs = flt(vs, func(v valve) bool {
		return v.rt > 0
	})
	// https://en.wikipedia.org/wiki/Depth-first_search#Pseudocode
	type state struct {
		tm, pt, rt, id int
	}
	for st := (stack[state]{{30, (1 << len(vs)) - 1, 0, 0}}); !st.empty(); {
		var s state
		st, s, _ = st.pop()
		if s.rt > n {
			n = s.rt
		}
		for i, j := 0, s.pt; j != 0; i, j = i+1, j>>1 {
			if v := vs[i]; j&1 != 0 {
				if t := s.tm - ds[s.id][v.id] - 1; t > 0 {
					st = st.push(state{t, s.pt &^ (1 << i), s.rt + v.rt*t, v.id})
				}
			}
		}
	}
	return n
}

func PartTwo(r io.Reader) (n int) {
	return n
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
		var s1, s2 string
		v := valve{id: i}
		fmt.Sscanf(rpl.Replace(s.Text()),
			"Valve %s has flow rate=%d; tunnels lead to valves %s", &s1, &v.rt, &s2)
		if s1 == "AA" {
			i, tls[0], vs[0] = i-1, strings.Split(s2, ","), valve{id: 0, rt: 0}
			continue
		}
		ids[s1], tls, vs = i, append(tls, strings.Split(s2, ",")), append(vs, v)
	}
	for i, t := range tls {
		vs[i].conns = make([]int, len(t))
		for j, s := range t {
			vs[i].conns[j] = ids[s]
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
