// https://adventofcode.com/2022/day/07
package d07_test

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

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// 1443806
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// 942298
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := 95437
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest))
	want := 24933642
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func PartOne(r io.Reader) (n int) {
	for _, size := range scan(r) {
		if size <= 100000 {
			n += size
		}
	}
	return n
}

func PartTwo(r io.Reader) (n int) {
	ss := scan(r)
	lm := ss[len(ss)-1] - 40_000_000
	for _, size := range ss {
		if size >= lm && (n == 0 || n > size) {
			n = size
		}
	}
	return n
}

func scan(r io.Reader) (ss []int) {
	st, n1, n2 := stack[int]{}, 0, 0
	for s := bufio.NewScanner(r); s.Scan(); {
		if s.Text() == "$ ls" ||
			strings.HasPrefix(s.Text(), "dir ") {
			continue // ignore
		}
		if s.Text() == "$ cd .." {
			st, n1, _ = st.pop()
			st, n2, _ = st.pop()
			st = st.push(n1 + n2)
			ss = append(ss, n1)
			continue
		}
		if strings.HasPrefix(s.Text(), "$ cd ") {
			st = st.push(0)
			continue
		}
		st, n1, _ = st.pop()
		fmt.Sscanf(s.Text(), "%d", &n2)
		st = st.push(n1 + n2)
	}
	for !st.empty() {
		st, n1, _ = st.pop()
		if !st.empty() {
			st, n2, _ = st.pop()
			st = st.push(n1 + n2)
		}
		ss = append(ss, n1)
	}
	return ss
}

const inputTest = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`
