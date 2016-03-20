package fileutils

import (
	"../fileutils"
	"testing"
	"fmt"
)

func TestCreatesFileCrawlerWithRootDirectory(t *testing.T) {
	var files fileutils.Directory = fileutils.Crawl("/test")
	fmt.Println("Crawling returned: " + files.Name)
}
