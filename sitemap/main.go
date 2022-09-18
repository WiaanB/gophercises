package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "url that you want to build a sitemap for")
	flag.Parse()

	fmt.Println(*urlFlag)

	// 1. GET webpage html

	// 2. How to parse all the links on the page

	// 3. Build proper urls with links (cleanup links)

	// 4. filter out links to other domains

	// 5. find all the pages (BFS)

	// 6. print out xml
}
