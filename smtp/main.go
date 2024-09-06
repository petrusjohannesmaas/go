package main

import (
	"fmt"
	"net/smtp"
)

func sendMailSimple() {
	auth := smtp.PlainAuth(
		"",
		"webadmin@megs.network",
		"g7QuY7bI-fGn",
		"mail.megs.network",
	)

	msg := "Subject: My special subject\nThis is the body of my email"

	err := smtp.SendMail(
		"mail.megs.network:465",
		auth,
		"support@megs.co.za",
		[]string{"petrusjohannesmaas@outlook.com"},
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	sendMailSimple()
}
