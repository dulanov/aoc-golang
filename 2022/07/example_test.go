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

type dir struct {
	size    int
	entries []entry
}

type entry struct {
	name string
	size int
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
	ds := scan(r)
	calcSize(ds, []string{"/"})
	for _, d := range ds {
		if d.size <= 100000 {
			n += d.size
		}
	}
	return n
}

func PartTwo(r io.Reader) (n int) {
	ds := scan(r)
	lm := calcSize(ds, []string{"/"}) - 40_000_000
	for _, d := range ds {
		if d.size >= lm && (n == 0 || n > d.size) {
			n = d.size
		}
	}
	return n
}

func calcSize(ds map[string]dir, path []string) (n int) {
	str := strings.Join(path, "")
	if ds[str].size != 0 {
		return ds[str].size
	}
	for _, e := range ds[str].entries {
		if e.size != 0 {
			n += e.size
		} else {
			n += calcSize(ds, append(path, e.name))
		}
	}
	ds[str] = dir{size: n, entries: ds[str].entries}
	return n
}

func scan(r io.Reader) (ds map[string]dir) {
	var path []string
	ds = make(map[string]dir)
	for s := bufio.NewScanner(r); s.Scan(); {
		if s.Text() == "$ ls" {
			continue // ignore
		}
		if s.Text() == "$ cd .." {
			path = path[:len(path)-1]
			continue
		}
		if strings.HasPrefix(s.Text(), "$ cd ") {
			var str string
			fmt.Sscanf(s.Text(), "$ cd %s", &str)
			path = append(path, str)
			ds[strings.Join(path, "")] = dir{}
			continue
		}
		entry := entry{}
		if strings.HasPrefix(s.Text(), "dir ") {
			fmt.Sscanf(s.Text(), "dir %s", &entry.name)
		} else {
			fmt.Sscanf(s.Text(), "%d %s", &entry.size, &entry.name)
		}
		str := strings.Join(path, "")
		ds[str] = dir{entries: append(ds[str].entries, entry)}
	}
	return ds
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
