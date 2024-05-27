package main_test

import (
	File "github.com/Meez25/AOC/01"
	"testing"
	"testing/fstest"
)

func TestFile(t *testing.T) {
	fs := fstest.MapFS{
		"input.txt": {Data: []byte(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)},
	}

	decodedFile, err := File.NewFileFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(decodedFile) != len(fs) {
		t.Errorf("got %d lines, wanted %d lines", len(decodedFile), len(fs))
	}

	got := File{decodedFile[0]}
}
