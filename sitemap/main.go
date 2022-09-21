package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	link "workspace/parser"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "url that you want to build a sitemap for")
	flag.Parse()

	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{Scheme: reqUrl.Scheme, Host: reqUrl.Host}
	base := baseUrl.String()

	pages := hrefs(resp.Body, base)

	for _, page := range pages {
		fmt.Println(page)
	}
}

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)
	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}

	return hrefs
}
