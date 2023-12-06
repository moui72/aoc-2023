package main

import (
	"fmt"
	"log"
	"moui72/aoc-2023/util"
	"strconv"
	"strings"
)

var path string = "../inputs/d2p1.txt"

func isPossible(rules map[string]int, game map[string]int) (bool, string) {
	for color, max := range rules {
		if game[color] > max {
			return false, fmt.Sprintf("too many %s cubes, found %d and max is %d", color, max, game[color])
		}
	}
	return true, ""
}

func findMaxColors(text string) map[string]int {
	result := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, turn := range strings.Split(text, ";") {
		for _, colorCount := range strings.Split(turn, ",") {
			count, color, _ := strings.Cut(strings.Trim(colorCount, " "), " ")
			parsedCount, err := strconv.Atoi(count)
			if err != nil {
				panic(err)
			}
			if result[color] < parsedCount {
				result[color] = parsedCount
			}
		}
	}
	return result
}

func main() {
	rules := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	f, scanner, err := util.ReadFileToScanner(path)
	defer f.Close()

	games := map[int]map[string]int{}

	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		gameName, gameResult, _ := strings.Cut(line, ":")
		gameId, err := strconv.Atoi(strings.TrimPrefix(gameName, "Game "))
		if err != nil {
			panic(err)
		}
		games[gameId] = findMaxColors(gameResult)
		gameIsPossible, reason := isPossible(rules, games[gameId])
		if gameIsPossible {
			sum += gameId
		} else {
			fmt.Printf("Game %d is not possible: %s\n\t%s\n\n", 
			gameId, 
			reason,
			strings.Trim(gameResult, " "), 
		)

		}
	}

	fmt.Printf("Sum of possible game IDs = %d\n", sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
