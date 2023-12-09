package util

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

func ReadFileToScanner(path string) (*os.File, *bufio.Scanner, error) {
	file, err := os.Open(path)
	if err != nil {
			return file, nil, err
	}
	scanner := bufio.NewScanner(file)
	return file, scanner, nil
}

func LineCount(path string) int {
	f, scanner, err := ReadFileToScanner(path)
	defer f.Close()
	if err != nil {
			log.Fatal(err)
	}
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
