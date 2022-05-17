package goSmtp

import (
	"log"
	"net/smtp"

	"github.com/rngallen/send-emails/config"
)

func SendStmp() {

	// Sender data
	from := config.Config("EMAIL_HOST.USER")
	password := config.Config("EMAIL_HOST.PASSWORD")

	// Receiver email addres
	toEmails := config.Config("EMAIL.TO")
	to := []string{toEmails}

	// Email Body
	msg := config.Config("EMAIL.BODY")
	message := []byte(msg)

	// smtp server configuration
	smtpHost := config.Config("EMAIL_HOST.SERVER")
	smtpPort := config.Config("EMAIL_HOST.PORT")

	// Authenticate
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending Email

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Println(err)
		// return
	} else {
		log.Printf("Email have been sent successfully")
	}

}
