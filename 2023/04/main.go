package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type card struct {
	winValues map[int]bool
	value     []int
}

type CardSet struct {
	cards []card
}

func main() {
	input, _ := os.ReadFile("example.txt")
	lines := strings.Split(string(input), "\n")
	cardRegx := regexp.MustCompile(`Card \d+:`)
	spaceRegx := regexp.MustCompile(`\s+`)
	cardSet := newCardSet()
	for _, line := range lines {
		line = cardRegx.ReplaceAllString(line, "")
		line = spaceRegx.ReplaceAllString(strings.TrimSpace(line), ",")
		split := strings.Split(line, ",|,")
		card := newCard()

		for _, d := range strings.Split(split[0], ",") {
			v, _ := strconv.Atoi(d)
			card.winValues[v] = true
		}
		for _, d := range strings.Split(split[1], ",") {
			v, _ := strconv.Atoi(d)
			card.value = append(card.value, v)

		}

		cardSet.cards = append(cardSet.cards, card)
	}

	// Part 1 = 23750
	total := 0
	for _, game := range cardSet.cards {
		total += game.score()
	}
	fmt.Println("Total:", total)

	// Part 2
	cardCount := 0
	toProcess := []CardSet{cardSet}
	for len(toProcess) > 0 {
		set := toProcess[0]
		toProcess = toProcess[1:]
		winningSets := set.getWinningCopies()
		fmt.Println("No of winning sets:", len(winningSets))
		if len(winningSets) > 0 {
			for _, s := range winningSets {
				wonCards := len(s.cards)
				fmt.Println("Won cards:", wonCards)
				cardCount += len(s.cards)
			}
			toProcess = append(toProcess, winningSets...)
		}

	}
	fmt.Println("No of cards:", cardCount)

}

func newCardSet() CardSet {
	return CardSet{
		cards: []card{},
	}
}

func (cs CardSet) getWinningCopies() []CardSet {
	winningSets := []CardSet{}
	for i, c := range cs.cards {
		score := len(c.winnings())
		if score > 0 {
			winningSet := newCardSet()
			start := i + 1
			end := start + len(c.winnings())
			fmt.Println("SCORE", score, "Start:", start, "End:", end, "Len:", len(cs.cards))
			for j := start; j < end; j++ {
				if j >= len(cs.cards) {
					break
				}
				winningSet.cards = append(winningSet.cards, cs.cards[j])
			}
			winningSets = append(winningSets, winningSet)
		} else {
			continue
		}

	}
	return winningSets
}

func newCard() card {
	return card{
		winValues: map[int]bool{},
		value:     []int{},
	}
}

func (c card) winnings() []int {
	numbers := []int{}
	for _, n := range c.value {
		if _, ok := c.winValues[n]; ok {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

func (c card) score() int {
	winningNumbers := c.winnings()
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

func (c card) copy() card {
	return c
}

func sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}
