package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readInputs(filename string) []int {
	if dat, err := ioutil.ReadFile(filename); err == nil {
		numStrs := strings.Split(string(dat), ",")
		var opcodes = make([]int, len(numStrs))
		for i, s := range numStrs {
			n, _ := strconv.Atoi(s)
			opcodes[i] = n
		}
		return opcodes
	} else {
		return []int{}
	}
}

func op1(state []int, p int) []int {
	var p1, p2, p3 int
	p1 = p + 1
	p2 = p + 2
	p3 = p + 3
	output := state[state[p1]] + state[state[p2]]
	outputAddr := state[p3]
	state[outputAddr] = output
	if output == 99 {
		fmt.Println("generated 99")
	}
	return state
}

func op2(state []int, p int) []int {
	var p1, p2, p3 int
	p1 = p + 1
	p2 = p + 2
	p3 = p + 3
	output := state[state[p1]] * state[state[p2]]
	outputAddr := state[p3]
	state[outputAddr] = output
	if output == 99 {
		fmt.Println("generated 99")
	}
	return state
}

func run(state []int) []int {
	position := 0
	op := state[position]
	var failure []int
	for op != 99 {
		switch {
		case op == 1:
			op1(state, position)
		case op == 2:
			op2(state, position)
		default:
			return failure
		}
		position += 4
		op = state[position]
	}
	return state
}

func main() {
	var s []int = readInputs("/tmp/aoc2")
	s[1] = 12
	s[2] = 2
	fmt.Println(run(s)[0])
}
