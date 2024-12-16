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

	now := time.Now().UTC()

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	claims["jti"] = fmt.Sprintf("%d-%x", now.UnixNano(), generateRandomBytes(16))

	tokenString, err := token.SignedString([]byte(secretJWTKey))
	if err != nil {
		return "", fmt.Errorf("generating JWT Token failed: %w", err)
	}

	fmt.Println("Generated Token:")
	fmt.Println("Token String:", tokenString)
	fmt.Println("Claims:", claims)

	return tokenString, nil
}

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

type TokenClaims struct {
	UserID string `json:"sub"`
	jwt.StandardClaims
}

func ValidateToken(accrefToken string, secretKey string) (string, error) {
	// Debug log to see the incoming token
	log.Println("Received token for validation:", accrefToken)

	// Parse the token with claims
	token, err := jwt.ParseWithClaims(accrefToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	// If error or invalid token, log and return
	if err != nil || !token.Valid {
		log.Println("Error parsing or invalid token:", err)
		return "", fmt.Errorf("invalid token")
	}

	// Log claims extraction step to debug
	if claims, ok := token.Claims.(*TokenClaims); ok {
		log.Println("Successfully extracted claims:", claims)
		log.Println("Extracted UserID:", claims.UserID)
		return claims.UserID, nil
	}

	// If no claims, log that as well
	log.Println("Failed to extract claims from token")
	return "", fmt.Errorf("invalid token claims")
}
