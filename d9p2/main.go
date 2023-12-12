package main

import (
	"fmt"
	"moui72/aoc-2023/util"
	"os"
	"slices"
	"strings"
)

func lastItem[T any](arr []T) T {
	return arr[len(arr)-1]
}

func buildDiffs(history []int) [][]int {
	allZero:= false
	currentSeries := make([]int, len(history))
	copy(currentSeries, history)
	series := [][]int{currentSeries}
	for allZero != true {
		nextSeries := []int{}
		allZero = true
		for i := 1; i < len(currentSeries); i++ {
			diff := currentSeries[i] - currentSeries[i-1]
			nextSeries = append(nextSeries, diff)
			if diff != 0 {
				allZero = false
			}
		}
		series = append(series, nextSeries)
		currentSeries = nextSeries
	}
	slices.Reverse(series)
	return series
}

func predictPrevious(history []int) int {
	series := buildDiffs(history)
	series[0] = append(series[0], 0)
	for _, s := range series {
		slices.Reverse(s)
	}
	nextValue := 0
	for i := 1; i < len(series); i++ {
		s := series[i]
		prevS := series[i-1]
		nextValue = lastItem(s) - lastItem(prevS)
		series[i] = append(series[i], nextValue)
	}
	return nextValue
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
	sum := 0
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		if len(text) > 0 {
			history := []int{}
			for _, value := range text {
				history = append(history, util.ParseIntOrRaise(value))
			}
			prev := predictPrevious(history)
			fmt.Printf("%d <- ", prev)
			println(strings.Join(text, ", "))
			sum += prev
		}
	}
	fmt.Printf("The sum of predicted values is %d\n", sum)
}
