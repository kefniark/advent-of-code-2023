package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

/**
 * Day 4: Scratchcards - Part 1
 * url: https://adventofcode.com/2023/day/4
 */
func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	cards := []int{}
	for _, s := range lines {
		win, got := parseLine(s)
		points := countIntersection(win, got)
		cards = append(cards, points)
	}

	sum := 0
	for _, val := range cards {
		if val > 0 {
			sum += int(math.Pow(2, float64(val-1)))
		}
	}

	fmt.Println("Part 1 = ", sum)
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
