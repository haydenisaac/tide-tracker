package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var errEnvVariableNotSet = errors.New("Environment variable not set")
var errUnableToParseHeight = errors.New("Cannot parse trigger environment")
var errNoRecipients = errors.New("No recipients found")

func loadConfig() (config, error) {
	loc := strings.TrimSpace(os.Getenv("TIDE__LOCATION"))
	if loc == "" {
		return config{}, fmt.Errorf("location: %w", errEnvVariableNotSet)
	}
	triggerHeight := strings.TrimSpace(os.Getenv("TIDE__TRIGGER_HEIGHT"))
	if triggerHeight == "" {
		return config{}, fmt.Errorf("trigger: %w", errEnvVariableNotSet)
	}
	trigger, err := strconv.ParseFloat(triggerHeight, 2)
	if err != nil {
		return config{}, fmt.Errorf("parse trigger height: %w, %s", errUnableToParseHeight, triggerHeight)
	}
	password := os.Getenv("TIDE__PASSWORD")
	user := os.Getenv("TIDE__USER")
	if user == "" {
		return config{}, fmt.Errorf("user: %w", errEnvVariableNotSet)
	}
	tideRecipients := os.Getenv("TIDE__RECIPIENTS")
	recipients := strings.Fields(tideRecipients)
	if len(recipients) == 0 {
		return config{}, fmt.Errorf("parse recipients: %w, %v", errNoRecipients, recipients)
	}

	return config{
		trigger:    trigger,
		location:   loc,
		password:   password,
		user:       user,
		recipients: recipients,
	}, nil
}
