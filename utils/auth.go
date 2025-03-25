package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(payload map[string]string, ExpiredAt time.Duration) (string, error) {
	secretKey := os.Getenv("API_KEY")
	expiredAt := time.Now().Add(time.Duration(time.Hour) * ExpiredAt).Unix()

	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt

	for i, v := range payload {
		claims[i] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return accessToken, err
	}

	return accessToken, nil
}

func GetPayloadInsideToken(tokenString string) (map[string]string, error) {
	secretKey := os.Getenv("API_KEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token expired")
		}
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	payload := make(map[string]string)
	for key, value := range claims {
		if strValue, ok := value.(string); ok {
			payload[key] = strValue
		}
	}

	return payload, nil
}
