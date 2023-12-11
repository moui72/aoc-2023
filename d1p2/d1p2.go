package main

import (
	"fmt"
	"log"
	"moui72/aoc-2023/util"
	"strconv"
	"strings"
)

var path string = util.RelativePathTo("inputs/d3.txt")



func main() {
	digits := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"0": 0,
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
		"zero": 0,
	}
	f, scanner := util.ReadFileToScanner(path)
	defer f.Close()

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		first := [2]int{9999999,9999999}
		last := [2]int{-1, -1}
		for key, value := range digits {
			firstIndex := strings.Index(line, key)
			if firstIndex > -1 && firstIndex < first[0] {
				first[0] = firstIndex
				first[1] = value
			}
			lastIndex := strings.LastIndex(line, key)
			if lastIndex > last[0] {
				last[0] = lastIndex
				last[1] = value
			}
		}
		n, err := strconv.Atoi(strconv.Itoa(first[1]) + strconv.Itoa(last[1]))
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		total += n
	}
	fmt.Printf("Calibration total is %d\n", total)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
