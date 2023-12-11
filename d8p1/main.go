package main

import (
	"fmt"
	"moui72/aoc-2023/util"
	"os"
	"strings"
)


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
	desertMap := map[string]map[rune]string

	for scanner.Scan() {
		text := scanner.Text()
		node, children, _ := strings.Cut(text, " = ")
		L, R, _ := strings.Cut(children, ", ")
		desertMap[node] = map[rune]string{'L': L[1:], 'R': R[:len(R)-1]}
	}
	loc := "AAA"
	steps := 0
	for loc != "ZZZ" {
		for _, d := range directions {
			steps += 1
			loc = desertMap[loc][d]
		}
	}
}
