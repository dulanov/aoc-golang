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

type typ int

const (
	TypeIdOpSum typ = iota
	TypeIdOpMul
	TypeIdOpMin
	TypeIdOpMax
	TypeIdVlLit
	TypeIdOpGt
	TypeIdOpLt
	TypeIdOpEq
)

var ops = []func(ns []int) int{
	sum, mul, min, max,
	func(ns []int) int { return ns[0] },
	func(ns []int) int { return gt(ns[0], ns[1]) },
	func(ns []int) int { return lt(ns[0], ns[1]) },
	func(ns []int) int { return eq(ns[0], ns[1]) },
}

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
	opType        typ
	posLimit      int
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
	// 10626195124371
}

func TestPartOne(t *testing.T) {
	for i, tc := range []struct {
		in   string
		want []int
	}{
		{
			in:   inputTest00,
			want: []int{6},
		},
		{
			in:   inputTest01,
			want: []int{1, 6, 2},
		},
		{
			in:   inputTest02,
			want: []int{7, 2, 4, 1},
		},
		{
			in:   inputTest03,
			want: []int{4, 1, 5, 6},
		},
		{
			in:   inputTest04,
			want: []int{3, 0, 0, 5, 1, 0, 3},
		},
		{
			in:   inputTest05,
			want: []int{6, 0, 0, 6, 4, 7, 0},
		},
		{
			in:   inputTest06,
			want: []int{5, 1, 3, 7, 6, 5, 2, 2},
		},
	} {
		t.Run(fmt.Sprintf("inputTest0%d", i), func(t *testing.T) {
			got := PartOne(strings.NewReader(tc.in))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	for i, tc := range []struct {
		in   string
		want int
	}{
		{
			in:   inputTest10,
			want: 3,
		},
		{
			in:   inputTest11,
			want: 54,
		},
		{
			in:   inputTest12,
			want: 7,
		},
		{
			in:   inputTest13,
			want: 9,
		},
		{
			in:   inputTest14,
			want: 1,
		},
		{
			in:   inputTest15,
			want: 0,
		},
		{
			in:   inputTest16,
			want: 0,
		},
		{
			in:   inputTest17,
			want: 1,
		},
	} {
		t.Run(fmt.Sprintf("inputTest1%d", i), func(t *testing.T) {
			got := PartTwo(strings.NewReader(tc.in))
			if got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}
}

func PartOne(r io.Reader) (vs []int) {
	eval(scan(r), func(hd head) {
		vs = append(vs, hd.version)
	})
	return vs
}

func PartTwo(r io.Reader) (n int) {
	return eval(scan(r), func(hd head) {})
}

func eval(bs []bool, fn func(hd head)) (n int) {
	type op struct {
		op     typ
		limit  int
		values []int
	}
	for pc, st := 0, (stack[op]{op{limit: 1}}); ; {
		var o1, o2 op
		st, o1, _ = st.pop()
		if o1.op == TypeIdVlLit && st.empty() {
			n = ops[o1.op](o1.values)
			break
		}
		if o1.op == TypeIdVlLit {
			st, o2, _ = st.pop()
			o2.values = append(o2.values, o1.values...)
			if o2.limit != pc && o2.limit != len(o2.values) {
				st = st.push(o2)
				continue
			}
			st = st.push(op{op: TypeIdVlLit,
				values: []int{ops[o2.op](o2.values)}})
			continue
		}
		var hd head
		hd, bs = parseHeader(bs, &pc)
		if fn(hd); hd.opType == TypeIdVlLit {
			var n int
			n, bs = parseLiteral(bs, &pc)
			st = st.push(o1, op{op: TypeIdVlLit, values: []int{n}})
			continue
		}
		st = st.push(o1, op{hd.opType, hd.posLimit | hd.numOfPackages, nil})
	}
	return n
}

func parseHeader(bs []bool, pc *int) (hd head, rs []bool) {
	id := typ(btoi(bs[3:6]))
	if id == TypeIdVlLit {
		*pc += 6
		return head{btoi(bs[:3]), TypeIdVlLit, 0, 0}, bs[6:]
	}
	if !bs[6] /* total length in bits */ {
		*pc += 22
		return head{btoi(bs[:3]), id, *pc + btoi(bs[7:22]), 0}, bs[22:]
	}
	*pc += 18
	return head{btoi(bs[:3]), id, 0, btoi(bs[7:18])}, bs[18:]
}

func parseLiteral(bs []bool, pc *int) (n int, rs []bool) {
	var l int
	for ; ; l += 5 {
		n = (n << 4) | btoi(bs[l+1:l+5])
		if !bs[l] {
			break
		}
	}
	*pc += l + 5
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

func sum(ns []int) (n int) {
	for _, v := range ns {
		n += v
	}
	return n
}

func mul(ns []int) (n int) {
	n = 1
	for _, v := range ns {
		n *= v
	}
	return n
}

func min(ns []int) (n int) {
	n = ns[0]
	for _, v := range ns[1:] {
		if n > v {
			n = v
		}
	}
	return n
}

func max(ns []int) (n int) {
	n = ns[0]
	for _, v := range ns[1:] {
		if n < v {
			n = v
		}
	}
	return n
}

func gt(n1, n2 int) (n int) {
	if n1 > n2 {
		return 1
	}
	return 0
}

func lt(n1, n2 int) (n int) {
	if n1 < n2 {
		return 1
	}
	return 0
}

func eq(n1, n2 int) (n int) {
	if n1 == n2 {
		return 1
	}
	return 0
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
	inputTest00 = `D2FE28`
	inputTest01 = `38006F45291200`
	inputTest02 = `EE00D40C823060`
	inputTest03 = `8A004A801A8002F478`
	inputTest04 = `620080001611562C8802118E34`
	inputTest05 = `C0015000016115A2E0802F182340`
	inputTest06 = `A0016C880162017C3686B18A3D4780`

	inputTest10 = `C200B40A82`
	inputTest11 = `04005AC33890`
	inputTest12 = `880086C3E88112`
	inputTest13 = `CE00C43D881120`
	inputTest14 = `D8005AC2A8F0`
	inputTest15 = `F600BC2D8F`
	inputTest16 = `9C005AC2F8F0`
	inputTest17 = `9C0141080250320F1802104A08`
)
