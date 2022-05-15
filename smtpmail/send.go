package smtpMail

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

	// smtp server configuration
	smtpHost := config.Config("EMAIL_HOST.SERVER")
	smtpPort := config.Config("EMAIL_HOST.PORT")

	msg := config.Config("EMAIL.MESSAGE")

	// Authenticate
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending Email
	message := []byte(msg)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Email have been sent successfully")

}
