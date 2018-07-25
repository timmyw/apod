package main

import (
	"flag"
	"github.com/timmyw/apod"
)

func main() {
	outputDir := flag.String("output", ".", "Output directory")
	flag.Parse()
	apod.ApodDownloadLatest(*outputDir)
}
