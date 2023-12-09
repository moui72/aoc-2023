package main

import (
	"fmt"
	"log"
	"moui72/aoc-2023/util"
	"strconv"
)

var path string = util.RelativePathTo("inputs/d3.txt")

var addressFmt string = "%d,%d-%d"

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

func neighborToStar(row, col int, numberLocationsRow [][]int) []string {
	uniqNeighbors := make(map[string]string)
	neighbors := []string{}
	for _, pos := range numberLocationsRow {
		colStart := pos[0]
		colEnd := pos[1]
		addr := fmt.Sprintf(addressFmt, row, colStart, colEnd)
		if col >= colStart-1 && col <= colEnd+1 {
			_, exists := uniqNeighbors[addr]
			if !exists {
				fmt.Printf("\t> Found number at: %s\n", addr)
				uniqNeighbors[addr] = addr
				neighbors = append(neighbors, addr)
			}
		}
	}
	return neighbors
}

func findGearNumberAddresses(row, col int, numberLocations [][][]int) (bool, string, string) {
	addrs := make([]string, 0, 3)
	for i := -1; i < 2; i++ {
		scanRow := row + i
		if scanRow >= 0 && scanRow < len(numberLocations) {
			addrs = append(addrs, neighborToStar(scanRow, col, numberLocations[scanRow])...)
			if len(addrs) > 2 {
				fmt.Printf("\t> Found more than 2 numbers adjacent to * at %d, %d\n", row, col)
				return false, "", ""
			}
		}
	}
	if len(addrs) == 2 {
		fmt.Printf("\t> Found exactly 2 numbers adjacent to * at %d, %d\n", row, col)
		return true, addrs[0], addrs[1]
	}
	fmt.Printf("\t> Found fewer than 2 numbers adjacent to * at %d, %d\n", row, col)
	return false, "", ""
}

func main() {
	f, scanner, err := util.ReadFileToScanner(path)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	starLocations := make([][]int, util.LineCount(path))
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
				if '*' == character {
					starLocations[row] = append(starLocations[row], pos)
				}
			}
			if (numberEnd > 0 || pos == len(line)-1) && numberString != "" {
				numberEnd = coalescePositive(numberEnd, pos)
				address := fmt.Sprintf(addressFmt, row, numberStart, numberEnd)
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
	for row, stars := range starLocations {
		for _, col := range stars {
			fmt.Printf("Checking if star at %d, %d has two neighbors\n", row, col)
			isGear, addr1, addr2 := findGearNumberAddresses(row, col, numberLocations)
			if isGear {
				num1 := numbersByAddress[addr1]
				num2 := numbersByAddress[addr2]
				fmt.Printf("\t> Ratio of %d (%s) * %d (%s) is %d\n", num1, addr1, num2, addr2, num1 * num2 )
				sum += numbersByAddress[addr1] * numbersByAddress[addr2]
			}
		}
	}
	fmt.Printf("Sum of all gear ratios is %d\n", sum)
}
