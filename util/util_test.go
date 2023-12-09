package util

import (
	"testing"
)

func TestReadFileToScanner(t *testing.T) {
	expectedValues := [2]string{"line one", "line two"}
	f, scanner, err := ReadFileToScanner(util.RelativePathTo("inputs/d3.txt"))
	defer f.Close()
	if err != nil {
		t.Fatalf(`ReadFileToScanner("inputs/test.txt") produced error %v`, err)
	}
	for _, expectedLine := range expectedValues {

		scanner.Scan()
		line := scanner.Text()
		err := scanner.Err();

		if line != expectedLine || err != nil {
			t.Fatalf(`Line 1 = %q, %v, want %q`, line, err, expectedLine)
		}
	}
}
