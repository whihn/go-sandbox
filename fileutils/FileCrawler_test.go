package fileutils

import (
	"../fileutils"
	"testing"
)

func TestCreatesFileCrawlerWithRootDirectory(t *testing.T) {
	file, err := fileutils.Crawl("/test")
	if err != nil {
		t.Fatal("Error was thrown", err)
	}
	assertEquals(t, "/test", file.Name)
}

func assertEquals(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Error("Expected", expected, "got", actual)
	}
}
