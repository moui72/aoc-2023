package main

import (
	"fmt"
	"log"
	"moui72/aoc-2023/util"
	"slices"
	"strings"
)

var path string = util.RelativePathTo("inputs/d7.txt")

func strongerHand(hand1, hand2 string) bool {
	if hand1 == hand2 {
		println("identical hands found: " + hand1 + " " + hand2)
	}
	handRank := map[string]int{
		"5ofKind": 7,
		"4ofKind": 6,
		"fullHouse": 5,
		"3ofKind": 4,
		"2Pair": 3,
		"pair": 2,
		"highCard": 1,
	}
	cardRank := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 0,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'1': 1,
	}
	hand1type := determineHandType(hand1)
	hand2type := determineHandType(hand2)
	if handRank[hand1type] != handRank[hand2type] {
		gt := "!="
		if handRank[hand1type] > handRank[hand2type] {
			gt = ">"
		} else {
			gt = "<"
		}
		fmt.Printf("...%s [%s] %s %s [%s]\n", hand1, hand1type, gt, hand2, hand2type)
		return handRank[hand1type] > handRank[hand2type]
	}
	for i, c1 := range hand1 {
		c2 := rune(hand2[i])
		strength1 := cardRank[c1]
		strength2 := cardRank[c2]
		if c1 != c2 {
			gt := "!="
				if strength1 > strength2 {
					gt = ">"
				} else {
					gt = "<"
				}
				fmt.Printf("...%#U [%s] %s %#U [%s]\n", c1, hand1, gt, c2, hand2)
				return strength1 > strength2
		}
	}
	return false
}

func determineHandType(hand string) string {
	counts := map[rune]int{}
	maxSame := 0
	maxSet := 'Z'
	secondMaxSame := 0
	for _, card := range hand {
		counts[card] += 1
		if counts[card] > maxSame {
			maxSame = counts[card]
			maxSet = card
		} else if counts[card] > secondMaxSame {
			secondMaxSame = counts[card]
		}
	}
	if counts['J'] > 0 {
		if maxSet == 'J' {
			maxSame += secondMaxSame
		} else {
			maxSame += counts['J']
		}
	}
	switch maxSame {
	case 5:
		return "5ofKind"
	case 4:
		return "4ofKind"
	case 3:
		if secondMaxSame > 1 {
			return "fullHouse"
		} else {
			return "3ofKind"
		}
	case 2:
		if secondMaxSame > 1 {
			return "2Pair"
		} else {
			return "pair"
		}
	default:
		return "highCard"
	} 
}

func main() {
	fmt.Printf("Loading data from %s\n", path)
	f, scanner, err := util.ReadFileToScanner(path)

	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	games := []map[string]string{}

	for scanner.Scan() {
		text := scanner.Text()
		hand, bid, _ := strings.Cut(text, " ")
		handType := determineHandType(hand)
		game := map[string]string{"hand": hand, "bid": bid, "type": handType}
		games = append(games, game)
	}
	sortedGames := []map[string]string{}
	for _, game := range games {
		inserted := false
		for i := 0; i < len(sortedGames); i++ {
			if strongerHand(game["hand"], sortedGames[i]["hand"]) {
				// fmt.Printf("%s > %s\n", game["hand"],  sortedGames[0]["hand"])
				sortedGames = append(sortedGames[:i], append([]map[string]string{game}, sortedGames[i:]...)...)
				inserted = true
				break
			}
		}
		if !inserted {
			sortedGames = append(sortedGames, game)
		}
	}
	slices.Reverse(sortedGames)
	fmt.Println("Rank\tHand\tBid\tHand Type")
	fmt.Println("----\t----\t---\t---------")
	totalWinnings := 0
	for rank, hand := range sortedGames {
		fmt.Printf("%d\t%s\t%s\t[%s]\n", rank + 1, hand["hand"], hand["bid"], hand["type"])
		totalWinnings += (rank + 1) * util.ParseIntOrRaise(hand["bid"])
	}
	fmt.Printf("Total winnings: %d\n", totalWinnings)

}
