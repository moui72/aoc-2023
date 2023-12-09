package main

import (
	"fmt"
	"log"
	"moui72/aoc-2023/util"
	"regexp"
	"slices"
	"strings"
)

var path string = util.RelativePathTo("inputs/d4.txt")

func trimAndSplitOnSpace(s string) []string {
	return strings.Split(strings.TrimSpace(s), " ")
}

func parseCard(line string) ([]string, []string) {
	multipleSpaces := regexp.MustCompile(" +")
	line = multipleSpaces.ReplaceAllString(line, " ")
	_, data, success := strings.Cut(line, ":")
	if success != true {
		panic("Could not split on ':'")
	}
	winningNumbers, playedNumbers, success := strings.Cut(data, "|")
	if success != true {
		panic("Could not split on '|'")
	}
	winning := trimAndSplitOnSpace(winningNumbers)
	played := trimAndSplitOnSpace(playedNumbers)
	return played, winning
}

func scoreCard(card string) int {
	played, winning := parseCard(card)
	score := 0
	for _, num := range played {
		if slices.Contains(winning, num) {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}
	return score
}

func main() {
	fmt.Printf("Loading data from %s\n", path)
	f, scanner, err := util.ReadFileToScanner(path)

	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	grandScore := 0
	card := 1
	for scanner.Scan() {
		score := scoreCard(scanner.Text())
		grandScore += score
		fmt.Printf("Card %d is worth %d points\n", card, score)
		card += 1
	}
	fmt.Printf("The total score is %d\n", grandScore)
}
