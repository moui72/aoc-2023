package main

import (
	"fmt"
	"log"
	"moui72/aoc-2023/util"
	"strconv"
	"strings"
)

var path string = util.RelativePathTo("inputs/d3.txt")
var digits string = "0123456789"

func main() {
	f, scanner := util.ReadFileToScanner(path)
	defer f.Close()

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		number := ""
		var digit rune
		for _, character := range line {
			if strings.ContainsRune(digits, character) {
				digit = character
				if len(number) == 0 {
					number += string(digit)
				}
			}
		}
		number += string(digit)
		n, err := strconv.Atoi(number)
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
