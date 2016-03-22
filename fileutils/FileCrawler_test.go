package fileutils

import (
	"../fileutils"
	"testing"
	"os"
)

func TestCreatesFileCrawlerWithRootDirectory(t *testing.T) {
	fileName := "../testdata"
	file, err := fileutils.Crawl(fileName)
	if err != nil {
		t.Fatal("Error was thrown", err)
	}
	assertEquals(t, fileName, file.Name)
	assertEquals(t, 1, len(file.FileInfos))
}

func TestReadDir(t *testing.T) {
	file, err := os.Open("../testdata")
	fileInfos, err := file.Readdir(0)
	if err != nil {
		t.Fatal(err)
	}
	assertEquals(t, 1, len(fileInfos))
}


func assertEquals(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Error("Expected", expected, "got", actual)
	}
}
