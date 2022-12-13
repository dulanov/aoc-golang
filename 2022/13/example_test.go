// https://adventofcode.com/2022/day/13
package d13_test

import (
	"bufio"
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

type packet struct {
	n  int
	ls []packet
}

func newPacket(ps ...packet) packet {
	return packet{n: -1, ls: ps}
}

func newPacketTerm(n int) packet {
	return packet{n: n}
}

func (p packet) term() bool {
	return p.n != -1
}

func (p packet) less(p2 packet) bool {
	if p.term() && p2.term() {
		return p.n < p2.n
	}
	if p.term() && !p2.term() {
		return newPacket(p).less(p2)
	}
	if !p.term() && p2.term() {
		return p.less(newPacket(p2))
	}
	for i := 0; i < len(p.ls) && i < len(p2.ls); i++ {
		if p2.ls[i].less(p.ls[i]) {
			return false
		}
		if p.ls[i].less(p2.ls[i]) {
			return true
		}
	}
	return len(p.ls) < len(p2.ls)
}

func ExamplePartOne() {
	ns := PartOne(strings.NewReader(input))
	fmt.Println(sum(ns))
	// Output:
	// 5682
}

func ExamplePartTwo() {
	ns := PartTwo(strings.NewReader(input))
	fmt.Println(ns[0] * ns[1])
	// Output:
	// 0
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := []int{1, 2, 4, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest))
	want := []int{10, 14}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func PartOne(r io.Reader) (ns []int) {
	for i, ps := 0, scan(r); i < len(ps); i += 2 {
		if ps[i].less(ps[i+1]) {
			ns = append(ns, i/2+1)
		}
	}
	return ns
}

func PartTwo(r io.Reader) (ns []int) {
	ps := scan(r)
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].less(ps[j])
	})
	for i, p := range []packet{
		newPacket(newPacket(newPacketTerm(2))),
		newPacket(newPacket(newPacketTerm(6)))} {
		ns = append(ns, sort.Search(len(ps), func(i int) bool {
			return p.less(ps[i])
		})+i+1)
	}
	return ns
}

func sum(ns []int) (n int) {
	for _, v := range ns {
		n += v
	}
	return n
}

func scan(r io.Reader) (ps []packet) {
	rpl := strings.NewReplacer(",", "", "10", ":")
	for s := bufio.NewScanner(r); s.Scan(); {
		if s.Text() == "" {
			continue
		}
		p, _ := parse([]byte(rpl.Replace(s.Text())))
		ps = append(ps, p)
	}
	return ps
}

func parse(bs []byte) (packet, int) {
	var ps []packet
	for i := 1; ; i++ {
		if bs[i] == ']' {
			return newPacket(ps...), i
		}
		if bs[i] == '[' {
			p, o := parse(bs[i:])
			i, ps = i+o, append(ps, p)
			continue
		}
		ps = append(ps, newPacketTerm(int(bs[i]-'0')))
	}
}

const inputTest = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`
