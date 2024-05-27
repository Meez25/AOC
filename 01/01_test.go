package main

import (
	"testing"
	"testing/fstest"
)

func TestSolution(t *testing.T) {
	fs := fstest.MapFS{
		"input.txt": {Data: []byte(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)},
	}

	decodedFile := file.NewInputFromFS(fs)

	if len(decodedFile) != len(fs) {
		t.Errorf("got %d lines, wanted %d lines", len(decodedFile), len(fs))
	}
}
