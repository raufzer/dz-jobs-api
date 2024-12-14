package utils

import (
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendOTP(email, otp, sendGridAPIKey string) error {
	if sendGridAPIKey == "" {
		log.Println("SendGrid API key is missing")
		return nil
	}

	client := sendgrid.NewSendClient(sendGridAPIKey)
	subject := "Your One-Time Password (OTP) for Dz Jobs"

	from := mail.NewEmail("Dz Jobs", "dzjobs.service@gmail.com")
	to := mail.NewEmail(email, email)
	message := mail.NewSingleEmail(from, subject, to, otp, otp)

	response, err := client.Send(message)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		return err
	}

	log.Printf("OTP email sent to %s. Status code: %d", email, response.StatusCode)
	return nil
}
