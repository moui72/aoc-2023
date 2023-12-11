package main

import (
	"fmt"
	"moui72/aoc-2023/util"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
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

	scanner.Scan()
	directions := scanner.Text()
	desertMap := map[string]map[rune]string{}

	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			node, children, _ := strings.Cut(text, " = ")
			L, R, _ := strings.Cut(children, ", ")
			desertMap[node] = map[rune]string{'L': L[1:], 'R': R[:len(R)-1]}
		}
	}
	paths := []string{}
	for node, _ := range desertMap {
		if strings.HasSuffix(node, "A") {
			paths = append(paths, node)
		}
	}
	spew.Dump(paths)
	lengths := make([]int, len(paths))
	steps := 0
	solved := 0
	directionsComplete := 0
	for solved < len(paths) {
		fmt.Printf(
			"After completing the direction set %d times, %d of the %d paths reached the end\n",
			directionsComplete,
			solved,
			len(paths),
		)
		for _, d := range directions {
			steps += 1
			for i, loc := range paths {
				paths[i] = desertMap[loc][d]
				if strings.HasSuffix(paths[i], "Z") && lengths[i] == 0 {
					fmt.Printf("Path %d took %d steps", i, steps)
					lengths[i] = steps
					solved += 1
				}
				if solved == len(paths) {
					break
				}
			}
			if solved == len(paths) {
				break
			}
		}
		directionsComplete += 1
	}
	spew.Dump(lengths)
	solution := 0
	if len(lengths) == 1 {
		solution = lengths[0]
	} else if len(lengths) == 2 {
		solution = LCM(lengths[0], lengths[1])
	} else {
		solution = LCM(lengths[0], lengths[1], lengths[2:]...)
	}
	fmt.Printf("\nIt took %d steps to reach all paths ending in Z\n", solution)
}
