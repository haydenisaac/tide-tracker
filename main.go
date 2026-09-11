package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"tide-tracker/scrape"
	"tide-tracker/smtp"
)

func main() {

	url := "https://www.tidetimes.org.uk/" + os.Getenv("LOCATION")
	webpage := scrape.New(url)
	allText := webpage.GetText()
	text := getCurrentTide(allText)
	if text == "" {
		log.Fatalf("no text found from %s", url)
	}
	height, err := strconv.Atoi(os.Getenv("HEIGHT_TRIGGER"))
	if err != nil {
		log.Fatalf("parsing height: %v", err)
	}
	if isHigh(float64(height), text) {
		e := smtp.New(os.Getenv("USER"), os.Getenv("PASSWORD"))
		notifyAll(e, strings.Fields(os.Getenv("RECIPIENTS")), text)
	}
}

func notifyAll(e *smtp.Client, recipients []string, text string) {
	for _, recipient := range recipients {
		fmt.Printf("notified %s\n", recipient)
		e.Notify(recipient, text)
	}
}

func getCurrentTide(allText []string) string {
	for _, text := range allText {
		if strings.Contains(text, "Right now") {
			return text
		}
	}

	return ""
}

func isHigh(high float64, text string) bool {
	re := regexp.MustCompile(`\d{1,2}\.\d{1,2}`)
	currentHeight, err := strconv.ParseFloat(re.FindString(text), 2)
	if err != nil {
		log.Fatalf("checking height. can't parse %s: %v", text, err)
	}
	return currentHeight > high
}
