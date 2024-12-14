package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(f))
	disk := []int{}
	id := 0

	for i, v := range input {
		val, _ := strconv.Atoi(string(v))
		isFile := i%2 == 0
		if isFile {
			for i := 0; i < val; i++ {
				disk = append(disk, (id))
			}
			id++
		} else {
			for i := 0; i < val; i++ {
				disk = append(disk, -1)
			}
		}
	}

	diskP1 := slices.Clone(disk)
	left := 0
	right := len(diskP1) - 1
	for left < right {
		if diskP1[left] >= 0 {
			left++
			continue
		}
		if diskP1[right] < 0 {
			right--
			continue
		}
		tmp := diskP1[left]
		diskP1[left] = diskP1[right]
		diskP1[right] = tmp

	}

	pt1 := 0
	for i, v := range diskP1 {
		if v == -1 {
			continue
		}
		pt1 += i * v
	}
	fmt.Println("Part 1:", pt1)

	// Part 2
	diskP2 := slices.Clone(disk)
	currentFile := -1
	currentFileSize := 0
	movedCache := map[int]bool{}

	for i := len(diskP2) - 1; i >= 0; i-- {
		val := diskP2[i]
		if currentFile == -1 {
			if val == -1 {
				continue
			} else {
				if movedCache[val] {
					continue
				} else {
					movedCache[val] = true
					currentFile = val
					currentFileSize = 1
				}
			}
		} else {
			if val == currentFile {
				currentFileSize++
				continue
			} else {
				// we reached the end of current file
				// we need to see if there is space to move it
				freeblockSize := 0
				for j := 0; j <= i; j++ {
					leftVal := diskP2[j]
					if leftVal == -1 {
						freeblockSize++
						if freeblockSize >= currentFileSize {
							// then we have space.... lets swap
							fileStart := i + 1
							fileEnd := fileStart + currentFileSize
							blockStart := j + 1 - freeblockSize
							blockEnd := blockStart + freeblockSize
							for k := blockStart; k < blockEnd; k++ {
								diskP2[k] = currentFile
							}
							for k := fileStart; k < fileEnd; k++ {
								diskP2[k] = -1
							}
							freeblockSize = 0
							break
						}
					} else {
						freeblockSize = 0
					}
				}

				// then we need to check current val again
				// if it's an empty space, we clear current file and space
				// or if it's another file, we set it to that and space = 1
				if val == -1 {
					currentFileSize = 0
					currentFile = val
				} else {
					if movedCache[val] {
						currentFile = -1
						currentFileSize = 0
					} else {
						movedCache[val] = true
						currentFile = val
						currentFileSize = 1
					}
				}
			}
		}
	}

	pt2 := 0
	for i, v := range diskP2 {
		if v == -1 {
			continue
		}
		pt2 += i * v
	}

	// Part 2 = 6307653242596
	fmt.Println("Part 2:", pt2)
}
