package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

import "fmt"

func mass(m int) int {
	return int(m/3) - 2
}

func sum(moduleMasses []int) int {
	i := 0
	total := 0
	for i < len(moduleMasses) {
		total += moduleMasses[i]
		i += 1
	}
	return total
}

func readInputs() string {
	if dat, err := ioutil.ReadFile("/tmp/aoc1"); err == nil {
		return string(dat)
	} else {
		return "oops"
	}
}

func splitIntoModules(inputs string) []int {
	var lines []string = strings.Split(inputs, "\n")
	modules := make([]int, len(lines))
	i := 0
	for i < len(lines) {
		m, _ := strconv.Atoi(lines[i])
		modules[i] = m
		i += 1
	}
	return modules
}

func convertToMasses(modules []int) []int {
	masses := make([]int, len(modules))
	i := 0
	for i < len(modules) {
		masses[i] = mass(modules[i])
		i += 1
	}
	return masses
}

var inputs [3]int = [3]int{3, 4, 5}

func main() {
	fmt.Println(sum(convertToMasses(splitIntoModules(readInputs()))))
}
