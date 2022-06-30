// https://adventofcode.com/2021/day/3
package d03_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"testing"
)

const input_test = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

//go:embed testdata/input
var input string

type rate struct {
	gamma, epsilon uint64
}

func (rt rate) String() string {
	return fmt.Sprintf("0b%b / 0b%b", rt.gamma, rt.epsilon)
}

type rating struct {
	o2, co2 uint64
}

func (rtg rating) String() string {
	return fmt.Sprintf("0b%b / 0b%b", rtg.o2, rtg.co2)
}

func ExamplePartOne() {
	rt := PartOne(strings.NewReader(input))
	fmt.Println(rt.gamma * rt.epsilon)
	// Output:
	// 3148794
}

func ExamplePartTwo() {
	rt := PartTwo(strings.NewReader(input))
	fmt.Println(rt.o2 * rt.co2)
	// Output:
	// 2795310
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(input_test))
	want := rate{gamma: 0b10110, epsilon: 0b1001}
	if got != want {
		t.Errorf("got %s; want %s", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(input_test))
	want := rating{o2: 0b10111, co2: 0b01010}
	if got != want {
		t.Errorf("got %s; want %s", got, want)
	}
}

func PartOne(r io.Reader) (rt rate) {
	ns, bits := scan(r)
	for i := bits - 1; i >= 0; i-- {
		ns0, ns1 := split(ns, i, false)
		if len(ns0) <= len(ns1) {
			rt.gamma |= 1 << i
		} else {
			rt.epsilon |= 1 << i
		}
	}
	return rt
}

func PartTwo(r io.ReadSeeker) (rt rating) {
	ns, bits := scan(r)
	sort.Slice(ns, func(i, j int) bool { return ns[i] < ns[j] })
	rt.o2 = filter(ns, bits, func(ns0, ns1 []uint64) bool {
		return len(ns1) >= len(ns0)
	})
	rt.co2 = filter(ns, bits, func(ns0, ns1 []uint64) bool {
		return len(ns1) < len(ns0)
	})
	return rt
}

func filter(ns []uint64, bits int, cmp func([]uint64, []uint64) bool) uint64 {
	for i := bits - 1; len(ns) != 1; i-- {
		ns0, ns1 := split(ns, i, true)
		if cmp(ns0, ns1) {
			ns = ns1
		} else {
			ns = ns0
		}
	}
	return ns[0]
}

func split(ns []uint64, shr int, sorted bool) (ns0, ns1 []uint64) {
	if sorted {
		i := sort.Search(len(ns), func(i int) bool {
			return (ns[i]>>shr)&1 == 1
		})
		return ns[:i], ns[i:]
	}
	for _, n := range ns {
		if (n>>shr)&1 == 0 {
			ns0 = append(ns0, n)
		} else {
			ns1 = append(ns1, n)
		}
	}
	return ns0, ns1
}

func scan(r io.Reader) (ns []uint64, bits int) {
	for s := bufio.NewScanner(r); s.Scan(); {
		if bits == 0 {
			bits = len(s.Text())
		}
		n, _ := strconv.ParseUint(s.Text(), 2, 64)
		ns = append(ns, n)
	}
	return ns, bits
}
