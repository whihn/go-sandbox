package main

import (
	"github.com/whihn/go-sandbox/fileutils"
	"testing"
)

func TestShouldPrintPassedDirectory(t *testing.T) {
	// result := stringutil.Reverse("foo")
	// if result != "foo" {
	// 	t.Error("expected " + result + " to be foo")
	// }
	fileutils.Crawl("/test")
}
