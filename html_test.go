// html_test.go
package html

import (
	"regexp"
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	page := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
	<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
	<head><title>Sample Page</title></head>
	<body>
	<img src="image.png" width="16" height="9" border="0" alt="PNG">
	<a href="#">A link</a>
	<a href="/">Home</a>
	<img src="image.jpeg" width="1024" height="640" alt="JPEG">
	<img src="image.gif" width="16" height="9" border="1" alt="GIF">
	</body></html>`
	// 1. Find the two <a>
	result := Find(strings.NewReader(page), "a", nil)
	if n := len(result); n != 2 {
		t.Errorf("Found %v <a> but expected 2.", n)
	}
	// 2. Find one <img> with border=0
	border0, _ := regexp.Compile("0")
	constraints := Constraints{"border": border0}
	result = Find(strings.NewReader(page), "img", &constraints)
	if n := len(result); n != 1 {
		t.Errorf("Found %v <img> with border=0 but expected 1.", n)
	}
	// 3. Find one JPEG
	jpeg, _ := regexp.Compile(".*jp.?g")
	constraints = Constraints{"src": jpeg}
	result = Find(strings.NewReader(page), "img", &constraints)
	if n := len(result); n != 1 {
		t.Errorf("Found %v JPEG but expected 1.", n)
	}
	// 4. Find the three <img>
	result = Find(strings.NewReader(page), "img", nil)
	if n := len(result); n != 3 {
		t.Errorf("Found %v <img> but expected 3.", n)
	}
}
