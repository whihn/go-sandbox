package main

import (
	"fmt"
	"github.com/whihn/go-sandbox/fileutils"
)

func main() {
	fmt.Println("starting crawler...")
	fileutils.Crawl("/Users/werner/")
}
