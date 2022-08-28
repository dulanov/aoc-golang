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
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 0
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
	want := 0
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (ns [][2]int) {
	tl, rs := scan(r)
	tl = steps(tl, rs, 10)
	sort.Slice(tl, func(i, j int) bool {
		return tl[i] < tl[j]
	})
	ns = group(tl)
	sort.Slice(ns, func(i, j int) bool {
		return ns[i][1] < ns[j][1]
	})
	return ns
}

func PartTwo(r io.Reader) int {
	return 0
}

func steps(tl []byte, rs [][3]byte, n int) []byte {
	rm := make(map[[2]byte]byte)
	for _, r := range rs {
		rm[[2]byte{r[0], r[1]}] = r[2]
	}
	b := append(make([]byte, 0, (1<<(n+1)-1)*(len(tl)-1)+n+1), tl...)
	for i, l := 0, len(b); i < n; i, l, b = i+1, len(b)-l, b[l:] {
		b = append(b, b[0])
		for j := 0; j < l-1; j++ {
			if r, ok := rm[[2]byte{b[j], b[j+1]}]; ok {
				b = append(b, r, b[j+1])
			}
		}
	}
	return b
}

func group(bs []byte) (ns [][2]int) {
	for i, j := 0, 1; i < len(bs); i, j = j, j+1 {
		for ; j < len(bs) && bs[j] == bs[i]; j++ {
		}
		ns = append(ns, [2]int{int(bs[i]), j - i})
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
