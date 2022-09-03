package main

import (
	"flag"
	"fmt"
	"os"
	"workspace/cyoa"
)

func main() {
	file := flag.String("file", "../../story.json", "json file with cyoa story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
