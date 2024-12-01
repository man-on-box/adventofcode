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
	lines := strings.Split(string(content), "\n")
	noOfValues := len(lines) - 1
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
		values := strings.Split(l, "   ")
		leftValue, _ := strconv.Atoi(values[0])
		rightValue, _ := strconv.Atoi(values[1])

		// part 1
		leftArr[i] = leftValue
		rightArr[i] = rightValue

		// part 2
		count, has := rightMap[rightValue]
		if has {
			rightMap[rightValue] = count + 1
		} else {
			rightMap[rightValue] = 1
		}
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
		count, has := rightMap[leftValue]
		if has {
			score := leftValue * count
			similarityScore += score
		}

	}

	// Answer part 1: 2378066
	fmt.Println(diffTotal)

	// Answer part 2:
	fmt.Println(similarityScore)
}
