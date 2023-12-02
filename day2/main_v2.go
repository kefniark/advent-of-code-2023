package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/**
 * Day 2: Cube Conundrum
 * url: https://adventofcode.com/2023/day/2
 *
 * Version 2: Try to improve a bit my code
 */

func main() {
	sum := 0
	sum2 := 0

	lines := findMaxColors("input.txt")
	for idx, line := range lines {
		// Part 1
		max := []int{12, 14, 13} // Red, Blue, Green
		if line[0] <= max[0] && line[1] <= max[1] && line[2] <= max[2] {
			sum += idx + 1
		}

		// Part 2
		val := line[0] * line[1] * line[2]
		sum2 += val
	}

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)
}

func findMaxColors(filename string) [][]int {
	colors := map[string]int{"red": 0, "blue": 1, "green": 2}
	bytesRead, _ := os.ReadFile(filename)
	re := regexp.MustCompile(`(\d+) (\w+)`)

	colorMaximums := [][]int{}
	for _, s := range strings.Split(strings.TrimSpace(string(bytesRead)), "\n") {
		lineMax := []int{0, 0, 0}
		for _, val := range re.FindAllStringSubmatch(s, -1) {
			num, _ := strconv.Atoi(val[1])
			color := colors[val[2]]

			lineMax[color] = max(lineMax[color], num)
		}

		colorMaximums = append(colorMaximums, lineMax)
	}

	return colorMaximums
}

// Hum I may need to update my local go version, I thought it was already in the stdlib

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
