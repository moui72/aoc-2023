package main

import (
	"fmt"
	"moui72/aoc-2023/util"
	"os"
	"slices"
	"strings"
)

func replaceAtIndex(in, rep string, index int) string {
	return in[:index] + rep + in[index+1:]

}

func neighbors(row, col int, m []string) (rune, rune, rune, rune) {
	up, right, down, left := 'X', 'X', 'X', 'X'
	if row > 0 {
		up = rune(m[row-1][col])
	}
	if col > 0 {
		left = rune(m[row][col-1])
	}
	if row < len(m)-1 {
		down = rune(m[row+1][col])
	}
	if col < len(m[row])-1 {
		right = rune(m[row][col+1])
	}
	return up, right, down, left
}

func identifyS (up, right, down, left rune) rune {
	con1, con2 := connectionsFromNeighbors(up, right, down, left) 
	return identifySFromConnections(con1, con2)
}

func connectionsFromNeighbors(up, right, down, left rune) (string, string) {
	cxns := []string{}
	connectsUp := []rune{'|', 'F', '7'}
	connectsRight := []rune{'-', '7', 'J'}
	connectsDown := []rune{'|', 'L', 'J'}
	connectsLeft := []rune{'-', 'F', 'L'}
	if slices.Contains(connectsUp, up) {
		cxns = append(cxns, "up")
	}
	if slices.Contains(connectsRight, right) {
		cxns = append(cxns, "right")
	}
	if slices.Contains(connectsDown, down) {
		cxns = append(cxns, "down")
	}
	if slices.Contains(connectsLeft, left) {
		cxns = append(cxns, "left")
	}
	if len(cxns) != 2 {
		panic(fmt.Sprintf("what the heck %d connections????", len(cxns)))
	}
	return cxns[0], cxns[1]
}

func identifySFromConnections(con1, con2 string) rune {
	identifyS := map[string]map[string]rune{
		"up": {
			"down": '|',
			"left": 'J',
			"right": 'L',
		},
		"down": {
			"up": '|',
			"left": '7',
			"right": 'F',
		},
		"left": {
			"up": 'J',
			"right": '-',
			"down": '7',
		},
		"right": {
			"left": '-',
			"up": 'L',
			"down": 'F',
		},
	}
	return identifyS[con1][con2]
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
	simpleMap := []string{}
	inOutMap := []string{}
	start := []int{}
	for scanner.Scan() {
		text := scanner.Text()
		pipeMap = append(pipeMap, text)
		simpleMap = append(simpleMap, strings.Repeat("+", len(text)))
		inOutMap = append(inOutMap, strings.Repeat("o", len(text)))
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
			s := string(identifyS(up, right, down, left))
			simpleMap[row] = replaceAtIndex(simpleMap[row], s, col)
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
			steps += 1
		}
		loc = rune(pipeMap[row][col])
		if loc != 'S' {
			simpleMap[row] = replaceAtIndex(simpleMap[row], string(loc), col)
		}
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
			panic("reached impossible symbol " + string(loc))
		}
		steps += 1
	}
	inside := 0
	println("\t0123456789")
	for row, symbols := range simpleMap {
		fmt.Printf("%d\t%s\n", row, symbols)
	}
	for row, symbols := range simpleMap {
		for col, symbol := range symbols {
			if symbol == '+' {
				crossings := map[string]int{
					"up": 0,
					"right": 0,
					"down": 0,
					"left": 0,
				}
				for i := row-1; i >= 0; i-- {
					if simpleMap[i][col] != '+' {
						crossings["up"]++
					}
				}
				for i := col+1; i < len(symbols)-1; i++ {
					if simpleMap[row][i] != '+' {
						crossings["right"]++
					}
				}
				for i := row+1; i < len(simpleMap)-1; i++ {
					if simpleMap[i][col] != '+' {
						crossings["down"]++
					}
				}
				for i := col-1; i >= 0; i-- {
					if simpleMap[row][i] != '+' {
						crossings["left"]++
					}
				}
				oddCrossings := 0
				for _, count := range crossings {
					if count == 0 {
						oddCrossings = 0
						break
					}
					if count % 2 == 1 {
						oddCrossings++
					}
				}
				if oddCrossings > 0 {
					inOutMap[row] = replaceAtIndex(inOutMap[row], "i", col)
					inside++
					fmt.Printf("Tile at position %d, %d is inside\n", row, col)
				}
			} else {
				inOutMap[row] = replaceAtIndex(inOutMap[row], "+", col)
			}
		}
	}
	println("\t0123456789")
	for row, symbols := range inOutMap {
		fmt.Printf("%d\t%s\n", row, symbols)
	}
	fmt.Printf("\nThere are %d tiles inside the loop\n", inside)
}
