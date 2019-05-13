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
}

// Constraints is a series of specific attributes we're looking for inside a tag.
type Constraints map[string]*regexp.Regexp

// Find all tags on the given page subject to the constraints.
func Find(page io.Reader, tag string, constraints *Constraints) []Tag {
	result := []Tag{}
	tokenizer := html.NewTokenizer(page)
	for looping := true; looping == true; {
		token := tokenizer.Next()
		switch token {
		case html.ErrorToken:
			// end of document
			looping = false
		case html.StartTagToken, html.SelfClosingTagToken:
			// is this the tag we want?
			if t := tokenizer.Token(); t.Data == tag {
				// capture all attributes
				attr := make(Attributes)
				for _, a := range t.Attr {
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
					result = append(result, Tag{tag, attr})
				}
			}
		}
	}
	return result
}
