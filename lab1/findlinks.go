// Copyright © 2016 Thw Go Programming Language
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"links"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string, depth int) {
	seen := make(map[string]bool)
	for len(worklist) > 0 && depth != 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
		depth--
	}
}

//!-breadthFirst

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

func testArgLength() (int, error) {
	if len(os.Args) != 3 {
		printUsageAndExit()
	}
	return strconv.Atoi(string(os.Args[1]))
}

func testError(err error) {
	if err != nil {
		fmt.Printf("An error occured: %e", err)
	}
}

func printUsageAndExit() {
	fmt.Printf("Usage: go run findlinksCon.go [DEPTH] [URL]\n")
	os.Exit(1)
}

//!+main
func main() {
	start := time.Now()
	depth, err := testArgLength()
	testError(err)
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[2:], depth+1)
	fmt.Println(time.Now().Sub(start))
}

//!-main
