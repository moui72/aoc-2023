package main

import (
	"fmt"
	"log"
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
	
	for scanner.Scan() {
		text := scanner.Text()
		println(text)
	}
}
