// https://adventofcode.com/2021/day/16
package d16_test

import (
	_ "embed"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

const (
	LitValTypeId = 4
)

type stack[T any] []T

func (s stack[T]) empty() bool {
	return len(s) == 0
}

func (s stack[T]) push(v ...T) stack[T] {
	return append(s, v...)
}

func (s stack[T]) pop() (stack[T], T, bool) {
	if len(s) == 0 {
		return s, *new(T), false
	}
	return s[:len(s)-1], s[len(s)-1], true
}

type head struct {
	version       int
	operator      bool
	lenInBits     int
	numOfPackages int
}

type item struct {
	startFrom     int
	lenInBits     int
	numOfPackages int
}

func ExamplePartOne() {
	fmt.Println(sum(PartOne(strings.NewReader(input))))
	// Output:
	// 908
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 0
}

func TestPartOne(t *testing.T) {
	for i, tc := range []struct {
		in   string
		want []int
	}{
		{
			in:   inputTest0,
			want: []int{6},
		},
		{
			in:   inputTest1,
			want: []int{1, 6, 2},
		},
		{
			in:   inputTest2,
			want: []int{7, 2, 4, 1},
		},
		{
			in:   inputTest3,
			want: []int{4, 1, 5, 6},
		},
		{
			in:   inputTest4,
			want: []int{3, 0, 0, 5, 1, 0, 3},
		},
		{
			in:   inputTest5,
			want: []int{6, 0, 0, 6, 4, 7, 0},
		},
		{
			in:   inputTest6,
			want: []int{5, 1, 3, 7, 6, 5, 2, 2},
		},
	} {
		t.Run(fmt.Sprintf("inputTest%d", i), func(t *testing.T) {
			got := PartOne(strings.NewReader(tc.in))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest0))
	want := 0
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (vs []int) {
	var ps int
	parse(scan(r), func(hd head) {
		vs = append(vs, hd.version)
	}, &ps)
	return vs
}

func PartTwo(r io.Reader) int {
	return 0
}

func parse(bs []bool, fn func(hd head), ps *int) (rs []bool) {
	for st := (stack[item]{{*ps, 0, 1}}); !st.empty(); {
		var it item
		st, it, _ = st.pop()
		if *ps != it.startFrom {
			if it.lenInBits > *ps-it.startFrom {
				st = st.push(item{*ps, it.lenInBits - *ps + it.startFrom, 0})
			} else if it.numOfPackages >= 2 {
				st = st.push(item{*ps, 0, it.numOfPackages - 1})
			}
			continue
		}
		var hd head
		hd, bs = parseHeader(bs, ps)
		if fn(hd); !hd.operator {
			_, bs = parseLiteral(bs, ps)
			st = st.push(it)
			continue
		}
		st = st.push(it, item{*ps, hd.lenInBits, hd.numOfPackages})
	}
	return bs
}

func parseHeader(bs []bool, ps *int) (hd head, rs []bool) {
	if n := btoi(bs[3:6]); n == LitValTypeId {
		*ps += 6
		return head{btoi(bs[:3]), false, 0, 0}, bs[6:]
	}
	if !bs[6] /* total length in bits */ {
		*ps += 22
		return head{btoi(bs[:3]), true, btoi(bs[7:22]), 0}, bs[22:]
	}
	*ps += 18
	return head{btoi(bs[:3]), true, 0, btoi(bs[7:18])}, bs[18:]
}

func parseLiteral(bs []bool, ps *int) (n int, rs []bool) {
	var l int
	for ; ; l += 5 {
		n = (n << 4) | btoi(bs[l+1:l+5])
		if !bs[l] {
			break
		}
	}
	*ps += l + 5
	return n, bs[l+5:]
}

func btoi(bs []bool) (n int) {
	for _, b := range bs {
		if n <<= 1; b {
			n |= 1
		}
	}
	return n
}

func sum(ns []int) (rs int) {
	for _, n := range ns {
		rs += n
	}
	return rs
}

func scan(r io.Reader) (bs []bool) {
	for {
		var n uint8
		if _, err := fmt.Fscanf(r, "%2X", &n); err == io.EOF {
			break
		}
		for i := 8; i != 0; i-- {
			bs = append(bs, (n>>(i-1))&1 == 1)
		}
	}
	return bs
}

const (
	inputTest0 = `D2FE28`
	inputTest1 = `38006F45291200`
	inputTest2 = `EE00D40C823060`
	inputTest3 = `8A004A801A8002F478`
	inputTest4 = `620080001611562C8802118E34`
	inputTest5 = `C0015000016115A2E0802F182340`
	inputTest6 = `A0016C880162017C3686B18A3D4780`
)
