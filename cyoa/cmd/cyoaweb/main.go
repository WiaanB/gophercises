package main

import (
	"encoding/json"
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

	d := json.NewDecoder(f)

	var story cyoa.Story

	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
