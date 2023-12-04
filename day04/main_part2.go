package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

/**
 * Day 4: Scratchcards - Part 2
 * url: https://adventofcode.com/2023/day/4
 */
func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// create a slice of instance counter
	cardCount := []int{}
	for i := 0; i < len(lines); i++ {
		cardCount = append(cardCount, 1)
	}

	// for each line, we count the number of win/got intersection
	for i, s := range lines {
		win, got := parseLine(s)
		points := countIntersection(win, got)

		// we increment the number of copies for each card
		// k: is the number of copies of the current card
		// j: is the number of winning point
		for k := 0; k < cardCount[i]; k++ {
			for j := 1; j <= points; j++ {
				cardCount[i+j] += 1
			}
		}
	}

	// count cards
	sum := 0
	for i := range cardCount {
		sum += cardCount[i]
	}

	fmt.Println("Part 2 = ", sum)
}

func parseLine(s string) ([]int, []int) {
	startAt := strings.Index(strings.TrimSpace(s), ":")
	sets := strings.Split(s[startAt+1:], "|")

	win := parseNumbers(sets[0])
	got := parseNumbers(sets[1])

	return win, got
}

func parseNumbers(s string) []int {
	re := regexp.MustCompile(`\d+`)
	numbers := []int{}
	for _, val := range re.FindAllStringSubmatch(s, -1) {
		num, _ := strconv.Atoi(val[0])
		numbers = append(numbers, num)
	}

	return numbers
}

func countIntersection(set1 []int, set2 []int) int {
	count := 0

	for _, val := range set1 {
		if slices.Contains(set2, val) {
			count++
		}
	}

	return count
}
