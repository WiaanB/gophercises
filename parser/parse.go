package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a hypertext link in html
type Link struct {
	Href string
	Text string
}

// parse will take in a HTML doc, and return a slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, l := range nodes {
		links = append(links, buildLink(l))
	}
	return links, nil
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

func buildLink(n *html.Node) (ret Link) {
	for _, a := range n.Attr {
		if a.Key == "href" {
			ret.Href = a.Val
			break
		}
	}
	ret.Text = text(n)
	return
}

func text(n *html.Node) (ret string) {
	if n.Type == html.TextNode {
		ret = n.Data
		return n.Data
	}
	if n.Type != html.ElementNode {
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}
	ret = strings.Join(strings.Fields(ret), " ")
	return
}
