package scrape

import (
	"net/http"

	"golang.org/x/net/html"
)

// Webpage information
type Webpage struct {
	node *html.Node
}

// New stores the webpage information
func New(url string) *Webpage {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	node, err := html.Parse(response.Body)
	if err != nil {
		panic(err)
	}
	return &Webpage{node: node}
}

// GetText collects text data from a webpage
func (w *Webpage) GetText() []string {
	textData := []string{}

	for n := range w.node.Descendants() {
		if n.Type == html.TextNode {
			textData = append(textData, n.Data)
		}
	}
	return textData
}
