package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

type Node struct {
	name  string
	left  string
	right string
}

/**
 * Day 8: Haunted Wasteland - Part 2
 * url: https://adventofcode.com/2023/day/8
 */
func main() {
	cmds, nodes := parseNodes()

	res := []int{}
	for k, _ := range nodes {
		if string(k[2:]) != "A" {
			continue
		}

		res = append(res, findPathLength(k, cmds, nodes))
	}

	fmt.Println("Part 2 =", lcm(res[0], res[1], res[2:]...))
}

func findPathLength(current string, cmds []string, nodes map[string]*Node) int {
	steps := 0
	for {
		if current[2:] == "Z" {
			break
		}

		dir := cmds[steps%len(cmds)]
		steps += 1

		if dir == "L" {
			current = nodes[current].left
		} else {
			current = nodes[current].right
		}
	}

	return steps
}

func parseNodes() ([]string, map[string]*Node) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	cmds := strings.Split(strings.TrimSpace(lines[0]), "")
	r := regexp.MustCompile(`(\w*) = \((\w*), (\w*)\)`)

	mapNodes := make(map[string]*Node)
	for _, s := range lines[1:] {
		res := r.FindAllStringSubmatch(strings.TrimSpace(s), -1)
		if len(res) == 0 {
			continue
		}

		mapNodes[res[0][1]] = &Node{
			name:  res[0][1],
			left:  res[0][2],
			right: res[0][3],
		}
	}

	return cmds, mapNodes
}

// find Least Common Multiple "lcm(a,b) = a * b / gcd(a, b)" and repeat for extra numbers
func lcm(a, b int, rest ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(rest); i++ {
		result = lcm(result, rest[i])
	}

	return result
}

// find Greatest Common Divisor (gcd) just repeat modulo until zero
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}
