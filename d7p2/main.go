package main

import (
	"fmt"
	"log"
	"moui72/aoc-2023/util"
)

var path string = util.RelativePathTo("inputs/[FILENAME].txt")

func main() {
	fmt.Printf("Loading data from %s\n", path)
	f, scanner, err := util.ReadFileToScanner(path)

	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		text := scanner.Text()
		println(text)
	}
}
