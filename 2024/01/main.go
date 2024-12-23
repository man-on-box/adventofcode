package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(content))
	lines := strings.Split(input, "\n")
	noOfValues := len(lines)
	leftArr := make([]int, noOfValues)
	rightArr := make([]int, noOfValues)
	diffArr := make([]int, noOfValues)

	// part 2
	rightMap := map[int]int{}
	similarityScore := 0

	for i, l := range lines {
		if len(l) == 0 {
			break
		}
		values := strings.Fields(l)
		leftValue, _ := strconv.Atoi(values[0])
		rightValue, _ := strconv.Atoi(values[1])

		// part 1
		leftArr[i] = leftValue
		rightArr[i] = rightValue

		// part 2
		rightMap[rightValue]++
	}

	slices.Sort(leftArr)
	slices.Sort(rightArr)

	diffTotal := 0
	for i, leftValue := range leftArr {
		rightValue := rightArr[i]
		diff := leftValue - rightValue
		if diff < 0 {
			diff = -diff
		}
		diffTotal += diff
		diffArr[i] = diff

		// part 2
		similarityScore += leftValue * rightMap[leftValue]
	}

	// Answer part 1: 2378066
	fmt.Println("Part 1:", diffTotal)

	// Answer part 2: 18934359
	fmt.Println("Part 2:", similarityScore)
}
