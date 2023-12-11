package util

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

func ReadFileToScanner(path string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(path)
	if err != nil {
			panic(err)
	}
	scanner := bufio.NewScanner(file)
	return file, scanner
}

func LineCount(path string) int {
	f, scanner := ReadFileToScanner(path)
	defer f.Close()
	n := 0
	for scanner.Scan() {
		n += 1
	}
	return n
}

func GetCurrentFilePath() string {
	pwd, _ := os.Getwd()
	return pwd
}

func RelativePathTo(path string) string {
	return filepath.Join(GetCurrentFilePath(), path)
}

func PathFromFileName(filename string) string {
	if ! strings.HasSuffix(filename, ".txt") {
		filename += ".txt"
	}
	return filepath.Join(GetCurrentFilePath(), "inputs", filename)
}


func ParseIntOrRaise(input string) int {
	str, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return str
}
