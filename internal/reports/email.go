package reports

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)


func SendEmail (to, subject, content string) error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	host := "smtp.gmail.com"
	port := "587"
	username := os.Getenv("smtp_email")
	password := os.Getenv("smtp_password")
	
	auth := smtp.PlainAuth("", username, password, host)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", host, port),
		auth,
		username,
		[]string{to},
		[]byte(fmt.Sprintf("Subject: %s\r\n%s", subject, content)),
	)

	return err
}
