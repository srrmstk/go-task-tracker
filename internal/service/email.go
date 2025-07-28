package service

import (
	"os"
	"strconv"

	"gopkg.in/mail.v2"
)

type EmailService struct{}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (s *EmailService) SendEmail(to string, addressHeader string, subject string, body string) error {
	from := os.Getenv("SMTP_FROM")
	password := os.Getenv("SMTP_PASS")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return err
	}

	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetAddressHeader("From", from, addressHeader)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := mail.NewDialer(smtpHost, smtpPort, from, password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
