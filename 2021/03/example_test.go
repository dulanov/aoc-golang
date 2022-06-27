// https://adventofcode.com/2021/day/3
package d03_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"log"
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
	gamma   uint64
	epsilon uint64
}

func (rt rate) String() string {
	return fmt.Sprintf("0b%b / 0b%b", rt.gamma, rt.epsilon)
}

type rating struct {
	o2  uint64
	co2 uint64
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
	ns, w := split(r)
	for i := w - 1; i >= 0; i-- {
		a := dissect(ns, i)
		if len(a[0]) <= len(a[1]) {
			rt.gamma |= 1 << i
		} else {
			rt.epsilon |= 1 << i
		}
	}
	return rt
}

func PartTwo(r io.ReadSeeker) (rt rating) {
	ns, w := split(r)
	cl := func(ns []uint64, fn func([2][]uint64) []uint64) uint64 {
		for i := w - 1; len(ns) != 1; i-- {
			ns = fn(dissect(ns, i))
		}
		return ns[0]
	}
	rt.o2 = cl(ns, func(a [2][]uint64) []uint64 {
		if len(a[1]) >= len(a[0]) {
			return a[1]
		} else {
			return a[0]
		}
	})
	rt.co2 = cl(ns, func(a [2][]uint64) []uint64 {
		if len(a[1]) < len(a[0]) {
			return a[1]
		} else {
			return a[0]
		}
	})
	return rt
}

func dissect(ns []uint64, shr int) (a [2][]uint64) {
	for _, n := range ns {
		if (n>>shr)&1 == 0 {
			a[0] = append(a[0], n)
		} else {
			a[1] = append(a[1], n)
		}
	}
	return a
}

func split(r io.Reader) (ns []uint64, w int) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		n, err := strconv.ParseUint(sc.Text(), 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		ns = append(ns, n)
		if w == 0 {
			w = len(sc.Text())
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	return ns, w
}
