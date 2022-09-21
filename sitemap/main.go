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

	pages := get(*urlFlag)

	for _, page := range pages {
		fmt.Println(page)
	}
}

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)
	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}

	return ret
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{Scheme: reqUrl.Scheme, Host: reqUrl.Host}
	base := baseUrl.String()

	pages := hrefs(resp.Body, base)

	return pages
}
