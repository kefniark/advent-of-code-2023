package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

/**
 * Day 3: Gear Ratios - Part 2
 * url: https://adventofcode.com/2023/day/3
 *
 * The mask approach used in part 1 is not working for part 2, so I decided to change the approach.
 * Find all the gear then from there flood fill to find the whole numbers around it.
 */

//go:embed input.txt
var input string
var digits = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// create grid of characters
	grid := make([][]string, len(lines))
	for j, s := range lines {
		s = strings.TrimSpace(s)
		grid[j] = make([]string, len(s))
		for i := 0; i < len(s); i++ {
			grid[j][i] = string(s[i])
		}
	}

	// find gear and calculate their gear ratio
	sumGearRatio := 0
	for j, line := range grid {
		for i, char := range line {
			if char != "*" {
				continue
			}

			// create a map of all adjacent numbers to a gear
			adjacentNumber := map[int]int{}
			for x := -1; x < 2; x++ {
				for y := -1; y < 2; y++ {
					x1 := i + x
					y1 := j + y

					if x1 >= 0 && x1 < len(line) && y1 >= 0 && y1 < len(grid) && slices.Contains(digits[:], grid[y1][x1]) {
						idx, num := floodFillNumber(grid, x1, y1)
						adjacentNumber[idx] = num
					}
				}
			}

			// sum the gear ratio
			if len(adjacentNumber) >= 2 {
				gearRatio := 1
				for _, num := range adjacentNumber {
					gearRatio = gearRatio * num
				}
				sumGearRatio = sumGearRatio + gearRatio
			}
		}
	}

	fmt.Println("Sum gear ratio:", sumGearRatio)
}

// from a starting location on the grid, flood fill in both direction (left & right) to find the whole number
func floodFillNumber(grid [][]string, x int, y int) (int, int) {
	buffer := grid[y][x]
	minX := x

	// on the left (-1), prepend to the buffer
	// on the right (+1), append to the buffer
	for _, dir := range []int{-1, 1} {
		for j := 1; j < 10; j++ {
			i := j * dir

			if x+i < 0 || x+i >= len(grid[y]) {
				break
			}

			char := string(grid[y][x+i])
			if !slices.Contains(digits[:], char) {
				break
			}

			if dir < 0 {
				buffer = fmt.Sprintf("%s%s", char, buffer)
				minX = x + i
			} else {
				buffer = fmt.Sprintf("%s%s", buffer, char)
			}
		}
	}

	// convert to int & create a unique index to deduplicate
	num, _ := strconv.Atoi(buffer)
	return (minX + 1) * (y + 1), num
}
