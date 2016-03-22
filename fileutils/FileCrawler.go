package fileutils

import (
	"fmt"
	"time"
	"errors"
	"os"
)

func Crawl(path string) (Directory, error) {
	file, err := os.Open(path)
	if err != nil {
		return Directory{}, errors.New("invalid path")
	}

	fileInfos, err := file.Readdir(0)
	dir := Directory{ Name: path, FileInfos: fileInfos }
	return dir, err
}

type File interface {

	timestamp() time.Time
	toString() string
}

type Directory struct {
	Name string
	FileInfos []os.FileInfo
}

func timestamp(file Directory) {
	fmt.Println("timestamp called for file: " + toString(file))
}

func toString(file Directory) string {
	return file.Name
}

// func (file *File) Timestamp{

// }