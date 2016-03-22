package fileutils

import (
	"../fileutils"
	"testing"
	"os"
	"time"
)

func TestCreatesFileCrawlerWithRootDirectory(t *testing.T) {
	fileName := "../testdata"
	file, err := fileutils.Crawl(fileName)
	if err != nil {
		t.Fatal("crawling directory", fileName, "failed: ", err)
	}
	assertEquals(t, fileName, file.Name)
	assertEquals(t, 1, len(file.FileInfos))
	assertTimestamp(t, file.FileInfos[0], 2016, time.March, 20, 12, 38)
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

func assertTimestamp(t *testing.T, fileInfo os.FileInfo, year int, month time.Month, day int, hour int, minute int) {
	time := fileInfo.ModTime()
	assertEquals(t, year, time.Year())
	assertEquals(t, month, time.Month())
	assertEquals(t, day, time.Day())
	assertEquals(t, hour, time.Hour())
	assertEquals(t, minute, time.Minute())
}