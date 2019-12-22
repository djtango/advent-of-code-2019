package main

import (
	"fmt"
	"strconv"
)

const input string = "138307-654504"
const lowerBound string = "138307"
const upperBound string = "654504"

func inspect(s string) bool {
	l := len(s)
	i, j := 0, 1
	allNotDecreasing := true
	hasAdjacentValue := false
	counts := map[string]int{
		"0": 0,
		"1": 0,
		"2": 0,
		"3": 0,
		"4": 0,
		"5": 0,
		"6": 0,
		"7": 0,
		"8": 0,
		"9": 0,
	}
	for j < l {
		counts[string(s[i])] += 1
		allNotDecreasing = allNotDecreasing && s[i] <= s[j]
		i += 1
		j += 1
	}
	counts[string(s[i])] += 1
	for _, v := range counts {
		hasAdjacentValue = hasAdjacentValue || v == 2
	}
	return allNotDecreasing && hasAdjacentValue
}

func checkRange(lower, upper int) int {
	i := 138307
	end := 654504
	count := 0
	for i <= end {
		if inspect(strconv.Itoa(i)) {
			count += 1
		}
		i += 1
	}
	return count
}

func main() {
	fmt.Println("--------------------")
	fmt.Println(checkRange(138307, 654504))
	fmt.Println("--------------------")
}
