package integrations

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/markbates/goth/providers/google"
	"golang.org/x/oauth2"
)

type UserInfo struct {
	ID    string `json:"sub"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FetchGoogleUserInfo(oauthConfig *oauth2.Config, token *oauth2.Token) (UserInfo, error) {
	client := oauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return UserInfo{}, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	var userInfo UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return UserInfo{}, fmt.Errorf("failed to decode user info: %w", err)
	}
	return userInfo, nil
}

func InitializeGoogleOAuthConfig(clientID, clientSecret, redirectURL string) *oauth2.Config {
	if clientID == "" || clientSecret == "" || redirectURL == "" {
		panic(fmt.Errorf("google OAuth configuration is incomplete: clientID, clientSecret, or redirectURL is missing"))
	}

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}
