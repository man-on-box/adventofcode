package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	var calibrationValues []int
	for _, line := range lines {
		parsedTexttoInt := textToIntInString(line)
		intsInLine := parseIntsFromString(parsedTexttoInt)
		firstAndLastInt := combineFirstAndLast(intsInLine)
		calibrationValues = append(calibrationValues, firstAndLastInt)
	}

	fmt.Println(sum(calibrationValues))
}

var textToIntMap = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "th3ee",
	"four":  "fo4r",
	"five":  "fi5e",
	"six":   "s6x",
	"seven": "se7en",
	"eight": "ei8ht",
	"nine":  "ni9e",
}

func textToIntInString(input string) string {
	parsed := input
	for key := range textToIntMap {
		parsed = strings.ReplaceAll(parsed, key, textToIntMap[key])
	}
	return parsed
}

func parseIntsFromString(input string) []int {
	var numbers []int
	for _, char := range input {
		number, err := strconv.Atoi(string(char))
		if err == nil {
			numbers = append(numbers, number)
		}
	}
	return numbers
}

func combineFirstAndLast(numbers []int) int {
	first := strconv.Itoa(numbers[0])
	last := strconv.Itoa(numbers[len(numbers)-1])
	combined, _ := strconv.Atoi(first + last)
	return combined
}

func sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
