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
			score += 1
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
	card := 1
	originalCards := map[int]int{}
	toScore := []int{}
	for scanner.Scan() {
		score := scoreCard(scanner.Text())
		originalCards[card] = score
		toScore = append(toScore, card)
		card += 1
	}
	grandScore := len(toScore)
	for i := 0; i < grandScore; i++ {
		card := toScore[i]
		cardScore := originalCards[card]
		for j := 0; j < cardScore; j++ {
			toScore = append(toScore, card + j + 1)
		}
		grandScore = len(toScore)
	}
	fmt.Printf("Total number of cards is %d\n", grandScore)
}
