package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"tide-tracker/email"
	"tide-tracker/scrape"
)

func main() {

	url := "https://www.tidetimes.org.uk/" + os.Getenv("LOCATION")
	text, err := scrape.FindPhrase(url, "Right now")
	if err != nil {
		panic(err)
	}
	height, err := strconv.Atoi(os.Getenv("HEIGHT_TRIGGER"))
	if err != nil {
		log.Fatalf("parsing height: %v", err)
	}
	if isHigh(float64(height), text) {
		e := email.New(os.Getenv("USER"), os.Getenv("PASSWORD"))
		e.NotifyAll(strings.Fields(os.Getenv("RECIPIENTS")), text, "Tide Times")
	}
}

func isHigh(high float64, text string) bool {
	re := regexp.MustCompile(`\d{1,2}\.\d{1,2}`)
	currentHeight, err := strconv.ParseFloat(re.FindString(text), 2)
	if err != nil {
		log.Fatalf("checking height. can't parse %s: %v", text, err)
	}
	return currentHeight > high
}
