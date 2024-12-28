package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(f))
	fields := strings.Fields(input)
	stones := map[int]int{}
	stonesInProcess := map[int]int{}

	for _, val := range fields {
		stone, _ := strconv.Atoi(val)
		stones[stone]++
	}

	start := time.Now()

	checkStone := func(stone int, count int) {
		if stone == 0 {
			stonesInProcess[1] += count
			return
		}

		asString := strconv.Itoa(stone)
		if len(asString)%2 == 0 {
			asSlice := strings.Split(asString, "")
			middle := len(asSlice) / 2
			first, _ := strconv.Atoi(strings.Join(asSlice[:middle], ""))
			second, _ := strconv.Atoi(strings.Join(asSlice[middle:], ""))
			stonesInProcess[first] += count
			stonesInProcess[second] += count
			return
		}

		stonesInProcess[stone*2024] += count
		return
	}

	blink := func(times int) {
		for i := 1; i <= times; i++ {
			stonesInProcess = map[int]int{}
			for stone, count := range stones {
				checkStone(stone, count)
			}
			stones = stonesInProcess
		}
	}

	blink(75)
	// Part 1, 25 blinks = 175006
	// Part 2, 75 blinks = 207961583799296
	output := 0
	for _, v := range stones {
		output += v
	}
	fmt.Println("Part #:", output)
	done := time.Since(start)
	fmt.Println("done in", done)
}
