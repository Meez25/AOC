package main

import (
	"io/fs"
)

type File struct {
	content []byte
}

func NewFileFromFS(fileSystem fs.FS) ([]File, error) {
	dir, _ := fs.ReadDir(fileSystem, ".")
	var files []File
	for range dir {
		files = append(files, File{})
	}
	return files, nil
}

func main() {
}
