package goMail

import (
	// "crypto/tls"
	"fmt"
	"log"
	"strings"

	gomail "github.com/go-mail/mail"
	"github.com/rngallen/send-emails/config"
)

func SendGomail() {
	// initiate new mail
	m := gomail.NewMessage()
	user := config.Config("EMAIL_HOST.USER")
	pwd := config.Config("EMAIL_HOST.PASSWORD")
	host := config.Config("EMAIL_HOST.SERVER")
	port := config.Intconfig("EMAIL_HOST.PORT")
	body := config.Config("EMAIL.BODY")
	boydType := config.Config("EMAIL.BODY_TYPE")

	type Headers map[string][]string

	headers := make(Headers, 0)
	headers["From"] = []string{user}
	headers["To"] = strings.Split(config.Config("EMAIL.TO"), ",")
	headers["Subject"] = []string{config.Config("EMAIL.SUBJECT")}

	m.SetHeaders(headers)
	m.SetAddressHeader("Cc", config.Config("EMAIL.CC"), "")
	m.SetAddressHeader("Bcc", config.Config("EMAIL.BCC"), "")
	m.SetBody(boydType, body)
	m.Attach(config.Config("EMAIL.ATTACHMENT"), gomail.Rename("sample.pdf"))
	fmt.Println(config.Config("EMAIL.ATTACHMENT"))

	// Smtp Server
	d := gomail.NewDialer(host, port, user, pwd)

	d.StartTLSPolicy = gomail.MandatoryStartTLS

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		log.Panicln(err)

	} else {
		log.Println("Gomail Basic email sentsuccessfully")
	}

}
