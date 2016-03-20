package main

import (
	"github.com/whihn/go-sandbox/fileutils"
	"testing"
)

func TestPrintsPassedDirectory(t *testing.T) {
	fileutils.Crawl("/test")
}
