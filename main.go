package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"tide-tracker/scrape"
	"tide-tracker/smtp"
)

type config struct {
	User       string
	Password   string
	Recipients []string
	Location   string
	MinHeight  float64
}

func main() {
	config := loadConfig()

	webpage := scrape.New("https://www.tidetimes.org.uk/" + config.Location)
	allText := webpage.GetText()
	text := getCurrentTide(allText)

	if isHigh(config.MinHeight, text) {
		e := smtp.New(config.User, config.Password)
		notifyAll(e, config.Recipients, text)
	}

}

func loadConfig() config {
	file, err := os.Open("./config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	conf := config{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}

	return conf
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
	re := regexp.MustCompile(`\d{2}\.\d{2}`)

	currentHeight, err := strconv.ParseFloat(re.FindString(text), 2)
	if err != nil {
		panic(err)
	}
	return currentHeight > high
}
