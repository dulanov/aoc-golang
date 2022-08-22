// https://adventofcode.com/2021/day/12
package d12_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"sort"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

type stack[T comparable] []T

func (s stack[T]) empty() bool {
	return len(s) == 0
}

func (s stack[T]) contains(v T) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

func (s stack[T]) push(v ...T) stack[T] {
	return append(s, v...)
}

func (s stack[T]) pop() (stack[T], T, bool) {
	if len(s) == 0 {
		return s, *new(T), false
	}
	return s[:len(s)-1], s[len(s)-1], true
}

type cave struct {
	name  string
	paths []string
}

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 3887
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 104834
}

func TestPartOne(t *testing.T) {
	for i, tc := range []struct {
		in   string
		want int
	}{
		{
			in:   input_test0,
			want: 10,
		},
		{
			in:   input_test1,
			want: 19,
		},
		{
			in:   input_test2,
			want: 226,
		},
	} {
		t.Run(fmt.Sprintf("input_test%d", i), func(t *testing.T) {
			if got := PartOne(strings.NewReader(tc.in)); got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	for i, tc := range []struct {
		in   string
		want int
	}{
		{
			in:   input_test0,
			want: 36,
		},
		{
			in:   input_test1,
			want: 103,
		},
		{
			in:   input_test2,
			want: 3509,
		},
	} {
		t.Run(fmt.Sprintf("input_test%d", i), func(t *testing.T) {
			if got := PartTwo(strings.NewReader(tc.in)); got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}
}

func PartOne(r io.Reader) int {
	cs := scan(r)
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].name < cs[j].name
	})
	return len(paths(cs, 0))
}

func PartTwo(r io.Reader) int {
	cs := scan(r)
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].name < cs[j].name
	})
	return len(paths(cs, 1))
}

func paths(cs []cave, lm int) (ps []string) {
	lbl, lbs, st := lm, stack[string]{}, stack[string]{"start"}
	for !st.empty() {
		var s string
		st, s, _ = st.pop()
		switch s {
		case "start":
			if lbs.empty() {
				lbs, st = lbs.push(s), st.push(search(cs, s).paths...)
			}
		case "end":
			ps = append(ps, strings.Join(append(lbs, "end"), ","))
		case "*":
			lbl++
			fallthrough
		case "+":
			lbs, _, _ = lbs.pop()
		default:
			if strings.ToLower(s) == s && lbs.contains(s) {
				if lbl == 0 {
					continue
				}
				lbl, st = lbl-1, st.push("*")
			} else {
				st = st.push("+")
			}
			lbs, st = lbs.push(s), st.push(search(cs, s).paths...)
		}
	}
	return ps
}

func search(cs []cave, name string) cave {
	return cs[sort.Search(len(cs), func(i int) bool {
		return cs[i].name >= name
	})]
}

func scan(r io.Reader) (cs []cave) {
	m := map[string]*cave{}
	for s := bufio.NewScanner(r); s.Scan(); {
		ss := strings.Split(s.Text(), "-")
		for _, s := range ss {
			if _, ok := m[s]; !ok {
				m[s] = &cave{name: s}
			}
		}
		m[ss[0]].paths = append(m[ss[0]].paths, ss[1])
		m[ss[1]].paths = append(m[ss[1]].paths, ss[0])
	}
	for _, c := range m {
		cs = append(cs, *c)
	}
	return cs
}

const (
	input_test0 = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

	input_test1 = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

	input_test2 = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`
)
