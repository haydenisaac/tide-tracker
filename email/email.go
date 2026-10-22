package email

import (
	"errors"
	"fmt"

	"gopkg.in/gomail.v2"
)

// ErrFailedToNotify if the notify fails
var ErrFailedToNotify = errors.New("failed to notify a client")

// Client to send emails from
type Client struct {
	user   string
	dialer *gomail.Dialer
}

// New initialises a new email client
func New(user, password string) *Client {
	return &Client{user: user, dialer: gomail.NewDialer("smtp.gmail.com", 587, user, password)}
}

// Notify emails the recipient the text message
func (c *Client) Notify(recipient, text, subject string) error {
	m := createMessage(c.user, recipient, text, subject)

	if err := c.dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("notify %s: %w with %w", recipient, ErrFailedToNotify, err)
	}
	return nil
}

// NotifyAll sends an email to all recipients
func (c *Client) NotifyAll(recipients []string, text, subject string) error {
	errs := []error{}
	for _, recipient := range recipients {
		err := c.Notify(recipient, text, subject)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func createMessage(user, recipient, text, subject string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", user)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", text)

	return m
}
