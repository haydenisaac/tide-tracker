package smtp

import (
	"gopkg.in/gomail.v2"
)

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
func (c *Client) Notify(recipient, text string) {
	m := createMessage(c.user, recipient, text)

	if err := c.dialer.DialAndSend(m); err != nil {
		panic(err)
	}
}

func createMessage(user, recipient, text string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", user)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", "Tide times")
	m.SetBody("text/html", text)

	return m
}
