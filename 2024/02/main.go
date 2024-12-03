package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(content))
	lines := strings.Split(input, "\n")
	safeReports := 0

	for _, l := range lines {
		report := strings.Fields(l)
		isSafe, problemIndex := isReportSafe(report)

		if !isSafe {
			for i := 0; i < 2; i++ {
				// run again but without the problem index
				// and again with the problem index - 1
				sliceIndex := problemIndex - i
				modifiedReport := []string{}
				modifiedReport = append(modifiedReport, report[:sliceIndex]...)
				modifiedReport = append(modifiedReport, report[sliceIndex+1:]...)
				isSafe, _ = isReportSafe(modifiedReport)
				if isSafe {
					break
				}
			}
		}
		if isSafe {
			safeReports++
		}
	}

	// Part 1 answer: 639
	// Part 2 answer: 674
	fmt.Println("Total safe reports:", safeReports)

}

func isReportSafe(report []string) (result bool, problemIndex int) {
	isSafe, asc, desc := true, true, true
	i := 1

	for isSafe && i < len(report) {
		value, _ := strconv.Atoi(report[i])
		prevVal, _ := strconv.Atoi(report[i-1])
		diff := value - prevVal
		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			isSafe = false
		}

		if value < prevVal {
			asc = false
		}
		if value > prevVal {
			desc = false
		}
		if asc == false && desc == false {
			isSafe = false
		}
		if isSafe {
			i++
		}

	}

	return isSafe, i

}
