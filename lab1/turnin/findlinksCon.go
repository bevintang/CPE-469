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
func breadthFirst(f func(item string, channel chan<- string), worklist []string, depth int) {
	seen := make(map[string]bool)
	worklistChan := make(chan string, 200)
	for len(worklist) > 0 && depth != 0 {
		items := worklist
		fmt.Printf("Size of worklist: %d\n", len(items))
		worklist = nil

		// Child Processing)
		for _, item := range items {
			// Crawl the children
			// if !seen[item] {
			// 	seen[item] = true
			// worklist = append(worklist, f(item)...)
			go crawl(item, worklistChan)
			// }
		}
		// Give the crawl some time to load

		// TODO: -- NEED HELP HERE --
		time.Sleep(500 * time.Millisecond)

		// Append new urls to worklist
		for len(worklistChan) != 0 {
			newURL := <-worklistChan
			// fmt.Println("New URL: " + newURL)
			if !seen[newURL] {
				seen[newURL] = true
				worklist = append(worklist, newURL)
			}
		}
		depth--
	}
}

//!-breadthFirst

//!+crawl
func crawl(url string, worklistChan chan<- string) {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	for _, newURL := range list {
		worklistChan <- newURL
	}
}

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

//!-crawl

//!+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	start := time.Now()
	depth, err := testArgLength()
	testError(err)

	// Create channel to track URL queue Crawl the URL
	breadthFirst(crawl, os.Args[2:], depth+1)

	// Calculate runtime
	fmt.Println(time.Now().Sub(start))
}

//!-main
