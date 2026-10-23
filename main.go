package main

import (
	"fmt"
	"log"
	"tide-tracker/email"
	"tide-tracker/tidetimes"

	"log/slog"
)

type config struct {
	location   string
	trigger    float64
	password   string
	user       string
	recipients []string
}

func main() {

	c, err := loadConfig()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	err = run(c)

	if err != nil {
		log.Fatalf("run: %v", err)
	}
}

func run(c config) error {
	predictedHeight, err := tidetimes.GetHeight(c.location)
	if err != nil {
		return fmt.Errorf("get predicted tide height: %w", err)
	}

	if c.trigger < predictedHeight {
		message := predictedHeightMessage(predictedHeight)
		slog.Info("height is larger than the trigger - sending email.", "message", message, "recipients", len(c.recipients))
		e := email.New(c.user, c.password)
		e.NotifyAll(c.recipients, message, "Possible Overtop")
	}
	return nil
}
