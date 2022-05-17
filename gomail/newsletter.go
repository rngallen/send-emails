package goMail

import (
	"fmt"
	"log"

	gomail "github.com/go-mail/mail"
	"github.com/rngallen/send-emails/config"
)

type UserAddress struct {
	Name    string
	Address string
}

var toUsers = []UserAddress{
	{Name: "User One", Address: "user1@gmail.com"},
	{Name: "User Two", Address: "user2@gmail.com"},
	{Name: "User Three", Address: "user3@gmail.com"},
	{Name: "User Four", Address: "user4@gmail.com"},
	{Name: "User Five", Address: "user5@gmail.com"},
	{Name: "User Six", Address: "user6@gmail.com"},
	{Name: "User Seven", Address: "user7@gmail.com"},
}

func SendNewletter() {

	host := config.Config("EMAIL_HOST.SERVER")
	port := config.Intconfig("EMAIL_HOST.PORT")
	username := config.Config("EMAIL_HOST.USER")
	password := config.Config("EMAIL_HOST.PASSWORD")

	d := gomail.NewDialer(host, port, username, password)

	d.StartTLSPolicy = gomail.MandatoryStartTLS

	sender, err := d.Dial()
	if err != nil {
		log.Panicln(err)
	}

	msg := gomail.NewMessage()

	for _, user := range toUsers {
		msg.SetHeader("From", config.Config("EMAIL_HOST.USER"))
		msg.SetAddressHeader("To", user.Address, user.Name)
		msg.SetHeader("Subject", config.Config("EMAIL.SUBJECT"))

		msg.SetBody("text/html", fmt.Sprintf("Hello %s !", user.Name))

		if err := gomail.Send(sender, msg); err != nil {
			log.Printf("Cound not send email to %q: %v", user.Address, err)

		}
		log.Printf("Email sent to %q successfully", user.Name)
		msg.Reset()
	}

}
