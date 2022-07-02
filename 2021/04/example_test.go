// https://adventofcode.com/2021/day/4
package d04_test

import (
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"
)

const input_test = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

//go:embed testdata/input
var input string

type board struct {
	nums [5][5]int
}

func (b board) sum() (n int) {
	for _, ns := range b.nums {
		n += sum(ns[:])
	}
	return n
}

func ExamplePartOne() {
	sc := PartOne(strings.NewReader(input))
	fmt.Println(sc.n * sc.p)
	// Output:
	// 29440
}

func ExamplePartTwo() {
	sc := PartTwo(strings.NewReader(input))
	fmt.Println(sc.n * sc.p)
	// Output:
	// 13884
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(input_test))
	want := struct{ n, p int }{24, 188}
	if got != want {
		t.Errorf("got %+v; want %+v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(input_test))
	want := struct{ n, p int }{13, 148}
	if got != want {
		t.Errorf("got %+v; want %+v", got, want)
	}
}

func PartOne(r io.Reader) (sc struct{ n, p int }) {
	bs, ns := scan(r)
	return score(bs, 1, ns)
}

func PartTwo(r io.ReadSeeker) (sc struct{ n, p int }) {
	bs, ns := scan(r)
	return score(bs, len(bs), ns)
}

func score(bs []board, bn int, ns []int) (sc struct{ n, p int }) {
	idx, ps := genIndex(bs), genPoints(bs)
	for _, n := range ns {
		for i := range bs {
			p, ok := idx.ps[i][n]
			if ok && ps[i][0] != 0 {
				ps[i][0] -= n
				ps[i][p.col+1]++
				ps[i][p.row+6]++
				if ps[i][p.col+1] == 5 ||
					ps[i][p.row+6] == 5 {
					if bn--; bn == 0 {
						return struct{ n, p int }{n, ps[i][0]}
					}
					ps[i][0] = 0
				}
			}
		}
	}
	return sc
}

type index struct {
	ps []map[int]struct{ row, col int }
}

func genIndex(bs []board) (idx index) {
	for _, b := range bs {
		m := make(map[int]struct{ row, col int }, 25)
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				m[b.nums[i][j]] = struct{ row, col int }{i, j}
			}
		}
		idx.ps = append(idx.ps, m)
	}
	return idx
}

func genPoints(bs []board) (ps [][]int) {
	for _, b := range bs {
		pa := [11]int{b.sum()}
		ps = append(ps, pa[:])
	}
	return ps
}

func scan(r io.Reader) (bs []board, ns []int) {
	var (
		s                       string
		n00, n01, n02, n03, n04 int
		n10, n11, n12, n13, n14 int
		n20, n21, n22, n23, n24 int
		n30, n31, n32, n33, n34 int
		n40, n41, n42, n43, n44 int
	)
	fmt.Fscanf(r, "%s\n", &s)
	for _, s := range strings.Split(s, ",") {
		n, _ := strconv.Atoi(s)
		ns = append(ns, n)
	}
	for {
		n, _ := fmt.Fscanf(r, `
%d%d%d%d%d
%d%d%d%d%d
%d%d%d%d%d
%d%d%d%d%d
%d%d%d%d%d
`,
			&n00, &n01, &n02, &n03, &n04,
			&n10, &n11, &n12, &n13, &n14,
			&n20, &n21, &n22, &n23, &n24,
			&n30, &n31, &n32, &n33, &n34,
			&n40, &n41, &n42, &n43, &n44)
		if n == 0 {
			break
		}
		bs = append(bs, board{[5][5]int{
			{n00, n01, n02, n03, n04},
			{n10, n11, n12, n13, n14},
			{n20, n21, n22, n23, n24},
			{n30, n31, n32, n33, n34},
			{n40, n41, n42, n43, n44}}})
	}
	return bs, ns
}

func sum(ns []int) (rs int) {
	for _, n := range ns {
		rs += n
	}
	return rs
}
