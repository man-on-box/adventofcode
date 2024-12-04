package main

import (
	"fmt"
	"os"
	"strings"
)

type point = struct {
	x int
	y int
}

var directions = []point{
	{0, -1},  // up
	{1, -1},  // up right
	{-1, -1}, // up left
	{0, 1},   // down
	{1, 1},   // down right
	{-1, 1},  // down left
	{-1, 0},  // left
	{1, 0},   // right
}

// part 2
var diagonals = []point{
	{1, -1},  // up right
	{-1, -1}, // up left
	{1, 1},   // down right
	{-1, 1},  // down left
}

func findWord(word string, wordsearch *[][]rune, p point, d point, wordIndex int) bool {
	ws := *wordsearch
	// BASE CASES
	// are we off the table?
	if p.x < 0 ||
		p.x > len(ws[0])-1 ||
		p.y < 0 ||
		p.y > len(ws)-1 {
		return false
	}
	// is the char not what we are looking for?
	if ws[p.y][p.x] != rune(word[wordIndex]) {
		return false
	}
	// have we found the word?
	if wordIndex == len(word)-1 {
		return true
	}
	// otherwise we have found the char so far, and we continue again in the same direction
	nextPoint := point{
		x: p.x + d.x,
		y: p.y + d.y,
	}
	return findWord(word, wordsearch, nextPoint, d, wordIndex+1)
}

func main() {
	content, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(content))
	lines := strings.Split(input, "\n")
	wordsearch := make([][]rune, len(lines))

	pt1Total := 0
	pt2Total := 0

	for i, line := range lines {
		row := []rune(line)
		wordsearch[i] = row
	}

	for colIndex, col := range wordsearch {
		for rowIndex, char := range col {
			p := point{x: rowIndex, y: colIndex}
			for _, d := range directions {
				found := findWord("XMAS", &wordsearch, p, d, 0)
				if found {
					pt1Total++
				}
			}

			// part 2
			if char == 'A' {
				masCount := 0
				for _, d := range diagonals {
					startPoint := point{
						x: p.x + d.x,
						y: p.y + d.y,
					}

					found := findWord("MAS", &wordsearch, startPoint, point{x: -d.x, y: -d.y}, 0)
					if found {
						masCount++
					}
				}
				if masCount == 2 {
					pt2Total++
				}
			}

		}
	}

	fmt.Println("Part 1 total:", pt1Total)
	fmt.Println("Part 2 total:", pt2Total)
}
