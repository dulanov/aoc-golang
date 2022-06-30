// https://adventofcode.com/2021/day/2
package d02_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"
)

const input_test = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

//go:embed testdata/input
var input string

type dir string

const (
	dirUp  dir = "up"
	dirDwn dir = "down"
	dirFwd dir = "forward"
)

type op struct {
	dir  dir
	step int
}

func ExamplePartOne() {
	ps := PartOne(strings.NewReader(input))
	fmt.Println(ps.h * ps.v)
	// Output:
	// 2039256
}

func ExamplePartTwo() {
	ps := PartTwo(strings.NewReader(input))
	fmt.Println(ps.h * ps.v)
	// Output:
	// 1856459736
}

func TestPartOne(t *testing.T) {
	got := PartOne(strings.NewReader(input_test))
	want := struct{ h, v int }{15, 10}
	if got != want {
		t.Errorf("got %+v; want %+v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(input_test))
	want := struct{ h, v int }{15, 60}
	if got != want {
		t.Errorf("got %+v; want %+v", got, want)
	}
}

func PartOne(r io.Reader) (pos struct{ h, v int }) {
	for _, o := range scan(r) {
		switch o.dir {
		case dirUp:
			pos.v -= o.step
		case dirDwn:
			pos.v += o.step
		case dirFwd:
			pos.h += o.step
		}
	}
	return pos
}

func PartTwo(r io.Reader) (pos struct{ h, v int }) {
	var aim int
	for _, in := range scan(r) {
		switch in.dir {
		case dirUp:
			aim -= in.step
		case dirDwn:
			aim += in.step
		case dirFwd:
			pos.h += in.step
			pos.v += aim * in.step
		}
	}
	return pos
}

func scan(r io.Reader) (ops []op) {
	for s := bufio.NewScanner(r); s.Scan(); {
		fs := strings.Fields(s.Text())
		n, _ := strconv.Atoi(fs[1])
		ops = append(ops, op{(dir)(fs[0]), n})
	}
	return ops
}
