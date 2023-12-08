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
 * Day 8: Haunted Wasteland - Part 1
 * url: https://adventofcode.com/2023/day/8
 */
func main() {
	cmds, nodes := parseNodes()

	steps := 0
	current := "AAA"
	for {
		if current == "ZZZ" {
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

	fmt.Println("Part 1 =", steps)
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
