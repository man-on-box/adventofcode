package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func evaluate(target int, values []int, concat bool) bool {
	// BASE CASES
	// if we have just one value left, and it matches
	// the target then this is a valid equation
	if len(values) == 1 {
		return values[0] == target
	}
	// if we are already above the target, lets quit early
	if values[0] > target {
		return false
	}
	// if we do an add and is true, we go true, and go down this branch
	if evaluate(target, slices.Concat([]int{values[0] + values[1]}, values[2:]), concat) {
		return true
	}
	// as well, we branch into multiplying so long as we're true
	if evaluate(target, slices.Concat([]int{values[0] * values[1]}, values[2:]), concat) {
		return true
	}
	// for part two, we'll try concatenating also with a 3rd branch
	if concat {
		concatVal, _ := strconv.Atoi(strconv.Itoa(values[0]) + strconv.Itoa(values[1]))
		if evaluate(target, slices.Concat([]int{concatVal}, values[2:]), concat) {
			return true
		}
	}
	return false
}

func main() {
	f, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(f))
	lines := strings.Split(input, "\n")
	pt1 := 0
	pt2 := 0

	for _, l := range lines {
		parts := strings.Split(l, ":")
		target, _ := strconv.Atoi(parts[0])
		strValues := strings.Fields(parts[1])
		values := make([]int, len(strValues))
		for i, v := range strValues {
			intValue, _ := strconv.Atoi(v)
			values[i] = intValue
		}
		if evaluate(target, values, false) {
			pt1 += target
		}
		if evaluate(target, values, true) {
			pt2 += target
		}

	}

	fmt.Println(pt1)
	fmt.Println(pt2)
}
