// https://adventofcode.com/2022/day/11
package d11_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

type monkey struct {
	items []int
	calc  func(int) int
	next  func(int) int
}

func ExamplePartOne() {
	ns := PartOne(strings.NewReader(input))
	sort.Sort(sort.Reverse(sort.IntSlice(ns)))
	fmt.Println(ns[0] * ns[1])
	// Output:
	// 112221
}

func ExamplePartTwo() {
	ns := PartTwo(strings.NewReader(input))
	sort.Sort(sort.Reverse(sort.IntSlice(ns)))
	fmt.Println(ns[0] * ns[1])
	// Output:
	// 25272176808
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := []int{101, 95, 7, 105}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest))
	want := []int{52166, 47830, 1938, 52013}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func PartOne(r io.Reader) (ns []int) {
	ms, _ := scan(r)
	return play(ms, 20, func(n int) int {
		return n / 3
	})
}

func PartTwo(r io.Reader) (ns []int) {
	ms, lcm := scan(r)
	return play(ms, 10_000, func(n int) int {
		return n % lcm
	})
}

func play(ms []monkey, nm int, fn func(int) int) (ns []int) {
	ns = make([]int, len(ms))
	for i := 0; i < nm; i++ {
		for j, m := range ms {
			for _, v := range m.items {
				v = fn(m.calc(v))
				ms[m.next(v)].items = append(ms[m.next(v)].items, v)
			}
			ns[j], ms[j].items = ns[j]+len(m.items), []int{}
		}
	}
	return ns
}

func scan(r io.Reader) (ms []monkey, lcm int) {
	lcm = 1
	b := strings.Builder{}
	for s := bufio.NewScanner(r); s.Scan(); {
		if s.Text() == "" {
			ms = append(ms, parse(b))
			b.Reset()
			continue
		}
		if s.Text()[2] == 'T' {
			n, _ := strconv.Atoi(s.Text()[21:])
			lcm *= n
		}
		b.WriteString(s.Text())
		b.WriteByte('\n')
	}
	return append(ms, parse(b)), lcm
}

func parse(b strings.Builder) monkey {
	var op byte
	var its, od string
	var i, tc, tn1, tn2 int
	fmt.Sscanf(strings.ReplaceAll(b.String(), ", ", ","), `Monkey %d:
  Starting items: %s
  Operation: new = old %c %s
  Test: divisible by %d
    If true: throw to monkey %d
    If false: throw to monkey %d`, &i, &its, &op, &od, &tc, &tn1, &tn2)
	var m monkey
	for _, s := range strings.Split(its, ",") {
		n, _ := strconv.Atoi(s)
		m.items = append(m.items, n)
	}
	if od == "old" {
		m.calc = func(n int) int { return n * n }
	} else {
		n1, _ := strconv.Atoi(od)
		if op == '+' {
			m.calc = func(n int) int { return n + n1 }
		} else {
			m.calc = func(n int) int { return n * n1 }
		}
	}
	m.next = func(n int) int {
		if n%tc == 0 {
			return tn1
		}
		return tn2
	}
	return m
}

const inputTest = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`
