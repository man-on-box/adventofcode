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

var pt2Total = 0

func move(grid *[][]rune, seen *map[point]struct{}, seenLoop *map[point]struct{}, curr point, turns int) bool {
	g := *grid
	s := *seen
	// BASE CASES
	// are we off the map? Then we're done
	if curr.x < 0 || curr.x > len(g[0])-1 || curr.y < 0 || curr.y > len(g)-1 {
		return true
	}
	// is our point an obstacle? If YES then we can't go further
	if g[curr.y][curr.x] == '#' {
		return false
	}

	// Otherwise we keep going
	d := directions[turns%len(directions)]

	nextPoint := point{
		x: curr.x + d.x,
		y: curr.y + d.y,
	}
	result := move(grid, seen, seenLoop, nextPoint, turns)

	if result == false {
		// then we hit an obstacle, lets turn right and go again
		return move(grid, seen, seenLoop, curr, turns+1)
	}

	if seenLoop != nil {
		// then we are checking if we get in a loop
		sl := *seenLoop
		_, looped := sl[curr]
		if looped {
			// fmt.Println("Loop at...", curr)
			pt2Total++
			return true
		}
	}

	// pt 2, simulate a loop if we have been here before
	_, beenHere := s[curr]
	if seenLoop == nil && beenHere {
		sl := map[point]struct{}{}
		sl[curr] = struct{}{}
		// fmt.Println("Checking loop from", curr)
		move(grid, seen, &sl, curr, turns+1)
	}

	// post recursion
	// we walked this point so append to our seen map
	if seenLoop != nil {
		sl := *seenLoop
		sl[curr] = struct{}{}
	} else {
		s[curr] = struct{}{}
	}

	return true

}

func main() {
	f, _ := os.ReadFile("example.txt")
	input := strings.TrimSpace(string(f))
	lines := strings.Split(input, "\n")
	grid := make([][]rune, len(lines))
	seen := map[point]struct{}{}
	var startPoint *point

	for i, l := range lines {
		row := []rune(l)
		grid[i] = row
		if startPoint != nil {
			continue
		}
		for j, r := range row {
			if r == '^' {
				startPoint = &point{
					x: j,
					y: i,
				}
			}
		}
	}

	move(&grid, &seen, nil, *startPoint, 0)

	pt1DistinctSteps := len(seen)

	fmt.Println("Part 1:", pt1DistinctSteps)
	fmt.Println("Part 2:", pt2Total)
}
