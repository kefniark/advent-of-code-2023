package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	dst    int
	src    int
	length int
}

//go:embed input.txt
var input string

/**
 * Day 4: Scratchcards - Part 1
 * url: https://adventofcode.com/2023/day/5
 */
func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	current := ""
	seeds := []int{}
	rules := map[string][]*Rule{}

	for _, s := range lines {
		s = strings.TrimSpace(s)
		if strings.Contains(s, ":") {
			current = strings.Replace(strings.Split(s, ":")[0], " map", "", -1)
			if current == "seeds" {
				re := regexp.MustCompile(`\d+`)
				numbers := []int{}
				for _, val := range re.FindAllStringSubmatch(strings.Split(s, ":")[1], -1) {
					num, _ := strconv.Atoi(val[0])
					numbers = append(numbers, num)
				}
				seeds = numbers
				continue
			}
		}

		rule := parseNumbers(s)
		if rule != nil {
			rules[current] = append(rules[current], rule)
		}
	}

	min := 0
	for _, val := range seeds {
		res := resolve(rules, val)
		if min == 0 || res["location"] < min {
			min = res["location"]
		}
	}

	fmt.Println("Part 1 = ", min)
}

func resolve(rules map[string][]*Rule, seed int) map[string]int {
	res := map[string]int{"seed": seed}
	resolveRule(rules, seed, "seed", res)
	return res
}

func resolveRule(rules map[string][]*Rule, id int, name string, res map[string]int) {
	for key, ranges := range rules {
		if !strings.Contains(key, name+"-") {
			continue
		}

		keySplit := strings.Split(key, "-")
		to := keySplit[2]

		idMap := id
		for _, rule := range ranges {
			if id >= rule.src && id <= rule.src+rule.length {
				idMap = rule.dst + (id - rule.src)
				break
			}
		}

		res[to] = idMap
		resolveRule(rules, idMap, to, res)
	}
}

func parseNumbers(s string) *Rule {
	re := regexp.MustCompile(`\d+`)
	numbers := []int{}
	for _, val := range re.FindAllStringSubmatch(s, -1) {
		num, _ := strconv.Atoi(val[0])
		numbers = append(numbers, num)
	}

	if len(numbers) != 3 {
		return nil
	}

	return &Rule{
		dst:    numbers[0],
		src:    numbers[1],
		length: numbers[2],
	}
}
