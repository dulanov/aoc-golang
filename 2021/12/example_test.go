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

func (s stack[T]) contains(v T) bool {
	for _, s := range s {
		if s == v {
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
	// 0
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
	got := PartTwo(strings.NewReader(input_test0))
	want := 0
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) int {
	cs := scan(r)
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].name < cs[j].name
	})
	return len(paths(cs))
}

func PartTwo(r io.Reader) int {
	return 0
}

func paths(cs []cave) (ps []string) {
	lbs, st := stack[string]{}, stack[string]{"start"}
	for len(st) != 0 {
		var s string
		st, s, _ = st.pop()
		switch s {
		case "":
			lbs, _, _ = lbs.pop()
		case "end":
			ps = append(ps, strings.Join(append(lbs, "end"), ","))
		default:
			if strings.ToUpper(s) == s || !lbs.contains(s) {
				lbs, st = lbs.push(s), st.push("")
				st = st.push(search(cs, s).paths...)
			}
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
