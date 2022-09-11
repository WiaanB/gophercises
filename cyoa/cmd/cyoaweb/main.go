package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"workspace/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "port to start server on")
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

	h := cyoa.NewHandler(story, cyoa.WithTemplate(nil))
	fmt.Printf("starting the server on port :%d\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), h)
}
