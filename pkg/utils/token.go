package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

func GenerateToken(ttl time.Duration, payload interface{}, secretJWTKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Current time (UTC)
	now := time.Now().UTC()

	// Create claims
	claims := token.Claims.(jwt.MapClaims)

	// Include the payload
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	// Generate a unique JWT ID using current timestamp and random bytes
	claims["jti"] = fmt.Sprintf("%d-%x", now.UnixNano(), generateRandomBytes(16)) // 16 bytes of randomness

	// Generate the token
	tokenString, err := token.SignedString([]byte(secretJWTKey))
	if err != nil {
		return "", fmt.Errorf("generating JWT Token failed: %w", err)
	}

	// Log token details (optional, for debugging)
	fmt.Println("Generated Token:")
	fmt.Println("Token String:", tokenString)
	fmt.Println("Claims:", claims)

	return tokenString, nil
}

// Helper to generate random bytes for jti
func generateRandomBytes(size int) []byte {
	randomBytes := make([]byte, size)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Printf("Error generating random bytes: %v", err)
	}
	return randomBytes
}

func GenerateSecureOTP(length int) string {
	const charset = "0123456789"
	otp := make([]byte, length)
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Printf("Error generating OTP: %v", err)
		return ""
	}

	for i := 0; i < length; i++ {
		otp[i] = charset[int(randomBytes[i])%len(charset)]
	}

	return string(otp)
}
