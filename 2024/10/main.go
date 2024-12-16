package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var directions = [][]int{
	{0, 1},
	{-1, 0},
	{0, -1},
	{1, 0},
}

func walk(grid *[][]int, curr [2]int, targetStep int, heads *map[[2]int]bool, score int) int {
	g := *grid
	h := *heads
	s := score
	x, y := curr[0], curr[1]
	// base cases
	if x < 0 || x > len(g[0])-1 || y < 0 || y > len(g)-1 {
		return 0
	}
	if g[y][x] != targetStep {
		return 0
	}
	if targetStep == 9 && g[y][x] == 9 {
		h[curr] = true
		return 1
	}
	for _, d := range directions {
		next := [2]int{x + d[0], y + d[1]}
		s += walk(grid, next, targetStep+1, heads, score)
	}
	return s
}

func main() {
	f, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(f))
	lines := strings.Split(input, "\n")
	grid := make([][]int, len(lines))

	for i, l := range lines {
		row := make([]int, len(l))
		for j, stringVal := range l {
			intVal, _ := strconv.Atoi(string(stringVal))
			row[j] = intVal
		}
		grid[i] = row
	}

	pt1 := 0
	pt2 := 0
	for y := range grid {
		for x, val := range grid[y] {
			if val == 0 {
				headsMap := map[[2]int]bool{}
				pt2 += walk(&grid, [2]int{x, y}, 0, &headsMap, 0)
				pt1 += len(headsMap)
			}
		}
	}

	// part 1 = 550
	fmt.Println("Part 1:", pt1)
	// part 2 = 1255
	fmt.Println("Part 2:", pt2)

}
