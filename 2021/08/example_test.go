// https://adventofcode.com/2021/day/8
package d08_test

import (
	"bufio"
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
	// 272
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 1007675
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(input_test))
	want := 26
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(input_test))
	want := 61229
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (n int) {
	for _, e := range scan(r) {
		for _, s := range e[1] {
			if len(s)%5 >= 2 {
				n++
			}
		}
	}
	return n
}

func PartTwo(r io.Reader) (n int) {
	for _, e := range scan(r) {
		m, v := demap(e[0]), 0
		for _, s := range e[1] {
			v = v*10 + decode(s, m)
		}
		n += v
	}
	return n
}

func demap(ps []string) (m [3]byte) {
	var cs [7]int
	for _, s := range ps {
		for _, r := range s {
			cs[r-'a']++
		}
	}
	for i, c := range cs {
		switch c {
		case 8 /* 'a|c' */ :
			if m[0] == 0 {
				m[0] = (byte)('a' + i)
			} else {
				m[1] = (byte)('a' + i)
			}
		case 4 /* 'e' */ :
			m[2] = (byte)('a' + i)
		}
	}
	return m
}

func decode(s string, m [3]byte) (n int) {
	switch len(s) {
	case 2:
		n = 1
	case 3:
		n = 7
	case 4:
		n = 4
	case 5:
		if strings.IndexByte(s, m[0] /* 'a|c' */) == -1 ||
			strings.IndexByte(s, m[1] /* 'a|c' */) == -1 {
			n = 5
		} else if strings.IndexByte(s, m[2] /* 'e' */) == -1 {
			n = 3
		} else {
			n = 2
		}
	case 6:
		if strings.IndexByte(s, m[0] /* 'a|c' */) == -1 ||
			strings.IndexByte(s, m[1] /* 'a|c' */) == -1 {
			n = 6
		} else if strings.IndexByte(s, m[2] /* 'e' */) == -1 {
			n = 9
		} else {
			n = 0
		}
	case 7:
		n = 8
	}
	return n
}

func scan(r io.Reader) (entries [][2][]string) {
	for s := bufio.NewScanner(r); s.Scan(); {
		sl := strings.Split(s.Text(), "|")
		entries = append(entries, [2][]string{
			strings.Fields(sl[0]), strings.Fields(sl[1])})
	}
	return entries
}

const input_test = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`
