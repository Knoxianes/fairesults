package helpers

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/gomail.v2"
)

func SendVerificationMailEmail(verificationToken string, to string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SENDEREMAIL"))
	m.SetHeader("To", to)
	m.SetHeader("subject", "Email verification")
	body := fmt.Sprintf(`
		<h1>Email verification</h1>
		<a href="http://localhost:8080/verify/%s">Click here</a>
	`,verificationToken)
	m.SetBody("text/html", body)
	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("SENDEREMAIL"), os.Getenv("EMAILPASS"))
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
