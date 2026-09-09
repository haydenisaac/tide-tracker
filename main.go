package main

import (
	"encoding/json"
	"os"
	"tide-tracker/smtp"
)

type config struct {
	User       string
	Password   string
	Recipients []string
}

func main() {
	conf := loadConfig()

	randomText := "this is more test text"

	e := smtp.New(conf.User, conf.Password)
	notifyAll(e, conf.Recipients, randomText)
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
		e.Notify(recipient, text)
	}
}
