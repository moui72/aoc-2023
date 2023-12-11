package main

import (
	"fmt"
	"moui72/aoc-2023/util"
	"slices"
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
	f, scanner := util.ReadFileToScanner(path)

	defer f.Close()
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


	reverseMapOrder := make(map[string]string, len(mapOrder))
	for src, dst := range mapOrder {
		reverseMapOrder[dst] = src
	}
	seeds := make([][]string, 0, len(rawSeedsData) / 2)
	for idx := 0; idx < len(rawSeedsData); idx += 2 {
		seeds = append(seeds, rawSeedsData[idx:idx+2])
	}
	result := maxVal
	prevVal := 0
	src := ""
	diff := 0
	println("---")
	fmt.Printf("Searching from %d to %d\n", prevVal, result)
	entryDst := 0
	entrySrc := 0
	//entryRng := 0
	chunk := maxVal / 100
	fmt.Print("[")
	for lowestLoc := 0; lowestLoc < result; lowestLoc++ {
		dst := "location"
		val := lowestLoc
		for dst != "" {
			src = reverseMapOrder[dst]
			slices.Reverse(almanac[src][dst])
			for _, entry := range almanac[src][dst] {
				prevVal = val

				if val >= entry["dst"] && val < entry["dst"] + entry["range"] {
					// dst src rng
					// 1   0   69
					// -1
					entryDst = entry["dst"]
					entrySrc = entry["src"]
					//entryRng = entry["range"]
					diff = entryDst - entrySrc
					val = val - diff
					break
				}
			}
			// fmt.Printf("%s, %d -> %d: %d -> %d; ", src, entryDst, entryDst + entryRng, prevVal, val)
			entryDst = 0
			entrySrc = 0
			//entryRng = 0
			dst = src
		}
		// fmt.Printf("\nlocation %d belongs to seed %d\n", lowestLoc, val)
		for _, seedRange := range seeds {
			startSeed := parseIntOrRaise(seedRange[0])
			endSeed := startSeed + parseIntOrRaise(seedRange[1])
			// fmt.Printf("seed range %d ... %d\n", startSeed, endSeed)
			if val >= startSeed && val < endSeed {
				// fmt.Printf("found! %d: %d\n", lowestLoc, val)
				result = lowestLoc
				break
			}
		}
		if lowestLoc == result {
			break
		}
		if lowestLoc % chunk == 0 {
			fmt.Print("=", lowestLoc/chunk)
		}
	}
	fmt.Println("]")
	fmt.Printf("The lowest numbered location for any of the seeds to be planted is %d\n", result)
}
