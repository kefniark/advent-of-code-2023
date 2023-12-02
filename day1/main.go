package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/**
 * Day 1: Calibration
 * url: https://adventofcode.com/2023/day/1
 */

func main() {
	calibration("input.txt")

	// For Testing
	// calibration("input_test.txt")
}

func calibration(filename string) {
	bytesRead, _ := os.ReadFile(filename)
	fileContent := string(bytesRead)

	replacements := [][]string{
		{"one", "1"},
		{"two", "2"},
		{"three", "3"},
		{"four", "4"},
		{"five", "5"},
		{"six", "6"},
		{"seven", "7"},
		{"eight", "8"},
		{"nine", "9"},
	}

	lines := strings.Split(fileContent, "\n")
	re := regexp.MustCompile("[0-9]")

	sum := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		words := getWords(replacements, line)

		// get the digits
		digits := re.FindAllString(line, -1)
		if len(digits) <= 0 {
			continue
		}

		// get the first and last digit misspelled
		first := digits[0]
		last := digits[len(digits)-1]

		if len(words) > 0 {
			// replace first occurrence
			firstWord := replacements[words[0]]
			if strings.Index(line, firstWord[0]) < strings.Index(line, first) {
				first = firstWord[1]
			}

			// replace last occurrence
			lastWord := replacements[words[len(words)-1]]
			if strings.LastIndex(line, lastWord[0]) > strings.LastIndex(line, last) {
				last = lastWord[1]
			}
		}

		// calculate the sum
		crc := fmt.Sprintf("%s%s", first, last)
		num, err := strconv.Atoi(crc)
		if err != nil {
			continue
		}
		sum += num
	}

	fmt.Printf("= %d", sum)
}

func getWords(replacements [][]string, line string) []int {
	words := []int{}
	for i := 0; i < len(line); i++ {
		for id, replacement := range replacements {
			if strings.HasPrefix(line[i:], replacement[0]) {
				words = append(words, id)
			}
		}
	}

	return words
}
