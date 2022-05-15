package main

import (
	"fmt"
	"github.com/rngallen/send-emails/smtpMail"
)

func main() {

	fmt.Println("Hello G mail")

	smtpMail.SendStmp()
}
