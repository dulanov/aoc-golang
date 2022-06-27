// https://adventofcode.com/2021/day/2
package d02_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"log"
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

type pos struct {
	h int
	v int
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
	want := pos{15, 10}
	if got != want {
		t.Errorf("got %+v; want %+v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(strings.NewReader(input_test))
	want := pos{15, 60}
	if got != want {
		t.Errorf("got %+v; want %+v", got, want)
	}
}

func PartOne(r io.Reader) (p pos) {
	for _, o := range split(r) {
		switch o.dir {
		case dirUp:
			p.v -= o.step
		case dirDwn:
			p.v += o.step
		case dirFwd:
			p.h += o.step
		}
	}
	return p
}

func PartTwo(r io.Reader) (p pos) {
	var aim int
	for _, in := range split(r) {
		switch in.dir {
		case dirUp:
			aim -= in.step
		case dirDwn:
			aim += in.step
		case dirFwd:
			p.h += in.step
			p.v += aim * in.step
		}
	}
	return p
}

func split(r io.Reader) (ops []op) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		fs := strings.Fields(sc.Text())
		n, err := strconv.Atoi(fs[1])
		if err != nil {
			log.Fatal(err)
		}
		ops = append(ops, op{(dir)(fs[0]), n})
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	return ops
}
