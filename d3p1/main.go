package main

import (
	"fmt"
	"log"
	"moui72/aoc-2023/util"
	"strconv"
)

var path string = "../inputs/d3.txt"

func coalescePositive(a, b int) int {
	if a > 0 {
			return a
	}
	return b
}

func isNumeral(r rune) bool {
	if r < '0' || r > '9' {
		return false
	}
	return true
}

func neighborInRow(colStart, colEnd int, symbolMapRow []int) bool {
	for _, pos := range symbolMapRow {
		if pos >= colStart - 1 && pos <= colEnd + 1{
			return true
		}
	}
	return false
}

func isPart(row, colStart, colEnd int, symbolMap [][]int) bool {
	for i := -1; i < 2; i++ {
		scanRow := row + i
		if scanRow >= 0 && scanRow < len(symbolMap) && neighborInRow(colStart, colEnd, symbolMap[row + i]) {
			return true
		}
	}
	return false
}

func main() {

	f, scanner, err := util.ReadFileToScanner(path)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	symbolLocations := make([][]int, util.LineCount(path))
	numberLocations := make([][][]int, util.LineCount(path))
	numbersByAddress := map[string]int{}
	numberString := ""
	row := 0
	numberStart := -1
	numberEnd := -1
	for scanner.Scan() {
		line := scanner.Text()
		for pos, character := range line {
			if isNumeral(character) {
				if numberStart < 0 {
					numberStart = pos
				}
				numberString += string(character)
			} else {
				if numberStart >= 0 {
					numberEnd = pos - 1
				}
				if '.' != character {
					symbolLocations[row] = append(symbolLocations[row], pos)
				}
			}
			if (numberEnd > 0 || pos == len(line) - 1) && numberString != "" {
				numberEnd = coalescePositive(numberEnd, pos)
				address := fmt.Sprintf("%d,%d-%d", row, numberStart, numberEnd)
				num, err := strconv.Atoi(numberString)
				numbersByAddress[address] = num
				if err != nil {
					panic(err)
				}
				numberLocations[row] = append(numberLocations[row], []int{numberStart, numberEnd})
				numberString = ""
				numberStart = -1
				numberEnd = -1
			}
		}
		row += 1
	}
	sum := 0
	for row, numberSpans := range numberLocations {
		for _, span := range numberSpans {
			if isPart(row, span[0], span[1], symbolLocations) {
				address := fmt.Sprintf("%d,%d-%d", row, span[0], span[1])
				sum += numbersByAddress[address]
			}
		}
	}
	fmt.Printf("Sum of all part numbers is %d\n", sum)
}
