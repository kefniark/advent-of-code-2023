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
 * Day 6: Wait For It - Part 1
 * url: https://adventofcode.com/2023/day/6
 */
func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	race := parseNumbers(lines[0])
	best := parseNumbers(lines[1])

	count := 1
	for i := 0; i < len(race); i++ {
		res := findWinning(race[i], best[i])
		count *= len(res)
	}

	fmt.Println("Part 1 = ", count)
}

func findWinning(ms int, goal int) []int {
	res := []int{}
	for i := 1; i < ms; i++ {
		val := (ms - i) * i
		if val > goal {
			res = append(res, i)
		}
	}

	return res
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
