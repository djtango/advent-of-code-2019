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
	hasAdjacentDigits := false
	allNotDecreasing := true
	for j < l {
		hasAdjacentDigits = hasAdjacentDigits || s[i] == s[j]
		allNotDecreasing = allNotDecreasing && s[i] <= s[j]
		i += 1
		j += 1
	}
	return hasAdjacentDigits && allNotDecreasing
}

func main() {
	fmt.Println("--------------------")
	i := 138307
	end := 654504
	count := 0
	for i <= end {
		if inspect(strconv.Itoa(i)) {
			count += 1
		}
		i += 1
	}
	fmt.Println(count)
	fmt.Println("--------------------")
}
