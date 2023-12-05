package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
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
 * Day 4: Scratchcards - Part 2
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
	locations := []int{}
	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		length := seeds[i+1]
		resolveMatchingInterval(rules, "seed", start, start+length, &locations)
	}

	for _, val := range locations {
		if min == 0 || val < min {
			min = val
		}
	}

	fmt.Println("Part 2 = ", min)
}

/*
 * Recursive method which for a given range of id and category, split in different interval of ids and categories.
 */
func resolveMatchingInterval(rules map[string][]*Rule, name string, start int, end int, locations *[]int) {
	r := filterValidRules(rules, name, start, end)

	// Exit condition, we found a location range (we only take the lowest one)
	if name == "location" {
		*locations = append(*locations, start)
	}

	// We split the range based on the rules
	for key, ranges := range r {
		sort.Slice(ranges, func(i, j int) bool {
			return ranges[i].start < ranges[j].start
		})

		s := start

		for _, ran := range ranges {
			rule := ran.rule

			maxStart := max(s, rule.src)
			minEnd := min(end, rule.src+rule.length)
			if s < maxStart {
				resolveMatchingInterval(rules, key, s, maxStart-1, locations)
			}

			resolveMatchingInterval(rules, key, maxStart+ran.offset, minEnd+ran.offset, locations)
			s = minEnd
		}

		if s < end {
			resolveMatchingInterval(rules, key, s, end, locations)
		}
	}
}

type Range struct {
	start  int
	end    int
	rule   *Rule
	offset int
}

func filterValidRules(rules map[string][]*Rule, name string, start int, end int) map[string][]Range {
	res := map[string][]Range{}
	for key, ranges := range rules {
		if !strings.Contains(key, name+"-") {
			continue
		}

		keySplit := strings.Split(key, "-")
		to := keySplit[2]
		if res[to] == nil {
			res[to] = []Range{}
		}

		for _, rule := range ranges {
			a1 := start
			b1 := end
			a2 := rule.src
			b2 := rule.src + rule.length

			if b2 < a1 || a2 > b1 {
				continue
			}

			res[to] = append(res[to], Range{
				start:  max(a1, a2),
				end:    min(b1, b2) - 1,
				rule:   rule,
				offset: rule.dst - rule.src,
			})
		}
	}
	return res
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
