// https://adventofcode.com/2022/day/11
package d11_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"math"
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
	fmt.Println(mul(ns[:2]))
	// Output:
	// 112221
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 0
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
	want := 0
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (ns []int) {
	ms := scan(r)
	ns = make([]int, len(ms))
	for i := 0; i < 20; i++ {
		for j, m := range ms {
			for _, v := range m.items {
				n := int(math.Round(float64(m.calc(v) / 3)))
				ms[m.next(n)].items = append(ms[m.next(n)].items, n)
			}
			ns[j], ms[j].items = ns[j]+len(m.items), []int{}
		}
	}
	return ns
}

func PartTwo(r io.Reader) int {
	return 0
}

func mul(ns []int) (n int) {
	n = 1
	for _, v := range ns {
		n *= v
	}
	return n
}

func scan(r io.Reader) (ms []monkey) {
	var m monkey
	for s := bufio.NewScanner(r); s.Scan(); {
		switch {
		case strings.HasPrefix(s.Text(), "M"):
			m = monkey{}
		case strings.HasPrefix(s.Text(), "  S"):
			for _, v := range strings.Split(s.Text()[18:], ", ") {
				n, _ := strconv.Atoi(v)
				m.items = append(m.items, n)
			}
		case strings.HasPrefix(s.Text(), "  O"):
			if strings.HasSuffix(s.Text(), "old") {
				m.calc = func(n int) int { return n * n }
				break
			}
			n1, _ := strconv.Atoi(s.Text()[25:])
			if strings.HasPrefix(s.Text()[23:], "+") {
				m.calc = func(n int) int { return n + n1 }
			} else {
				m.calc = func(n int) int { return n * n1 }
			}
		case strings.HasPrefix(s.Text(), "  T"):
			var n1, n2, n3 int
			fmt.Sscanf(s.Text(), "  Test: divisible by %d", &n1)
			s.Scan()
			fmt.Sscanf(s.Text(), "    If true: throw to monkey %d", &n2)
			s.Scan()
			fmt.Sscanf(s.Text(), "    If false: throw to monkey %d", &n3)
			m.next = func(n int) int {
				if n%n1 == 0 {
					return n2
				}
				return n3
			}
		case s.Text() == "":
			ms = append(ms, m)
		}
	}
	return append(ms, m)
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
