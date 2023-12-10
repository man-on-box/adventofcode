package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	winValues map[int]bool
	value     []int
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	cardRegx := regexp.MustCompile(`Card \d+:`)
	spaceRegx := regexp.MustCompile(`\s+`)
	games := []Game{}
	for _, line := range lines {
		line = cardRegx.ReplaceAllString(line, "")
		line = spaceRegx.ReplaceAllString(strings.TrimSpace(line), ",")
		split := strings.Split(line, ",|,")
		game := newGame()

		for _, d := range strings.Split(split[0], ",") {
			v, _ := strconv.Atoi(d)
			game.winValues[v] = true
		}
		for _, d := range strings.Split(split[1], ",") {
			v, _ := strconv.Atoi(d)
			game.value = append(game.value, v)

		}

		games = append(games, game)
	}

	// Part 1 = 23750
	total := 0
	for _, game := range games {
		total += game.score()
	}
	fmt.Println("Total:", total)
}

func newGame() Game {
	return Game{
		winValues: map[int]bool{},
		value:     []int{},
	}
}

func (g Game) winnings() []int {
	numbers := []int{}
	for _, n := range g.value {
		if _, ok := g.winValues[n]; ok {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

func (g Game) score() int {
	winningNumbers := g.winnings()
	score := 0
	for i := range winningNumbers {
		if i > 0 {
			score = score * 2
		} else {
			score = 1
		}
	}
	return score
}

func sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}
