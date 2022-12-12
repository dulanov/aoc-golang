// https://adventofcode.com/2022/day/05
package d05_test

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
	if s.empty() {
		return s, *new(T), false
	}
	return s[:len(s)-1], s[len(s)-1], true
}

type op struct {
	nm, fr, to int
}

func ExamplePartOne() {
	fmt.Println(PartOne(strings.NewReader(input)))
	// Output:
	// CNSZFDVLJ
}

func ExamplePartTwo() {
	fmt.Println(PartTwo(strings.NewReader(input)))
	// Output:
	// QNDWLMGNS
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(inputTest))
	want := "CMZ"
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(inputTest))
	want := "MCD"
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func PartOne(r io.Reader) string {
	ss, ops := scan(r)
	for _, op := range ops {
		for i := 0; i < op.nm; i++ {
			var v byte
			ss[op.fr], v, _ = ss[op.fr].pop()
			ss[op.to] = ss[op.to].push(v)
		}
	}
	return str(ss)
}

func PartTwo(r io.Reader) string {
	ss, ops := scan(r)
	for _, op := range ops {
		buf := make([]byte, op.nm)
		for i := op.nm; i != 0; i-- {
			ss[op.fr], buf[i-1], _ = ss[op.fr].pop()
		}
		ss[op.to] = ss[op.to].push(buf...)
	}
	return str(ss)
}

func str(ss []stack[byte]) string {
	var b strings.Builder
	for _, st := range ss {
		if !st.empty() {
			_, v, _ := st.pop()
			b.WriteByte(v)
		}
	}
	return b.String()
}

func rev[T any](vs []T) {
	for i, j := 0, len(vs)-1; i < j; i, j = i+1, j-1 {
		vs[i], vs[j] = vs[j], vs[i]
	}
}

func scan(r io.Reader) (ss []stack[byte], ops []op) {
	for s := bufio.NewScanner(r); s.Scan(); {
		if strings.HasPrefix(s.Text(), " 1") {
			for _, st := range ss {
				rev(st)
			}
			s.Scan() /* skip empty line */
			continue
		}
		if strings.HasPrefix(s.Text(), "move ") {
			var op op
			fmt.Sscanf(s.Text(), "move %d from %d to %d", &op.nm, &op.fr, &op.to)
			op.fr, op.to = op.fr-1, op.to-1
			ops = append(ops, op)
			continue
		}
		if len(ss) == 0 {
			ss = make([]stack[byte], (len(s.Text())+1)/4)
		}
		for i := range ss {
			if s.Text()[i*4] == '[' {
				ss[i] = ss[i].push(s.Text()[i*4+1])
			}
		}
	}
	return ss, ops
}

const inputTest = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
