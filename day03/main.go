package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

/**
 * Day 3: Gear Ratios - Part 1
 * url: https://adventofcode.com/2023/day/3
 *
 * Part 1: Find all numbers in the grid adjacent to a symbol
 * Part 2: I preferred to do it in a separated file, too different implementation
 */

//go:embed input.txt
var input string
var digits = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var symbols = [...]string{"#", "$", "+", "*", "%", "@", "&", "-", "/", "="}

func main() {
	// create a mask to find where are valid numbers
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// create binary mask
	mask := make([][]bool, len(lines))
	for i, s := range lines {
		s = strings.TrimSpace(s)
		mask[i] = make([]bool, len(s))
	}

	// set the mask, every neighbor of a symbol is valid
	for j, s := range lines {
		s = strings.TrimSpace(s)
		for i := 0; i < len(s); i++ {
			if !slices.Contains(symbols[:], string(s[i])) {
				continue
			}

			// for neighbors of a symbol, set the mask
			for x := -1; x < 2; x++ {
				for y := -1; y < 2; y++ {
					x1 := i + x
					y1 := j + y

					if mask[y1][x1] {
						continue
					}

					if x1 >= 0 && x1 < len(s) && y1 >= 0 && y1 < len(lines) {
						mask[y1][x1] = true
					}
				}
			}
		}
	}

	// find all numbers
	numbers := []int{}
	for j, s := range lines {
		s = strings.TrimSpace(s)
		buffer := ""
		bufferValid := false
		for i := 0; i < len(s); i++ {
			char := string(s[i])
			if slices.Contains(symbols[:], char) || char == "." {
				if bufferValid {
					num, _ := strconv.Atoi(buffer)
					numbers = append(numbers, num)
				}
				buffer = ""
				bufferValid = false
				continue
			}

			buffer += char
			if mask[j][i] {
				bufferValid = true
			}
		}

		if bufferValid {
			num, _ := strconv.Atoi(buffer)
			numbers = append(numbers, num)
		}
	}

	// sum all numbers
	sum := 0
	for _, num := range numbers {
		sum = sum + num
	}

	fmt.Println("Sum:", sum)
}
