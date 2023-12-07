package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var kind = []string{"H", "1P", "2P", "3", "F", "4", "5"}
var cards = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

type Hand struct {
	cards      []string
	bid        int
	kind       int
	cardValues []int
}

/**
 * Day 7: Camel Cards - Part 1
 * url: https://adventofcode.com/2023/day/7
 */
func main() {
	hands := sortHand(parseHands())

	sum := 0
	for i, hand := range hands {
		sum += (i + 1) * hand.bid
		fmt.Println("-", i, hand.bid, hand.kind, hand.cards)
	}

	fmt.Println("Part 1 =", sum)
}

func parseHands() []Hand {
	hands := []Hand{}
	for _, s := range strings.Split(strings.TrimSpace(input), "\n") {
		str := strings.Split(s, " ")

		num, _ := strconv.Atoi(strings.TrimSpace(str[1]))
		hand := Hand{cards: strings.Split(str[0], ""), bid: num}
		evaluateHand(&hand)

		hands = append(hands, hand)
	}
	return hands
}

func sortHand(hands []Hand) []Hand {
	slices.SortFunc(hands, func(a, b Hand) int {
		if a.kind == b.kind {
			for i := 0; i < len(a.cardValues); i++ {
				if a.cardValues[i] == b.cardValues[i] {
					continue
				}
				return a.cardValues[i] - b.cardValues[i]
			}
		}
		return a.kind - b.kind
	})

	return hands
}

func evaluateHand(hand *Hand) {
	count := map[int]int{}
	values := []int{}
	for _, card := range hand.cards {
		cardValue := slices.Index(cards, card)
		count[cardValue]++

		values = append(values, cardValue)
	}

	kinds := map[string]int{"5": 0, "4": 0, "F": 0, "3": 0, "2P": 0, "1P": 0, "H": 0}
	for _, cpt := range count {
		if cpt == 5 {
			kinds["5"] += 1
		} else if cpt == 4 {
			kinds["4"] += 1
		} else if cpt == 3 {
			kinds["3"] += 1
		} else if cpt == 2 {
			kinds["1P"] += 1
		}
	}

	if kinds["3"] == 1 && kinds["1P"] == 1 {
		kinds["3"] -= 1
		kinds["1P"] -= 1
		kinds["F"] += 1
	}

	if kinds["1P"] > 1 {
		kinds["1P"] -= 2
		kinds["2P"] += 1
	}

	maxKind := 0
	for idx, val := range kinds {
		if val == 0 {
			continue
		}
		maxKind = slices.Index(kind, idx)
	}

	hand.kind = maxKind
	hand.cardValues = values
}
