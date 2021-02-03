package cmd

import (
	"fmt"
	"github.com/fprotimaru/test-project/models"
	"github.com/urfave/cli/v2"
	"net/smtp"
)

var SendEmail = cli.Command{
	Name:   "sendemail",
	Usage:  "send email to user",
	Action: sendEmail,
}

func sendEmail(c *cli.Context) error {
	for _, user := range models.Users {
		sendEMAIL(&user)
	}
	return nil
}

func sendEMAIL(user *models.User) {
	// Sender data.
	from := "fprotimaru@gmail.com"
	password := "*************"

	// Receiver email address.
	to := []string{
		"fprotimaru@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
