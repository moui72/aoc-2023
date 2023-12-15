package main

import (
	"fmt"
	"moui72/aoc-2023/util"
	"os"
	"slices"
	"strings"
)

func neighbors(row, col int, m []string) (rune, rune, rune, rune) {
	up, right, down, left := 'X', 'X', 'X', 'X'
	if row > 0 {
		up = rune(m[row-1][col])
	}
	if col > 0 {
		left = rune(m[row][col-1])
	}
	if row < len(m) - 1 {
		down = rune(m[row+1][col])
	}
	if col < len(m[row]) - 1 {
		right = rune(m[row][col+1])
	}
	return up, right, down, left
}

func main() {
	arg1 := os.Args[1]
	path := ""
	if strings.Contains(arg1, "/") {
		path = util.RelativePathTo(arg1)
	} else {
		path = util.PathFromFileName(os.Args[1])
	}

	fmt.Printf("Loading data from %s\n", path)
	f, scanner := util.ReadFileToScanner(path)
	defer f.Close()
	pipeMap := []string{}
	start := []int{}
	for scanner.Scan() {
		text := scanner.Text()
		println(text)
		pipeMap = append(pipeMap, text)
		col := strings.Index(text, "S")
		if col >= 0 {
			row := len(pipeMap) - 1
			start = append(start, row, col)
		}
	}
	steps := 0
	loc := 'S'
	connectsUp := []rune{'|', 'F', '7'}
	connectsDown := []rune{'|', 'L', 'J'}
	connectsLeft := []rune{'-', 'F', 'L'}
	connectsRight := []rune{'-', '7', 'J'}
	prevStep := 'X'
	fmt.Printf("Starting at %d, %d\n", start[0], start[1])
	row := start[0]
	col := start[1]
	for steps == 0 || loc != 'S' {
		if loc == 'S' {
			up, right, down, left := neighbors(row, col, pipeMap)
			if slices.Contains(connectsUp, up) {
				row -= 1
				prevStep = 'U'
			} else if slices.Contains(connectsRight, right) {
				col += 1
				prevStep = 'R'
			} else if slices.Contains(connectsDown, down) {
				row += 1
				prevStep = 'D'
			} else if slices.Contains(connectsLeft, left) {
				col -= 1
				prevStep = 'L'
			}
			print(string(prevStep))
			steps += 1
		}
		loc = rune(pipeMap[row][col])
		switch loc {
			case '|':
				if prevStep == 'U' {
					row -= 1
					prevStep = 'U'
				} else {
					row += 1
					prevStep = 'D'
				}
			case 'F':
				if prevStep == 'U' {
					col += 1
					prevStep = 'R'
				} else {
					row += 1
					prevStep = 'D'
				}
			case '7': 
				if prevStep == 'U' {
					col -= 1
					prevStep = 'L'
				} else {
					row += 1
					prevStep = 'D'
				}
			case '-': 
				if prevStep == 'L' {
					col -= 1
					prevStep = 'L'
				} else {
					col += 1
					prevStep = 'R'
				}
			case 'J': 
				if prevStep == 'D' {
					col -= 1
					prevStep = 'L'
				} else {
					row -= 1
					prevStep = 'U'
				}
			case 'L': 
				if prevStep == 'D' {
					col += 1
					prevStep = 'R'
				} else {
					row -= 1
					prevStep = 'U'
				}
			case 'S':
				break
			default:
				panic("reached impossible symbol "+ string(loc))
		}
		print(string(prevStep))
		steps += 1
	}
	fmt.Printf("\nThe furthest point is %d steps away\n", steps/2)
}
