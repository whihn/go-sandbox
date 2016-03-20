package main

import (
	"fmt"
	"os"
	"github.com/whihn/go-sandbox/fileutils"
)

func main() {
	args := os.Args[1:]
	crawlDir := args[0]
	fmt.Println("Crawling directory " + crawlDir)
	fileutils.Crawl(crawlDir)
}
