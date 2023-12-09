package main

import (
	"fmt"
	"log"
	"moui72/aoc-2023/util"
	"strconv"
	"strings"
)

var path string = util.RelativePathTo("inputs/d5.txt")
var max_loops int = 99

func parseIntOrRaise(input string) int {
	str, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return str
}

func parseAlmanacEntry(input string) (int, int, int) {
	values := strings.Split(input, " ")
	if len(values) != 3 {
		panic("Wrong number of almanac entries")
	}
	return parseIntOrRaise(values[0]), parseIntOrRaise(values[1]), parseIntOrRaise(values[2])

}

func main() {
	fmt.Printf("Loading data from %s\n", path)
	f, scanner, err := util.ReadFileToScanner(path)

	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	building := ""
	almanac := make(map[string]map[string][]map[string]int)
	mapOrder := map[string]string{}
	source := ""
	dest := ""
	maxVal := 0

	scanner.Scan()
	rawSeedsData := strings.Split(scanner.Text(), " ")[1:]
	prevLineBlank := false

	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasSuffix(text, ":") {
			building, _ = strings.CutSuffix(text, " map:")
			source, dest, _ = strings.Cut(building, "-to-")
			mapOrder[source] = dest
			fmt.Printf("building %s to %s\n", source, dest)
		} else if text != ""{
			if source == "" {
				panic("what almanac????")
			}
			destStart, sourceStart, rangeLength := parseAlmanacEntry(text)
			maxVal = max(destStart, sourceStart, maxVal)
			entry := map[string]int{"dst": destStart, "src": sourceStart, "range": rangeLength}
			if len(almanac[source]) <= 0 {
				almanac[source] = make(map[string][]map[string]int)
			}
			almanac[source][dest] = append(almanac[source][dest], entry)
		} else {
			if !prevLineBlank {
				prevLineBlank = true
				continue
			}
			fmt.Printf("done building %s to %s\n", source, dest)
			building = ""
			source = ""
			dest = ""
			continue
		}
		prevLineBlank = false
	}
	seedToLoc := map[int]int{}
	lowestLoc := maxVal

	seeds := make([]int, len(rawSeedsData))
	for idx, val := range rawSeedsData {
		numVal := parseIntOrRaise(val)
		if idx % 2 == 0 {
			seeds = append(seeds, numVal)
		} else {
			seeds = append(seeds, seeds[idx] + numVal)

		}
	}
	for _, seed := range seeds {
		next := "seed"
		val := seed
		fmt.Printf("Locating seed %s ...\n", seed)
		
		for loops := 0; loops < max_loops; loops++ {
			src := next
			next = mapOrder[next]
			if next == "" {
				continue
			}
			dst := next
			for _, entry := range almanac[src][dst] {
				if val >= entry["src"] && val < entry["src"] + entry["range"] {
					fmt.Printf("\n  - %d > %d >= %d\n", entry["src"] + entry["range"], val, entry["src"])
					val = val - (entry["src"] - entry["dst"])
					break
				}
			}
			fmt.Printf("    - Seed %s needs %s %d\n", seed, next, val)
		}
		seedToLoc[seed] = val
		if val < lowestLoc {
			lowestLoc = val
		}
	}
	fmt.Printf("The lowest numbered location for any of the seeds to be planted is %d\n", lowestLoc)
}
