package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/**
 * Day 2: Cube Conundrum
 * url: https://adventofcode.com/2023/day/2
 *
 * Version 2: Try to improve a bit my code
 * - use embed to read the input file
 * - use regexp to extract values
 * - use map resolve colors <> int
 */

//go:embed input.txt
var input string

func main() {
	sum := 0
	sum2 := 0

	lines := findMaxColors()
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

func findMaxColors() [][]int {
	re := regexp.MustCompile(`(\d+) (\w+)`)
	colors := map[string]int{"red": 0, "blue": 1, "green": 2}

	colorMaximums := [][]int{}
	for _, s := range strings.Split(strings.TrimSpace(input), "\n") {
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
