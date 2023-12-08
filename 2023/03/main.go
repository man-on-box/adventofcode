package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	value      int
	startIndex int
	boxValues  []string
}

func main() {
	input, _ := os.ReadFile("example.txt")
	rows := strings.Split(string(input), "\n")
	numbers := []Number{}
	for i := range rows {
		ns := parseNumbersFromLine(rows, i)
		numbers = append(numbers, ns...)
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
	fmt.Println(partNumbers)
	fmt.Println("Total", sum(partNumbers))
}

func parseNumbersFromLine(lines []string, lineIndex int) []Number {
	line := lines[lineIndex]
	splitLine := strings.Split(line, "")
	var matrices = []Number{}
	startIndex := 0
	digitCache := ""

	for i, char := range splitLine {
		isEoL := i == len(splitLine)-1
		if isDigit(char) {
			if digitCache == "" {
				startIndex = i
			}
			digitCache += char
			if isEoL || !isDigit(splitLine[i+1]) {
				value, _ := strconv.Atoi(digitCache)
				fmt.Println("Found:", value)
				boxValues := getBoxValues(lines, lineIndex, startIndex, i)
				fmt.Println("Box:", boxValues)

				matrices = append(matrices, Number{value, startIndex, boxValues})
				digitCache = ""
			}
		}
	}
	return matrices
}

func getBoxValues(lines []string, lineIndex int, startIndex int, endIndex int) []string {
	fmt.Println("Line:", lineIndex, "Start:", startIndex, "End:", endIndex)
	var boxValues = []string{}
	currentLine := strings.Split(lines[lineIndex], "")
	sliceStart := max(startIndex-1, 0)
	sliceEnd := min(endIndex+1, len(currentLine)-1)

	// Get above line values
	if lineIndex != 0 {
		line := strings.Split(lines[lineIndex-1], "")
		for i := sliceStart; i <= sliceEnd; i++ {
			boxValues = append(boxValues, line[i])
		}
	}
	// Get current line values
	if startIndex != 0 {
		boxValues = append(boxValues, currentLine[sliceStart])
	}

	if endIndex != len(currentLine)-1 {
		boxValues = append(boxValues, currentLine[sliceEnd])
	}

	// Get below line values
	if lineIndex != len(lines)-1 {
		line := strings.Split(lines[lineIndex+1], "")
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
