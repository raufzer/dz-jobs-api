package utils

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-redis/redis/v8"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"golang.org/x/oauth2"
)

func CheckDatabaseHealth(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		log.Printf("Database health check failed: %v", err)
		return err
	}
	log.Println("Database is healthy")
	return nil
}

func CheckCacheHealth(redisClient *redis.Client) error {
	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Printf("Cache health check failed: %v", err)
		return err
	}
	log.Println("Cache is healthy")
	return nil
}

func CheckSendGridHealth(apiKey string) error {
	client := sendgrid.NewSendClient(apiKey)

	from := mail.NewEmail("Health Check", "dzjobs.service@gmail.com")
	to := mail.NewEmail("Health Check", "zerkhefraouf90@gmail.com")
	message := mail.NewSingleEmail(from, "Health Check", to, "This is a test.", "<p>This is a test.</p>")

	response, err := client.Send(message)
	if err != nil {
		log.Printf("SendGrid health check failed: %v", err)
		return err
	}

	if response.StatusCode >= 400 {
		log.Printf("SendGrid health check failed. Status: %d, Body: %s", response.StatusCode, response.Body)
		return fmt.Errorf("SendGrid returned error with status code %d: %s", response.StatusCode, response.Body)
	}

	log.Println("SendGrid is healthy")
	return nil
}

func CheckCloudinaryHealth(cloudName, apiKey, apiSecret string) error {
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		log.Printf("Cloudinary health check failed: %v", err)
		return err
	}

	_, err = cld.Admin.Ping(context.Background())
	if err != nil {
		log.Printf("Cloudinary health check failed: %v", err)
		return err
	}
	log.Println("Cloudinary is healthy")
	return nil
}

func CheckGoogleOAuthHealth(oauthConfig *oauth2.Config) error {

	url := oauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Google OAuth health check failed: %v", err)
		return fmt.Errorf("failed to reach Google OAuth authorization endpoint: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Google OAuth health check failed with status code: %d", resp.StatusCode)
		return fmt.Errorf("google OAuth returned error with status code %d", resp.StatusCode)
	}

	log.Println("Google OAuth is healthy")
	return nil
}
