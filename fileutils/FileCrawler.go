package fileutils

import (
	"fmt"
	"time"
	"errors"
)

func Crawl(path string) (Directory, error) {
	if path == "/wrongPath" {
		return Directory{}, errors.New("invalid path")
	}
	fmt.Println(path)
	file := Directory{ Name: path }
	return file, nil
}

type File interface {

	timestamp() time.Time
	toString() string
}

type Directory struct {
	Name string
}

func timestamp(file Directory) {
	fmt.Println("timestamp called for file: " + toString(file))
}

func toString(file Directory) string {
	return file.Name
}

// func (file *File) Timestamp{

// }