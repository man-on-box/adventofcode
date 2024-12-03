package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	startStr = "mul("
	endStr   = ")"
	dontMod  = "don't()"
	doMod    = "do()"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(content))

	pt1Total := 0
	pt2Total := 0

	// part 2
	mulEnabled := true

	searchIndex := 0
	for searchIndex >= 0 {
		subStr := input[searchIndex:]
		startIndex := strings.Index(subStr, startStr)

		if startIndex >= 0 {
			if mulEnabled {
				disableMul := strings.Contains(subStr[:startIndex], dontMod)
				mulEnabled = !disableMul
			} else {
				enableMul := strings.Contains(subStr[:startIndex], doMod)
				mulEnabled = enableMul
			}
			endIndex := strings.Index(subStr[startIndex:], ")")

			if endIndex > 0 {
				potentialParams := subStr[startIndex+len(startStr) : startIndex+endIndex]
				isValid, params := parseParams(potentialParams)
				if isValid {
					pt1Total += params[0] * params[1]
				}
				if isValid && mulEnabled {
					pt2Total += params[0] * params[1]
				}
			}
			// increment search index with addition
			// so we move to the next instance
			searchIndex += startIndex + 4
		} else {
			// no more instances
			searchIndex = -1
		}
	}

	// Part 1 answer: 185797128
	fmt.Println("Part 1 total:", pt1Total)
	// Part 2 answer: 89798695
	fmt.Println("Part 2 total:", pt2Total)
}

// takes a slice of strings and parses the params
func parseParams(str string) (isValid bool, params []int) {
	if len(str) > len("xxx,xxx") {
		return false, nil
	}

	parts := strings.Split(str, ",")
	if len(parts) != 2 {
		return false, nil
	}

	paramValues := make([]int, 2)

	for i, v := range parts {
		value, err := strconv.Atoi(v)
		if err != nil {
			return false, nil
		}
		paramValues[i] = value
	}

	return true, paramValues
}
