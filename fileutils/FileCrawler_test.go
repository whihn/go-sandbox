package fileutils

import (
	"../fileutils"
	"testing"
)

func TestCreatesFileCrawlerWithRootDirectory(t *testing.T) {
	fileutils.Crawl("/test")
}
