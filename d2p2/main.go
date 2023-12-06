package main

import (
	"fmt"
	"log"
	"moui72/aoc-2023/util"
	"strconv"
	"strings"
)

var path string = "../inputs/d2p1.txt"

func powerOfSet(cubes map[string]int) int {
	result := 1
	for _, count := range cubes {
		result *= count
	}
	return result
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
		sum += powerOfSet(games[gameId])

	}

	fmt.Printf("Sum powers = %d\n", sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
