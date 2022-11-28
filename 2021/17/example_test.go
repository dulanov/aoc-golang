// https://adventofcode.com/2021/day/17
package d17_test

import (
	_ "embed"
	"fmt"
	"io"
	"math"
	"sort"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

type pos [2]int

func (p pos) less(o pos) bool {
	return (p[0] != o[0] && p[0] < o[0]) ||
		(p[0] == o[0] && p[1] < o[1])
}

type vel = pos

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 5460
}

func ExamplePartTwo() {
	fmt.Println(len(PartTwo(strings.NewReader(input))))
	// Output:
	// 3618
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := 45
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest))
	want := []vel{
		{23, -10}, {25, -9}, {27, -5}, {29, -6}, {22, -6}, {21, -7},
		{9, 0}, {27, -7}, {24, -5}, {25, -7}, {26, -6}, {25, -5},
		{6, 8}, {11, -2}, {20, -5}, {29, -10}, {6, 3}, {28, -7},
		{8, 0}, {30, -6}, {29, -8}, {20, -10}, {6, 7}, {6, 4},
		{6, 1}, {14, -4}, {21, -6}, {26, -10}, {7, -1}, {7, 7},
		{8, -1}, {21, -9}, {6, 2}, {20, -7}, {30, -10}, {14, -3}, {20, -8},
		{13, -2}, {7, 3}, {28, -8}, {29, -9}, {15, -3}, {22, -5}, {26, -8},
		{25, -8}, {25, -6}, {15, -4}, {9, -2}, {15, -2}, {12, -2}, {28, -9},
		{12, -3}, {24, -6}, {23, -7}, {25, -10}, {7, 8}, {11, -3}, {26, -7},
		{7, 1}, {23, -9}, {6, 0}, {22, -10}, {27, -6}, {8, 1}, {22, -8},
		{13, -4}, {7, 6}, {28, -6}, {11, -4}, {12, -4}, {26, -9}, {7, 4},
		{24, -10}, {23, -8}, {30, -8}, {7, 0}, {9, -1}, {10, -1}, {26, -5},
		{22, -9}, {6, 5}, {7, 5}, {23, -6}, {28, -1}, {10, -2}, {11, -1},
		{20, -9}, {14, -2}, {29, -7}, {13, -3}, {23, -5}, {24, -8}, {27, -9},
		{30, -7}, {28, -5}, {21, -1}, {7, 9}, {6, 6}, {21, -5}, {27, -10},
		{7, 2}, {30, -9}, {21, -8}, {22, -7}, {24, -9}, {20, -6}, {6, 9},
		{29, -5}, {8, -2}, {27, -8}, {30, -5}, {24, -7}}
	if len(got) != len(want) {
		t.Fatalf("got %d; want %d", len(got), len(want))
	}
	for _, v := range want {
		if sort.Search(len(got), func(i int) bool {
			return !got[i].less(v)
		}) == len(got) {
			t.Errorf("%v wasn't found", v)
		}
	}
}

func PartOne(r io.Reader) (hp int) {
	ps := scan(r)
	ss := genY(ps[1][1], ps[0][1])
	return sumOf(ss[len(ss)-2][0])
}

func PartTwo(r io.Reader) (rs []vel) {
	ps := scan(r)
	vs := flt(mul(genX(ps[0][0], ps[1][0]),
		genY(ps[1][1], ps[0][1])))
	sort.Slice(vs, func(i, j int) bool {
		return vs[i].less(vs[j])
	})
	return unq(vs)
}

func genX(p1, p2 int) (ss [][]int) {
	ss = make([][]int, int(math.Sqrt(float64(p2*2))))
	for vl, ps, st := p2, p2, 1; vl >= st; vl, ps = vl-1, ps-st {
		for ps < p1 && vl >= st {
			ps, st = ps+vl-st, st+1
		}
		for i, p := 0, ps; p >= p1 && p <= p2; i, p = i+1, p+vl-st-i {
			if vl == i+st {
				for j := i + st - 1; j < len(ss); j++ {
					ss[j] = append(ss[j], vl)
				}
				break
			}
			ss[i+st-1] = append(ss[i+st-1], vl)
		}
	}
	return ss
}

func genY(p1, p2 int) (ss [][]int) {
	ss = make([][]int, -p2*2+1)
	for vl, ps, st := p2, p2, 1; vl <= 0; vl, ps = vl+1, ps+st {
		if ps > p1 {
			ps, st = ps+vl-st, st+1
		}
		for i, p := 0, ps; p >= p2; i, p = i+1, p+vl-st-i {
			ss[i+st-1] = append(ss[i+st-1], vl)
			if vl <= -2 /* from the opposite side */ {
				ss[i+st-vl*2-2] = append(ss[i+st-vl*2-2], -vl-1)
			}
		}
	}
	return ss
}

func mul(ss1, ss2 [][]int) (ss [][]vel) {
	ss = make([][]vel, max(len(ss1), len(ss2)))
	for i := range ss {
		vs1, vs2 := elm(ss1, ss2, i)
		ss[i] = make([]vel, len(vs1)*len(vs2))
		for j, v1 := range vs1 {
			for k, v2 := range vs2 {
				ss[i][j*len(vs2)+k] = vel{v1, v2}
			}
		}
	}
	return ss
}

func elm(ss1, ss2 [][]int, n int) (vs1, vs2 []int) {
	if n >= len(ss1) {
		return ss1[len(ss1)-1], ss2[n]
	}
	if n >= len(ss2) {
		return ss1[n], ss2[len(ss2)-1]
	}
	return ss1[n], ss2[n]
}

func flt[T any](ss [][]T) (vs []T) {
	for _, v := range ss {
		vs = append(vs, v...)
	}
	return vs
}

func unq[T comparable](vs []T) (rs []T) {
	pr := 0
	for i := 1; i < len(vs); i++ {
		if vs[i] != vs[i-1] {
			vs[pr+1], pr = vs[i], pr+1
		}
	}
	return vs[:pr+1]
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func sumOf(n int) int {
	return n * (n + 1) / 2
}

func scan(r io.Reader) (ps [2]pos) {
	fmt.Fscanf(r, "target area: x=%d..%d, y=%d..%d",
		&ps[0][0], &ps[1][0], &ps[0][1], &ps[1][1])
	return ps
}

const inputTest = `target area: x=20..30, y=-10..-5`
