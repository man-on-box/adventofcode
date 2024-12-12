package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(f))
	lines := strings.Split(input, "\n")
	ants := map[rune][][2]int{}
	antinodes := map[[2]int]bool{}
	repeatingAntinodes := map[[2]int]bool{}

	for row, l := range lines {
		for col, r := range l {
			if r != '.' {
				ants[r] = append(ants[r], [2]int{col, row})
			}
		}
	}

	for _, points := range ants {
		for i := range points {
			for j := i + 1; j < len(points); j++ {
				pos1 := points[i]
				pos2 := points[j]
				// here we compare each matching antenna and get the
				// deltas of their coords, to calculate the antinodes
				delta := [2]int{
					pos2[0] - pos1[0],
					pos2[1] - pos1[1],
				}
				// now get the antinode coords, for the first and second
				// antenna, by using their positions +- deltas
				nodes := [][2]int{
					{pos1[0] - delta[0], pos1[1] - delta[1]},
					{pos2[0] + delta[0], pos2[1] + delta[1]},
				}
				// Now we loop through the nodes, and remove any that are
				// off of the grid
				for _, node := range nodes {
					if node[0] >= 0 && node[0] < len(lines[0]) && node[1] >= 0 && node[1] < len(lines) {
						antinodes[node] = true
					}
				}
				// part 2
				// we know we are already in a pair, so push current antenna to
				// repeating antennas
				repeatingAntinodes[pos1] = true
				repeatingAntinodes[pos2] = true
				// pos1 repeating
				for {
					antinode := [2]int{
						pos1[0] - delta[0], pos1[1] - delta[1],
					}

					if antinode[0] >= 0 && antinode[0] < len(lines[0]) && antinode[1] >= 0 && antinode[1] < len(lines) {
						repeatingAntinodes[antinode] = true
						pos1 = antinode
					} else {
						break
					}
				}
				// pos2 repeating
				for {
					antinode := [2]int{
						pos2[0] + delta[0], pos2[1] + delta[1],
					}

					if antinode[0] >= 0 && antinode[0] < len(lines[0]) && antinode[1] >= 0 && antinode[1] < len(lines) {
						repeatingAntinodes[antinode] = true
						pos2 = antinode
					} else {
						break
					}
				}
			}
		}
	}

	// Part 1 = 311
	fmt.Println("Part 1:", len(antinodes))
	fmt.Println("Part 2:", len(repeatingAntinodes))
}
