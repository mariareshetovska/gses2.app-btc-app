package services

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strconv"
	"text/template"

	"gses2.app-btc/config"
)

func SendBitcoinRateEmail(subject string, templatePath string, to []string) error {
	t, err := parseEmailTemplate(templatePath)
	if err != nil {
		return fmt.Errorf("failed to parse email template: %v", err)
	}

	rate, err := GetBitcoinRate()
	if err != nil {
		return fmt.Errorf("failed to get Bitcoin rate: %v", err)
	}

	body, err := renderEmailBody(t, formatRate(rate))
	if err != nil {
		return fmt.Errorf("failed to render email body: %v", err)
	}

	cnf, err := config.LoadEmailConfig()
	if err != nil {
		return fmt.Errorf("error loading email config: %v", err)
	}

	err = sendEmail(subject, body.String(), to, cnf)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

func parseEmailTemplate(templatePath string) (*template.Template, error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func formatRate(rate float64) string {
	return strconv.FormatFloat(rate, 'f', 2, 64)
}

func renderEmailBody(t *template.Template, rateString string) (*bytes.Buffer, error) {
	var body bytes.Buffer
	err := t.Execute(&body, struct{ Rate string }{Rate: rateString})
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func sendEmail(subject, body string, to []string, cnf config.ConfigEmail) error {
	auth := smtp.PlainAuth("", cnf.Username, cnf.Password, "smtp.gmail.com")
	headers := "MIME-version: 1.0;\nContent-Type:text/html;charset=\"UTF-8\";"
	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body

	err := smtp.SendMail("smtp.gmail.com:587", auth, cnf.Username, to, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}
