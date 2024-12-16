package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/golang-jwt/jwt"
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

	return tokenString, nil
}

func generateRandomBytes(size int) []byte {
	randomBytes := make([]byte, size)
	_, err := rand.Read(randomBytes)
	if err != nil {

		return nil
	}
	return randomBytes
}

func GenerateSecureOTP(length int) string {
	const charset = "0123456789"
	otp := make([]byte, length)
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
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

	token, err := jwt.ParseWithClaims(accrefToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(*TokenClaims); ok {
		return claims.UserID, nil
	}

	return "", fmt.Errorf("invalid token claims")
}
