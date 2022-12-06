// https://adventofcode.com/2022/day/06
package d06_test

import (
	_ "embed"
	"fmt"
	"io"
	"strings"
	"testing"
)

//go:embed testdata/input
var input string

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 1953
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 2301
}

func TestPartOne(t *testing.T) {
	for i, tc := range []struct {
		in   string
		want int
	}{
		{
			in:   inputTest0,
			want: 7,
		},
		{
			in:   inputTest1,
			want: 5,
		},
		{
			in:   inputTest2,
			want: 6,
		},
		{
			in:   inputTest3,
			want: 10,
		},
		{
			in:   inputTest4,
			want: 11,
		},
	} {
		t.Run(fmt.Sprintf("inputTest%d", i), func(t *testing.T) {
			got := PartOne(strings.NewReader(tc.in))
			if got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
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
			in:   inputTest0,
			want: 19,
		},
		{
			in:   inputTest1,
			want: 23,
		},
		{
			in:   inputTest2,
			want: 23,
		},
		{
			in:   inputTest3,
			want: 29,
		},
		{
			in:   inputTest4,
			want: 26,
		},
	} {
		t.Run(fmt.Sprintf("inputTest%d", i), func(t *testing.T) {
			got := PartTwo(strings.NewReader(tc.in))
			if got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}
}

func PartOne(r io.Reader) int {
	return look(scan(r), 4)
}

func PartTwo(r io.Reader) int {
	return look(scan(r), 14)
}

func look(cs []byte, n int) int {
	bf := make([]int, 'z'-'a'+1)
	for i := range cs {
		if i < n-1 {
			bf[cs[i]]++
			continue
		}
		if i != n-1 {
			bf[cs[i-n]]--
		}
		bf[cs[i]]++
		if fold(bf, cs[i-n+1:i+1]) == 1 {
			return i + 1
		}
	}
	return 0
}

func fold(ns []int, cs []byte) (n int) {
	n = ns[cs[0]]
	for _, c := range cs[1:] {
		n *= ns[c]
	}
	return n
}

func scan(r io.Reader) (cs []byte) {
	bs, _ := io.ReadAll(r)
	cs = append(cs, bs...)
	for i := range cs {
		cs[i] = cs[i] - 'a'
	}
	return cs
}

const (
	inputTest0 = `mjqjpqmgbljsphdztnvjfqwrcgsmlb`
	inputTest1 = `bvwbjplbgvbhsrlpgdmjqwftvncz`
	inputTest2 = `nppdvjthqldpwncqszvftbrmjlhg`
	inputTest3 = `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`
	inputTest4 = `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`
)
