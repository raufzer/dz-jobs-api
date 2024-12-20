package integrations

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendOTPEmail(email, otp, sendGridAPIKey string) error {
	if sendGridAPIKey == "" {
		return fmt.Errorf("SendGrid API key is missing")
	}

	client := sendgrid.NewSendClient(sendGridAPIKey)
	subject := "Dz Jobs password assistance"

	from := mail.NewEmail("Dz Jobs", "dzjobs.service@gmail.com")
	to := mail.NewEmail(email, email)

	templatePath := filepath.Join("internal", "templates", "otp_email_template.html")

	templateBytes, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read email template at %s: %w", templatePath, err)
	}

	emailBodyHTML := strings.ReplaceAll(string(templateBytes), "{{OTP}}", otp)
	emailBodyPlainText := "Your OTP is: " + otp

	message := mail.NewSingleEmail(from, subject, to, emailBodyPlainText, emailBodyHTML)

	response, err := client.Send(message)
	if err != nil {
		return fmt.Errorf("failed to send OTP email: %w", err)
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("failed to send OTP email: received status code %d", response.StatusCode)
	}

	return nil
}
