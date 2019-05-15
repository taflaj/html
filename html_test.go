// html_test.go
package html

import (
	"regexp"
	"strings"
	"testing"
)

const page = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
<head><title>Sample Page</title></head>
<body><!-- insert comment here -->
<h1>A Page</h1>
<div>Some text here.</div>
<img src="image.png" width="16" height="9" border="0" alt="PNG">
<p><a href="#">A link</a></p>
<div><a href="/">Home</a></div>
<div id="internal" class="enclosed">
Some more text here.<br />
<img src="image.jpeg" width="1024" height="640" alt="JPEG" />
<div><img src="image.gif" width="16" height="9" border="1" alt="GIF"></div>
</div>
<img src="picture.jpg" alt="Picture">
</body>
</html>`

func TestFind(t *testing.T) {
	data, err := Read(strings.NewReader(page))
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	// 1. Find the two <a>
	result := Find(data, "a", nil)
	if n := len(result); n != 2 {
		t.Errorf("Found %v <a> but expected 2.", n)
	}
	// 2. Find one <div> with id="internal"
	internal, _ := regexp.Compile("internal")
	constraints := Constraints{"id": internal}
	result = Find(data, "div", &constraints)
	if n := len(result); n != 1 {
		t.Errorf("Found %v <div> with id=\"internal\" but expected 1.", n)
	}
	// 3. Find two internal <img>
	result = Find(result[0].Node, "img", nil)
	if n := len(result); n != 2 {
		t.Errorf("Found %v internal <img> but expected 2.", n)
	}
	// 4. Find the four <img>
	result = Find(data, "img", nil)
	if n := len(result); n != 4 {
		t.Errorf("Found %v <img> but expected 3.", n)
	}
}
