package util

import (
	"fmt"
	"github.com/go-gomail/gomail"
)

type EmailDialer struct {
	dialer *gomail.Dialer
}

func NewEmailDialer(smtpHost string, smtpPort int, from string, password string) *EmailDialer {
	return &EmailDialer{dialer: gomail.NewDialer(smtpHost, smtpPort, from, password)}
}

func (d *EmailDialer) SendEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", d.dialer.Username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	if err := d.dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Printf("Email to %s sent successfully!\n", to)
	return nil
}
