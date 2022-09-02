// https://adventofcode.com/2021/day/14
package d14_test

import (
	_ "embed"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

func ExamplePartOne() {
	ns := PartOne(strings.NewReader(input))
	fmt.Println(ns[len(ns)-1][1] - ns[0][1])
	// Output:
	// 2967
}

func ExamplePartTwo() {
	ns := PartTwo(strings.NewReader(input))
	fmt.Println(ns[len(ns)-1][1] - ns[0][1])
	// Output:
	// 3692219987038
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(input_test))
	want := [][2]int{
		{int('H'), 161},
		{int('C'), 298},
		{int('N'), 865},
		{int('B'), 1749}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(input_test))
	want := [...][2]int{
		{int('H'), 3849876073},
		{int('B'), 2192039569602}}
	if got := [...][2]int{got[0], got[len(got)-1]}; got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func PartOne(r io.Reader) (ns [][2]int) {
	tl, rs := scan(r)
	ns = times(tl[len(tl)-1], steps(tl, rs, 10))
	sort.Slice(ns, func(i, j int) bool {
		return ns[i][1] < ns[j][1]
	})
	return ns
}

func PartTwo(r io.Reader) (ns [][2]int) {
	tl, rs := scan(r)
	ns = times(tl[len(tl)-1], steps(tl, rs, 40))
	sort.Slice(ns, func(i, j int) bool {
		return ns[i][1] < ns[j][1]
	})
	return ns
}

func steps(tl []byte, rs [][3]byte, n int) (ps []struct {
	pr [2]byte
	nm int
}) {
	rm := make(map[[2]byte]byte)
	for _, r := range rs {
		rm[[2]byte{r[0], r[1]}] = r[2]
	}
	pm := make(map[[2]byte]int, len(rm))
	for i := 0; i < len(tl)-1; i++ {
		pm[[2]byte{tl[i], tl[i+1]}]++
	}
	for i := 0; i < n; i++ {
		pm2 := make(map[[2]byte]int, len(rm))
		for pr, n := range pm {
			pm2[[2]byte{pr[0], rm[pr]}] += n
			pm2[[2]byte{rm[pr], pr[1]}] += n
		}
		pm = pm2
	}
	for pr, n := range pm {
		ps = append(ps, struct {
			pr [2]byte
			nm int
		}{pr, n})
	}
	return ps
}

func times(lt byte, ps []struct {
	pr [2]byte
	nm int
}) (ns [][2]int) {
	nm := map[byte]int{lt: 1}
	for _, p := range ps {
		nm[p.pr[0]] += p.nm
	}
	for b, n := range nm {
		ns = append(ns, [2]int{int(b), n})
	}
	return ns
}

func scan(r io.Reader) (tl []byte, rs [][3]byte) {
	var s string
	var bs [3]byte
	fmt.Fscanf(r, "%s\n\n", &s)
	for {
		if _, err := fmt.Fscanf(r, "%c%c -> %c\n", &bs[0], &bs[1], &bs[2]); err == io.EOF {
			break
		}
		rs = append(rs, bs)
	}
	return []byte(s), rs
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
