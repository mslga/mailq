package mail

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

// SendMail sends an HTML email to one or more recipients.
//
// If password is empty, SMTP authentication is skipped.
// The sender address is also used as the SMTP username when authentication is enabled.
// Set insecureSkipVerify=true only for trusted internal SMTP servers.
func SendMail(host string, port int, from string, password string, insecureSkipVerify bool, recipients []string, subject string, body string) error {

	dialer := gomail.NewDialer(host, port, from, password)
	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
	}

	sender, err := dialer.Dial()
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server %s:%d: %w", host, port, err)
	}
	defer sender.Close()

	for _, recipient := range recipients {
		message := gomail.NewMessage()

		message.SetHeader("From", from)
		message.SetHeader("To", recipient)
		message.SetHeader("Subject", subject)
		message.SetBody("text/html", body)

		if err := gomail.Send(sender, message); err != nil {
			return fmt.Errorf("failed to send email to %s: %w", recipient, err)
		}
	}

	return nil
}
