package helpers

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/gomail.v2"
)

func SendMail(body string, subject string, to ...string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", "strahinja2001@gmail.com")
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)

	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "strahinja2001@gmail.com", os.Getenv("EMAILPASS"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil

}
