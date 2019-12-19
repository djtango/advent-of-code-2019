package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Navigation struct {
	Θ string
	R int
}

type Input struct {
	Wire1 []Navigation
	Wire2 []Navigation
}

func parseNavs(line string) []Navigation {
	instructions := strings.Split(line, ",")
	var navs = make([]Navigation, len(instructions))
	for i, nav := range instructions {
		θ := string(nav[0])
		r, _ := strconv.Atoi(string(nav[1:]))
		navs[i] = Navigation{θ, r}
	}
	return navs
}

func readInputs(filename string) Input {
	if dat, err := ioutil.ReadFile(filename); err == nil {
		lines := strings.Split(string(dat), "\n")
		wire1 := parseNavs(lines[0])
		wire2 := parseNavs(lines[1])
		return Input{wire1, wire2}
	} else {
		empty := make([]Navigation, 0)
		return Input{empty, empty}
	}
}

func main() {
	fmt.Println(readInputs("/tmp/aoc3"))
}
