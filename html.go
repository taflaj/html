// html.go

// Package html contains wrapper functions for handling html.
package html

import (
	"io"
	"regexp"

	"golang.org/x/net/html"
)

// Attributes is a series of attributes inside a tag.
type Attributes map[string]string

// Tag is a structure representing tags and all attributes.
type Tag struct {
	Tag        string
	Attributes Attributes
	Node       *html.Node
}

// Constraints is a series of specific attributes we're looking for inside a tag.
type Constraints map[string]*regexp.Regexp

// Read and parse a web page.
func Read(page io.Reader) (*html.Node, error) {
	top, err := html.Parse(page)
	return top, err
}

func find(node *html.Node, tag string, constraints *Constraints, result *[]Tag) {
	if node == nil {
		return
	}
	if node.Type == html.ElementNode {
		if node.Data == tag { // this is the tag we want
			// capture all attributes
			attr := make(Attributes)
			for _, a := range node.Attr {
				attr[a.Key] = a.Val
			}
			// validate against all constraints
			valid := true
			if constraints != nil {
				for k, v := range *constraints {
					if !v.MatchString(attr[k]) {
						valid = false
					}
				}
			}
			if valid {
				*result = append(*result, Tag{tag, attr, node})
			}
		}
	}
	// continue downwards and sideways
	find(node.FirstChild, tag, constraints, result)
	find(node.NextSibling, tag, constraints, result)
}

// Find all tags starting at the given node subject to the constraints.
func Find(node *html.Node, tag string, constraints *Constraints) []Tag {
	result := []Tag{}
	find(node.FirstChild, tag, constraints, &result) // start right below the current node
	return result
}
