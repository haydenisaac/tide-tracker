package scrape

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Errors
var (
	ErrPhraseNotFound = errors.New("phrase not found in webpage")
	ErrPhraseIsEmpty  = errors.New("phrase cannot be empty")
	ErrBadStatus      = errors.New("unexpected status")
)

// FindPhrase finds a phrase from an html webpage
func FindPhrase(url, phrase string) (string, error) {
	if phrase == "" {
		return "", fmt.Errorf("finding phrase: %w", ErrPhraseIsEmpty)
	}

	node, err := getWebpage(url)
	if err != nil {
		return "", fmt.Errorf("fetching webpage: %w", err)
	}

	textNodes := collectTextNodes(node)
	for _, text := range textNodes {
		if strings.Contains(text, phrase) {
			return text, nil
		}
	}
	return "", fmt.Errorf("finding phrase \"%s\": %w", phrase, ErrPhraseNotFound)
}

func getWebpage(url string) (*html.Node, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: bad status %d", ErrBadStatus, response.StatusCode)
	}

	node, err := html.Parse(response.Body)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// collectTextNodes collects text data from a webpage
func collectTextNodes(node *html.Node) []string {
	textData := []string{}

	for n := range node.Descendants() {
		if n.Type == html.TextNode {
			textData = append(textData, n.Data)
		}
	}
	return textData
}
