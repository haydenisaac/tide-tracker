package tidetimes

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"tide-tracker/scrape"
)

var re = regexp.MustCompile(`\d{1,2}\.\d{1,2}`)

// ErrUnableToParse unable to parse text for heigh
var ErrUnableToParse = errors.New("Unable to parse text")

// GetHeight collects height from tidetimes
func GetHeight(loc string) (float64, error) {

	txt, err := scrape.FindPhrase("https://www.tidetimes.org.uk/"+loc, "Right now")
	if err != nil {
		return -1, err
	}

	height, err := parseText(txt)
	if err != nil {
		return -1, err
	}

	return height, nil
}

func parseText(txt string) (float64, error) {
	currentHeight, err := strconv.ParseFloat(re.FindString(txt), 2)
	if err != nil {
		return -1, fmt.Errorf("parse text: %w, %w", ErrUnableToParse, err)
	}
	return currentHeight, nil
}
