package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

/**
 * Day 6: Wait For It - Part 2
 * url: https://adventofcode.com/2023/day/6
 */
func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	race, _ := parseNumbers(lines[0])
	best, _ := parseNumbers(lines[1])

	fmt.Println("Part 2 = ", findWinning(race, best))
}

func findWinning(ms int, goal int) int {
	count := 0
	for i := 1; i < ms; i++ {
		val := (ms - i) * i
		if val > goal {
			count += 1
		}
	}

	return count
}

func parseNumbers(s string) (int, error) {
	re := regexp.MustCompile(`\d+`)
	numbers := ""
	for _, val := range re.FindAllStringSubmatch(s, -1) {
		numbers += val[0]
	}

	return strconv.Atoi(numbers)
}
