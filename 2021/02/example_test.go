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

type cmd struct {
	dir dir
	stp int
}

func ExamplePartOne() {
	ps, dp := PartOne(strings.NewReader(input))
	fmt.Println(ps * dp)
	// Output:
	// 2039256
}

func ExamplePartTwo() {
	ps, dp := PartTwo(strings.NewReader(input))
	fmt.Println(ps * dp)
	// Output:
	// 1856459736
}

func TestPartOne(t *testing.T) {
	ps, dp := PartOne(strings.NewReader(input_test))
	got, want := ps*dp, 150
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	ps, dp := PartTwo(strings.NewReader(input_test))
	got, want := ps*dp, 900
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func PartOne(r io.Reader) (ps int, dp int) {
	for _, in := range split(r) {
		switch in.dir {
		case dirUp:
			dp -= in.stp
		case dirDwn:
			dp += in.stp
		case dirFwd:
			ps += in.stp
		}
	}
	return ps, dp
}

func PartTwo(r io.Reader) (ps int, dp int) {
	var aim int
	for _, in := range split(r) {
		switch in.dir {
		case dirUp:
			aim -= in.stp
		case dirDwn:
			aim += in.stp
		case dirFwd:
			ps += in.stp
			dp += aim * in.stp
		}
	}
	return ps, dp
}

func split(r io.Reader) (vs []cmd) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		fs := strings.Fields(sc.Text())
		n, err := strconv.Atoi(fs[1])
		if err != nil {
			log.Fatal(err)
		}
		vs = append(vs, cmd{(dir)(fs[0]), n})
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	return vs
}
