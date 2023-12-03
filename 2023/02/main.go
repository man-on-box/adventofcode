package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	r int
	g int
	b int
}

type Game struct {
	sets       []Set
	maxCubes   Set
	isPossible bool
}

var MaxCubes = Set{
	r: 12,
	g: 13,
	b: 14,
}

var maxRedCubes = 0
var maxGreenCubes = 0
var mainBlueCubes = 0

func main() {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	var games = make(map[int]Game)
	for _, line := range lines {
		id, game := parseGame(line)
		games[id] = game
	}

	var possibleGames []int
	var powerOfSets []int
	for id, game := range games {
		if game.isPossible {
			possibleGames = append(possibleGames, id)
		}
		power := game.maxCubes.r * game.maxCubes.g * game.maxCubes.b
		powerOfSets = append(powerOfSets, power)
	}

	fmt.Println("Part 1: ", sum(possibleGames))
	fmt.Println("Part 2: ", sum(powerOfSets))

}

func parseGame(line string) (int, Game) {
	parts := strings.Split(line, ": ")
	gameId, _ := strconv.Atoi(strings.Replace(parts[0], "Game ", "", -1))
	sets, maxCubes, isPossible := parseSets(parts[1])
	game := Game{
		sets:       sets,
		maxCubes:   maxCubes,
		isPossible: isPossible,
	}
	return gameId, game
}

func parseSets(line string) ([]Set, Set, bool) {
	parts := strings.Split(line, "; ")
	var sets []Set
	var maxCubes Set

	for _, setString := range parts {
		set := parseSet(setString)
		sets = append(sets, set)
	}

	for _, set := range sets {
		if set.r > maxCubes.r {
			maxCubes.r = set.r
		}
		if set.g > maxCubes.g {
			maxCubes.g = set.g
		}
		if set.b > maxCubes.b {
			maxCubes.b = set.b
		}
	}

	isPossible := maxCubes.r <= MaxCubes.r && maxCubes.g <= MaxCubes.g && maxCubes.b <= MaxCubes.b

	return sets, maxCubes, isPossible
}

func parseSet(setLine string) Set {
	cubes := strings.Split(setLine, ", ")
	set := Set{}
	for _, cube := range cubes {
		parts := strings.Split(cube, " ")
		count, _ := strconv.Atoi(parts[0])
		color := parts[1]
		switch color {
		case "red":
			set.r = count
		case "green":
			set.g = count
		case "blue":
			set.b = count
		}
	}
	return set
}

func sum(a []int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}
