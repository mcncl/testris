package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mcncl/testris/internal/finder"
	"github.com/mcncl/testris/internal/splitter"
)

func main() {
	index := flag.Int("index", 0, "Current parallel index (0-based)")
	total := flag.Int("total", 1, "Total number of parallel runs")
	dir := flag.String("dir", ".", "Root directory to scan for tests")
	flag.Parse()

	if err := validateFlags(*index, *total); err != nil {
		log.Fatal(err)
	}

	tests, err := finder.FindTests(*dir)
	if err != nil {
		log.Fatal(err)
	}

	pattern := splitter.GenerateTestPattern(tests, *index, *total)
	fmt.Print(pattern)
}

func validateFlags(index, total int) error {
	if total < 1 {
		return fmt.Errorf("total must be >= 1")
	}
	if index < 0 || index >= total {
		return fmt.Errorf("index must be between 0 and %d", total-1)
	}
	return nil
}
