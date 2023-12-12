package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Line struct {
	txt      string
	checksum []int
	possible int
}

func main() {
	lines := parse()

	sum := 0
	for i := range lines {
		solve(&lines[i], lines[i].txt)
		sum += lines[i].possible
	}

	fmt.Println("Part 1 =", sum)
}

func solve(line *Line, current string) {
	crc := calculateCRC(current)

	next := strings.Index(current, "?")
	if next == -1 {
		if equal(crc, line.checksum) {
			line.possible++
		}
		return
	}

	solve(line, current[:next]+"."+current[next+1:])
	solve(line, current[:next]+"#"+current[next+1:])
}

func calculateCRC(txt string) []int {
	checksum := []int{}

	count := 0
	for _, c := range txt {
		if string(c) == "." {
			if count > 0 {
				checksum = append(checksum, count)
				count = 0
			}
			continue
		}
		count++
	}

	if count > 0 {
		checksum = append(checksum, count)
		count = 0
	}

	return checksum
}

func parse() []Line {
	lines := []Line{}
	for _, s := range strings.Split(strings.TrimSpace(input), "\n") {
		val := strings.Split(strings.TrimSpace(s), " ")

		check := []int{}

		for _, n := range strings.Split(val[1], ",") {
			num, _ := strconv.Atoi(n)
			check = append(check, num)
		}

		lines = append(lines, Line{
			txt:      val[0],
			checksum: check,
		})
	}

	return lines
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
