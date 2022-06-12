// https://adventofcode.com/2021/day/1
package day1_test

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func numberOfTimes(r io.Reader, fn func(int, int) bool) (n int) {
	scnr := bufio.NewScanner(r)
	scnr.Split(bufio.ScanLines)
	if !scnr.Scan() {
		return
	}
	i, err := strconv.Atoi(scnr.Text())
	if err != nil {
		log.Fatal(err)
	}
	for scnr.Scan() {
		j, err := strconv.Atoi(scnr.Text())
		if err != nil {
			log.Fatal(err)
		}
		if fn(i, j) {
			n++
		}
		i = j
	}
	return
}

func Example() {
	file, err := os.Open(filepath.Join("testdata", "input"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println(numberOfTimes(file,
		func(i, j int) bool { return j > i }))
	// Output:
	// 1342
}

func TestNumberOfTimes(t *testing.T) {
	const in = "199\n200\n208\n210\n200\n207\n240\n269\n260\n263"
	got := numberOfTimes(strings.NewReader(in),
		func(i, j int) bool { return j > i })
	want := 7
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
