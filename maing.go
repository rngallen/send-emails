package main

import (
	"fmt"
	"github.com/rngallen/send-emails/gomail"
	// "github.com/rngallen/send-emails/smtpmail"
)

func main() {

	//fmt.Println("Starting to send smtp email...")
	//goSmtp.SendStmp() // Standard go smtp library

	fmt.Println("Starting to send basic email...")
	goMail.SendGomail() // standard email using mail

	fmt.Println("Starting to send newsletter email...")
	goMail.SendNewletter() // Send newslater using mail

}
