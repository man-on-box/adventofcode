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
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func main() {
	f, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(f))
	lines := strings.Split(input, "\n")
	grid := make([][]rune, len(lines))
	seen := map[point]bool{}
	placedObstacles := map[point]bool{}

	startPoint := point{-1, -1}

	for i, l := range lines {
		row := []rune(l)
		grid[i] = row
		if startPoint.x > 0 {
			continue
		}
		for j, r := range row {
			if r == '^' {
				startPoint.x = j
				startPoint.y = i
			}
		}
	}

	current := startPoint
	direction := 0

	for {
		seen[current] = true
		dOuter := directions[direction]
		next := point{
			x: current.x + dOuter.x,
			y: current.y + dOuter.y,
		}

		// check if we made it off the map
		if next.x < 0 ||
			next.x > len(grid[0])-1 ||
			next.y < 0 ||
			next.y > len(grid)-1 {
			break
		}

		// are we hitting a wall? Then we just change direction
		if grid[next.y][next.x] == '#' {
			direction = (direction + 1) % 4
			continue
		}

		// otherwise move to next point
		current = next

	}

	pt1DistinctSteps := len(seen)
	// Part 1 = 5531
	fmt.Println("Part 1:", pt1DistinctSteps)

	// Now we go through each place the guard visited,
	// and we place an obstacle, to see if we get a loop
	for p := range seen {
		// if we're not a point, we pass over it
		if grid[p.y][p.x] != '.' {
			continue
		}

		guardCurrent := startPoint
		guardDirection := 0
		guardSeen := map[[3]int]bool{}

		// we set the current to a wall
		grid[p.y][p.x] = '#'

		for {
			state := [3]int{guardCurrent.x, guardCurrent.y, guardDirection}
			if guardSeen[state] {
				placedObstacles[p] = true
				break
			}

			guardSeen[state] = true
			d := directions[guardDirection%4]
			guardNext := point{
				x: guardCurrent.x + d.x,
				y: guardCurrent.y + d.y,
			}

			// check if we made it off the map
			if guardNext.x < 0 ||
				guardNext.x > len(grid[0])-1 ||
				guardNext.y < 0 ||
				guardNext.y > len(grid)-1 {
				break
			}

			// are we hitting a wall? Then we just change direction
			if grid[guardNext.y][guardNext.x] == '#' {
				guardDirection = (guardDirection + 1) % 4
				continue
			}

			guardCurrent = guardNext

		}

		// reset the value back to a point
		grid[p.y][p.x] = '.'

	}

	// part 2 = 2165
	fmt.Println("Part 2:", len(placedObstacles))

}
