package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(content))
	sections := strings.Split(input, "\n\n")

	orderingSection := strings.Split(sections[0], "\n")
	// Orderings is a map of AFTERs, and each contain a map
	// of page numbers that should be BEFORE.
	// This way we can iterate a report, and if any numbers that appear
	// after the current value (by referring to the map of BEFORE values)
	// we know it is not valid
	orderings := map[string]map[string]struct{}{}
	for _, item := range orderingSection {
		parts := strings.Split(item, "|")
		if orderings[parts[1]] == nil {
			orderings[parts[1]] = map[string]struct{}{}
		}
		orderings[parts[1]][parts[0]] = struct{}{}
	}

	updatesSection := strings.Split(sections[1], "\n")
	updates := make([][]string, len(updatesSection))
	for i, line := range updatesSection {
		updates[i] = strings.Split(line, ",")
	}

	pt1Total := 0
	pt2Total := 0

	for _, u := range updates {
		isNotValid := false
		for i := 0; i < len(u); i++ {
			for j := i + 1; j < len(u); j++ {
				current := u[i]
				compared := u[j]
				if orderings[current] != nil {
					_, exists := orderings[current][compared]
					if exists {
						isNotValid = true
						// part 2
						// now swap values
						u[i] = compared
						u[j] = current
					}
				}
			}
		}

		middleIndex := (len(u) - 1) / 2
		middleValue, _ := strconv.Atoi(u[middleIndex])

		if isNotValid {
			pt2Total += middleValue
			continue
		}

		pt1Total += middleValue
	}

	// part 1 answer: 6242
	fmt.Println("Part 1 total:", pt1Total)
	// part 2 answer: 5169
	fmt.Println("Part 2 total:", pt2Total)

}
