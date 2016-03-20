package fileutils

import (
	"fmt"
	"time"
)

func Crawl(dir string) Directory {
	fmt.Println(dir)
	file := Directory{Name: "/test"}
	return file
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