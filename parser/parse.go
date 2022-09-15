package link

import "io"

// Link represents a hypertext link in html
type Link struct {
	Href string
	Text string
}

// parse will take in a HTML doc, and return a slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
