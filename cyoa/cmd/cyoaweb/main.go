package main

import (
	"flag"
	"os"
	"workspace/cyoa"
)

func main() {
	file := flag.String("file", "../../story.json", "json file with cyoa story")
	flag.Parse()

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}
}
