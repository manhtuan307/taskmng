package utils

import (
	"log"
	"net/smtp"
)

var smtpHost string
var smtpPort int
var senderEmail string
var senderPassword string

//InitMailSettings - init mail settings
func InitMailSettings() {
	smtpHost = "smtp.gmail.com"
	smtpPort = 587
	senderEmail = "minhnha30@gmail.com"
	senderPassword = "TenLao307"
}

//SendMail - send mail to list of recipients
func SendMail(recipientEmails []string, bodyContent string) {
	// Set up authentication information.
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(smtpHost+":"+string(smtpPort),
		auth, senderEmail, recipientEmails, []byte(bodyContent),
	)
	if err != nil {
		log.Fatal(err)
	}
}

//SendMailToOne - send mail to only one recipient
func SendMailToOne(recipientEmail string, bodyContent string) {
	SendMail([]string{recipientEmail}, bodyContent)
}
