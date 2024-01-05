package day07

import (
	"fmt"
	"sort"
	"strings"

	"github.com/orlando-pt/aoc/2023/utils"
	"golang.org/x/exp/slices"
)

const cards = "AKQJT98765432"

func Part1(lines []string) int {
	game := parseLines(lines)

	game.sort()
	fmt.Println(game)

	ret := 0

	for i, h := range game.hands {
		ret += h.bid * (i + 1)
	}
	return ret
}

func Part2(lines []string) int {
	return 0
}

func (game Game) sort() {
	hands := game.hands
	sort.Slice(hands, func(i, j int) bool {
		code1 := findStrongestHand(hands[i])
		code2 := findStrongestHand(hands[j])

		if code1 == code2 {
			return !hands[i].isStronger(hands[j])
		}

		return code1 < code2
	})
}

func (h1 Hand) isStronger(h2 Hand) bool {
	for i := 0; i < len(cards); i++ {
		h1Index := strings.Index(cards, string(h1.cards[i]))
		h2Index := strings.Index(cards, string(h2.cards[i]))

		if h1Index != h2Index {
			return h1Index < h2Index
		}
	}
	return false
}

func findStrongestHand(hand Hand) int {
	code := 0

	var repeatedCards []int
	for _, c := range distinct(hand.cards) {
		repeatedCards = append(repeatedCards, strings.Count(hand.cards, string(c)))
	}

	switch len(repeatedCards) {
	case 5:
		// all different
		code = 1
	case 4:
		// one pair
		code = 2
	case 3:
		// two pairs or three of a kind
		if slices.Contains(repeatedCards, 3) {
			code = 4
		} else {
			code = 3
		}
	case 2:
		// full house or four of a kind
		if slices.Contains(repeatedCards, 4) {
			code = 7
		} else {
			code = 6
		}
	case 1:
		// five of a kind
		code = 8
	}

	return code
}

func distinct(s string) string {
	var distinct string
	for _, c := range s {
		if strings.Index(distinct, string(c)) == -1 {
			distinct += string(c)
		}
	}
	return distinct
}

type Hand struct {
	cards string
	bid   int
}

type Game struct {
	hands []Hand
}

func parseLines(lines []string) Game {
	var hands []Hand

	for _, line := range lines {
		handString := strings.Split(line, " ")
		hands = append(hands, Hand{cards: handString[0], bid: utils.StrToInt(handString[1])})
	}

	return Game{hands: hands}
}
