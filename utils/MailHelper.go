package utils

import (
	"github.com/kataras/go-mailer"
)

var config mailer.Config

//InitMailSettings - init mail settings
func InitMailSettings() {
	config = mailer.Config{
		Host:       "smtp.gmail.com",
		Username:   "appchecker1988",
		Password:   "TryIt307",
		FromAddr:   "appchecker1988@gmail.com",
		Port:       587,
		UseCommand: false,
	}

}

//SendMail - send mail to list of recipients
func SendMail(recipientEmails []string, subject string, bodyContent string) {
	sender := mailer.New(config)
	err := sender.Send(subject, bodyContent, recipientEmails...)
	if err != nil {
		println("error while sending the e-mail: " + err.Error())
	}
}

//SendMailToOne - send mail to only one recipient
func SendMailToOne(recipientEmail string, subject string, bodyContent string) {
	SendMail([]string{recipientEmail}, subject, bodyContent)
}
