package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * Day 2: Cube Conundrum
 * url: https://adventofcode.com/2023/day/2
 */

func main() {
	sum := 0
	sum2 := 0

	lines := findMaxColors("input.txt")

	// For Testing
	// lines := findMaxColors("input_test.txt")

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

/**
 * Find the maximum number of cubes of each color for each line of a file
 */
func findMaxColors(filename string) [][]int {
	bytesRead, _ := os.ReadFile(filename)
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")

	colorMaximums := [][]int{}

	for _, line := range lines {
		line = strings.TrimSpace(line)

		args := strings.Split(line, ":")
		if len(args) != 2 {
			fmt.Println("Invalid input", args)
			return nil
		}

		tries := strings.Split(args[1], ";")
		lineMax := []int{0, 0, 0}
		for _, try := range tries {
			cubes := strings.Split(try, ",")
			for _, cube := range cubes {
				val := strings.Split(strings.TrimSpace(cube), " ")
				num, _ := strconv.Atoi(val[0])
				color := parseColor(val[1])
				if color == -1 {
					fmt.Println("Invalid input", args)
					return nil
				}

				lineMax[color] = max(lineMax[color], num)
			}
		}
		colorMaximums = append(colorMaximums, lineMax)
	}

	return colorMaximums
}

func parseColor(color string) int {
	switch color {
	case "red":
		return 0
	case "blue":
		return 1
	case "green":
		return 2
	default:
		return -1
	}
}

// Hum I may need to update my local go version, I thought it was already in the stdlib

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
