package fileutils

import (
	"../fileutils"
	"testing"
)

func TestCreatesFileCrawlerWithRootDirectory(t *testing.T) {
	var file fileutils.Directory = fileutils.Crawl("/test")
	assertEquals(t, "/test", file.Name)
}

func assertEquals(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Error("Expected", expected, "got", actual)
	}
}
