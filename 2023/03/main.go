package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	value     int
	pos       [2]int
	boxValues []string
}

type Contents [][]string

/*
	This one is quite a mess, but it works.
*/

func main() {
	input, _ := os.ReadFile("input.txt")
	contents := Contents{}
	for _, row := range strings.Split(string(input), "\n") {
		contents = append(contents, strings.Split(row, ""))
	}
	numbers := []Number{}
	stars := [][2]int{}
	for i := range contents {
		parseRow(&contents, &numbers, &stars, i)
	}
	var partNumbers []int
	for _, n := range numbers {
		if n.isPartNo() {
			partNumbers = append(partNumbers, n.value)
		}
	}

	// Part 1
	// Example result = 4361
	// 521515 < after fixing sliceEnd value
	fmt.Println("Total", sum(partNumbers))

	// Part 2
	// fmt.Println("Stars", stars)
	gears := getGears(&stars, &numbers)
	// fmt.Println("Gears", gears)
	gearRatios := []int{}
	for _, g := range gears {
		gearRatios = append(gearRatios, g[0]*g[1])
	}

	fmt.Println("Gear ratio totals", sum(gearRatios))
}

/*
Get lists of intersections, per star:
[[467, 35], [617], [755, 598]]
*/
func getGears(stars *[][2]int, numbers *[]Number) [][]int {
	intersections := [][]int{}
	for _, pos := range *stars {
		gears := []int{}
		for _, n := range *numbers {
			// Check Y intercepts
			yDistance := pos[1] - n.pos[1]
			if yDistance <= 1 && yDistance >= -1 {
				// Check X intercepts
				xDistance := pos[0] - n.pos[0]
				if xDistance >= -1 && xDistance <= len(strconv.Itoa(n.value)) {
					gears = append(gears, n.value)
				}
			}
		}
		if len(gears) == 2 {
			intersections = append(intersections, gears)
		}
	}
	return intersections
}

func parseRow(contents *Contents, numbers *[]Number, stars *[][2]int, lineIndex int) {
	row := (*contents)[lineIndex]
	pos := [2]int{}
	digitCache := ""

	for i, char := range row {
		isEoL := i == len(row)-1
		if char == "*" {
			*stars = append(*stars, [2]int{i, lineIndex})
		}
		if isDigit(char) {
			if digitCache == "" {
				pos = [2]int{i, lineIndex}
			}
			digitCache += char
			if isEoL || !isDigit(row[i+1]) {
				value, _ := strconv.Atoi(digitCache)
				boxValues := getBoxValues(contents, lineIndex, pos[1], i)

				*numbers = append(*numbers, Number{value, pos, boxValues})
				digitCache = ""
			}
		}
	}
}

func getBoxValues(contents *Contents, lineIndex int, startIndex int, endIndex int) []string {
	var boxValues = []string{}
	row := (*contents)[lineIndex]
	sliceStart := max(startIndex-1, 0)
	sliceEnd := min(endIndex+1, len(row)-1)

	// Get above line values
	if lineIndex != 0 {
		line := (*contents)[lineIndex-1]
		for i := sliceStart; i <= sliceEnd; i++ {
			boxValues = append(boxValues, line[i])
		}
	}
	// Get current line values
	if startIndex != 0 {
		boxValues = append(boxValues, row[sliceStart])
	}

	if endIndex != len(row)-1 {
		boxValues = append(boxValues, row[sliceEnd])
	}

	// Get below line values
	if lineIndex != len(*contents)-1 {
		line := (*contents)[lineIndex+1]
		for i := sliceStart; i <= sliceEnd; i++ {
			boxValues = append(boxValues, line[i])
		}
	}
	return boxValues
}

func (n Number) isPartNo() bool {
	result := false
	for _, v := range n.boxValues {
		if isSymbol(v) {
			result = true
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var digits = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
var nonSymbols = append(digits, ".")

func isDigit(char string) bool {
	result := false
	for _, d := range digits {
		if d == char {
			result = true
		}
	}
	return result
}

func isSymbol(char string) bool {
	result := true
	for _, c := range nonSymbols {
		if c == char {
			result = false
		}
	}
	return result
}

func sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
