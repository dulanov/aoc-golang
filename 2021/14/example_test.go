// https://adventofcode.com/2021/day/14
package d14_test

import (
	_ "embed"
	"fmt"
	"io"
	"math"
	"reflect"
	"sort"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

type res struct {
	pl byte
	nm int
}

func ExamplePartOne() {
	rs := PartOne(strings.NewReader(input))
	fmt.Println(rs[len(rs)-1].nm - rs[0].nm)
	// Output:
	// 2967
}

func ExamplePartTwo() {
	rs := PartTwo(strings.NewReader(input))
	fmt.Println(rs[len(rs)-1].nm - rs[0].nm)
	// Output:
	// 3692219987038
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(input_test))
	want := []res{
		{'H', 161},
		{'C', 298},
		{'N', 865},
		{'B', 1749}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(input_test))
	want := [...]res{
		{'H', 3849876073},
		{'B', 2192039569602}}
	if got := [...]res{got[0], got[len(got)-1]}; got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func PartOne(r io.Reader) (rs []res) {
	tl, rls := scan(r)
	sort.Slice(rls, func(i, j int) bool {
		return rls[i][0] < rls[j][0]
	})
	rs = steps(tl, rls, 10)
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].nm < rs[j].nm
	})
	return rs
}

func PartTwo(r io.Reader) (rs []res) {
	tl, rls := scan(r)
	sort.Slice(rls, func(i, j int) bool {
		return rls[i][0] < rls[j][0]
	})
	rs = steps(tl, rls, 40)
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].nm < rs[j].nm
	})
	return rs
}

func steps(tl []byte, rls [][3]byte, n int) (rs []res) {
	pm := make(map[[2]byte]int, len(rls))
	for i, r := range rls {
		pm[[2]byte{r[0], r[1]}] = i
	}
	ps := make([][2]int, len(rls))
	for i, r := range rls {
		ps[i] = [2]int{
			pm[[2]byte{r[0], r[2]}],
			pm[[2]byte{r[2], r[1]}]}
	}
	ns := make([]int, len(rls))
	for i := 0; i < len(tl)-1; i++ {
		ns[pm[[2]byte{tl[i], tl[i+1]}]]++
	}
	for i, nx := 0, make([]int, len(rls)); i < n; i++ {
		for j, n := range ns {
			nx[ps[j][0]] += n
			nx[ps[j][1]] += n
			ns[j] = 0
		}
		ns, nx = nx, ns
	}
	rs = make([]res, int(math.Sqrt(float64(len(rls)))))
	for i, j := 0, 0; i < len(rs); i, j = i+1, j+len(rs) {
		rs[i] = res{rls[j][0], sum(ns[j : j+len(rs)])}
	}
	rs[sort.Search(len(rs), func(i int) bool {
		a, b := rs[i], tl[len(tl)-1]
		return a.pl >= b
	})].nm++
	return rs
}

func sum(ns []int) (rs int) {
	for _, n := range ns {
		rs += n
	}
	return rs
}

func scan(r io.Reader) (tl []byte, rls [][3]byte) {
	var s string
	var bs [3]byte
	fmt.Fscanf(r, "%s\n\n", &s)
	for {
		if _, err := fmt.Fscanf(r, "%c%c -> %c\n", &bs[0], &bs[1], &bs[2]); err == io.EOF {
			break
		}
		rls = append(rls, bs)
	}
	return []byte(s), rls
}

const input_test = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`
