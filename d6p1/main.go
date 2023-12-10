package main

import (
	"bufio"
	"fmt"
	"log"
	"moui72/aoc-2023/util"
	"regexp"
	"strings"
)

var path string = util.RelativePathTo("inputs/d6.txt")

func determineResult(charge, totalTime int) int {
	return charge * (totalTime - charge)
}

func findMinCharge(time, recordDistanceToBeat int) int {
	charge := time / 4 
	for determineResult(charge, time) >= recordDistanceToBeat {
		charge -= 1
	}
	if charge < time / 4 {
		return charge
	}
	for determineResult(charge, time) <= recordDistanceToBeat {
		charge += 1
	}
	return charge
}


func parseRaces(scanner *bufio.Scanner) []map[string]int {
	lines := make([][]string, 0, 2)
	multipleSpaces := regexp.MustCompile(" +")
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		text = multipleSpaces.ReplaceAllString(text, " ")
		_, line, ok := strings.Cut(text, ":")
		if ok != true {
			panic(text)
		}
		values := strings.Split(strings.TrimSpace(line), " ")
		lines = append(lines, values)
	}

	races := make([]map[string]int, 0, len(lines[0]))

	for i := 0; i < len(lines[0]); i++ {
		race := map[string]int{"time": util.ParseIntOrRaise(lines[0][i]), "recordDistance": util.ParseIntOrRaise(lines[1][i])}
		races = append(races, race)
	}
	return races
}

func main() {
	fmt.Printf("Loading data from %s\n", path)
	f, scanner, err := util.ReadFileToScanner(path)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	races := parseRaces(scanner)
	margin := 1
	for _, race := range races {
		time := race["time"]
		record := race["recordDistance"]
		minCharge := findMinCharge(time, record)
		maxCharge := time - minCharge
		fmt.Printf("Minimum charge time to beat %d in %d is %d\n", record, time, minCharge)
		fmt.Printf("Maximum charge time to beat %d in %d is %d\n", record, time, maxCharge)
		waysToWin := 1 + maxCharge - minCharge
		fmt.Printf("There are %d ways to win\n", waysToWin)
		margin *= waysToWin
	}
	fmt.Printf("The margin for error is %d\n", margin)
}
