package day07

import (
	"sort"
	"strings"

	"github.com/orlando-pt/aoc/2023/utils"
	"golang.org/x/exp/slices"
)

const cards = "AKQJT98765432"
const cardsPart2 = "AKQT98765432J"

func Part1(lines []string) int {
	game := parseLines(lines)

	game.sort()

	ret := 0
	for i, h := range game.hands {
		ret += h.bid * (i + 1)
	}
	return ret
}

func Part2(lines []string) int {
	game := parseLines(lines)

	game.sortPart2()

	ret := 0
	for i, h := range game.hands {
		ret += h.bid * (i + 1)
	}

	return ret
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

func (game Game) sortPart2() {
	hands := game.hands
	sort.Slice(hands, func(i, j int) bool {
		code1 := findStrongestHandPart2(hands[i])
		code2 := findStrongestHandPart2(hands[j])

		if code1 == code2 {
			return !hands[i].isStrongerPart2(hands[j])
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
func (h1 Hand) isStrongerPart2(h2 Hand) bool {
	for i := 0; i < len(cardsPart2); i++ {
		h1Index := strings.Index(cardsPart2, string(h1.cards[i]))
		h2Index := strings.Index(cardsPart2, string(h2.cards[i]))

		if h1Index != h2Index {
			return h1Index < h2Index
		}
	}
	return false
}

func findStrongestHand(hand Hand) int {
	var repeatedCards []int
	for _, c := range distinct(hand.cards) {
		repeatedCards = append(repeatedCards, strings.Count(hand.cards, string(c)))
	}

	return getHandCode(repeatedCards)
}

func findStrongestHandPart2(hand Hand) int {

	distinctCards := distinctNoJ(hand.cards)

    possibleHands := []string{hand.cards}

	// replace cards for J
	if strings.Contains(hand.cards, "J") {
		nrOfJokers := strings.Count(hand.cards, "J")
		for _, c := range distinctCards {
			possibleHands = append(possibleHands, strings.Replace(hand.cards, "J", string(c), nrOfJokers))
		}
	}

	highestHandCode := 0
	for _, h := range possibleHands {
        repeatedCards := []int{}
		for _, c := range distinct(h) {
			repeatedCards = append(repeatedCards, strings.Count(h, string(c)))
		}

		handCode := getHandCode(repeatedCards)
		if handCode > highestHandCode {
			highestHandCode = handCode
		}
	}

	return highestHandCode
}

func getHandCode(repeatedCards []int) int {
	var code int

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

func distinctNoJ(s string) string {
	var distinct string
	for _, c := range s {
		if strings.Index(distinct, string(c)) == -1 && string(c) != "J" {
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
