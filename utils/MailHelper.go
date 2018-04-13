package utils

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
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

//SendMailSSL - send mail which host has SSL setting
func SendMailSSL(recipientEmail string, subject string, bodyContent string) {

	from := mail.Address{"", senderEmail}
	to := mail.Address{"", recipientEmail}
	subj := subject
	body := bodyContent

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := smtpHost + ":" + string(smtpPort)

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", senderEmail, senderPassword, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()

}
