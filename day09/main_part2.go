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
 * Day 9: Mirage Maintenance - Part 2
 * url: https://adventofcode.com/2023/day/9
 */
func main() {
	sum := 0
	for _, s := range strings.Split(strings.TrimSpace(input), "\n") {
		line := parseNumbers(strings.TrimSpace(s))

		tower := generateTower(line)
		next := extrapolateNextValue(tower)

		sum += next
	}

	fmt.Println("Part 2 =", sum)
}

func extrapolateNextValue(tower [][]int) int {
	lastLine := len(tower) - 1
	tower[lastLine] = append([]int{0}, tower[lastLine]...)
	for i := len(tower) - 2; i >= 0; i-- {
		newVal := tower[i][0] - tower[i+1][0]
		tower[i] = append([]int{newVal}, tower[i]...)
	}

	return tower[0][0]
}

func generateTower(num []int) [][]int {
	tower := [][]int{}
	tower = append(tower, num)
	for {
		row := tower[len(tower)-1]
		nextRow := []int{}
		finish := true

		for i := 0; i < len(row)-1; i++ {
			diff := row[i+1] - row[i]
			if diff != 0 {
				finish = false
			}
			nextRow = append(nextRow, diff)
		}

		tower = append(tower, nextRow)
		if finish {
			break
		}
	}

	return tower
}

func parseNumbers(s string) []int {
	re := regexp.MustCompile(`-?\d+`)
	numbers := []int{}
	for _, val := range re.FindAllStringSubmatch(s, -1) {
		num, _ := strconv.Atoi(val[0])
		numbers = append(numbers, num)
	}

	return numbers
}
