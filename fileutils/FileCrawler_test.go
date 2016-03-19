package fileutils

import (
	"testing"
	"../fileutils"
)

func TestCreatesFileCrawlerWithRootDirectory(t *testing.T) {
	fileutils.Crawl("/test")
}