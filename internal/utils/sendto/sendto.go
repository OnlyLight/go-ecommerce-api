package sendto

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

const (
	SMTPHost = "smtp.gmail.com"
	SMTPPort = "587"
	SMTPUser = "username"
	SMTPPass = "password"
)

type EmailAdress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAdress
	To      []string
	Subject string
	Body    string
}

func BuildMessage(mail Mail) string {
	header := "Mine-Version: 1.0\nContent-Type: text/plain; charset=\"utf-8\"\n"
	return fmt.Sprintf("%s From: %s\nTo: %s\nSubject: %s\n\n%s", header, mail.From.Address, mail.To, mail.Subject, mail.Body)
}

func SendTemplateEmailOtp(to []string, from string, templateName string, templateData map[string]interface{}) error {
	htmlBody, err := getMailTemplate(templateName, templateData)
	if err != nil {
		return err
	}

	return send(to, from, htmlBody)
}

func getMailTemplate(templateName string, templateData map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	t := template.Must(template.New(templateName).ParseFiles("template-email/" + templateName))
	err := t.Execute(htmlTemplate, templateData)
	if err != nil {
		return "", err
	}

	return htmlTemplate.String(), nil
}

func SendTextEmailOtp(to []string, from string, otp string) error {
	// Send text message
	contentEmail := Mail{
		From:    EmailAdress{Address: from, Name: "Go Ecommerce"},
		To:      to,
		Subject: "OTP for Go Ecommerce",
		Body:    fmt.Sprintf("Your OTP is $s", otp),
	}

	messageMail := BuildMessage(contentEmail)

	auth := smtp.PlainAuth("", SMTPUser, SMTPPass, SMTPHost)
	err := smtp.SendMail(SMTPHost+":"+SMTPPort, auth, SMTPUser, to, []byte(messageMail))
	if err != nil {
		global.Logger.Error("Error sending email", zap.Error(err))
		return err
	}

	return nil
}

func send(to []string, from string, htmlBody string) error {
	// Send text message
	contentEmail := Mail{
		From:    EmailAdress{Address: from, Name: "Go Ecommerce"},
		To:      to,
		Subject: "OTP for Go Ecommerce",
		Body:    htmlBody,
	}

	messageMail := BuildMessage(contentEmail)

	auth := smtp.PlainAuth("", SMTPUser, SMTPPass, SMTPHost)
	err := smtp.SendMail(SMTPHost+":"+SMTPPort, auth, SMTPUser, to, []byte(messageMail))
	if err != nil {
		global.Logger.Error("Error sending email", zap.Error(err))
		return err
	}

	return nil
}
