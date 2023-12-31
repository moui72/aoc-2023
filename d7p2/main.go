package main

import (
	"fmt"
	"moui72/aoc-2023/util"
)

var path string = util.RelativePathTo("inputs/[FILENAME].txt")

func main() {
	fmt.Printf("Loading data from %s\n", path)
	f, scanner := util.ReadFileToScanner(path)

	defer f.Close()
	for scanner.Scan() {
		text := scanner.Text()
		println(text)
	}
}
