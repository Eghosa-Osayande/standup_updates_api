package jwt

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct {
	jwt.RegisteredClaims
	CustomClaims map[string]any `json:"custom_claims"`
}

func GenerateJwtToken(customClaims map[string]any, tokenSecret string,duration time.Duration) (string, error) {
	// accessDuration := 30 * 24 * time.Hour
	accessclaims := JWTData{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration).UTC()),
		},
		CustomClaims: customClaims,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessclaims)
	tokenString, err := token.SignedString([]byte(tokenSecret))

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil

}

func extractClaims(accessToken string, secret string) (*JWTData, error) {

	claims := &JWTData{}
	token, err := jwt.ParseWithClaims(accessToken, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("unexpected signing method in auth token")
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		log.Println("unable to parse claims", "error", err)
		return nil, err
	}

	if !token.Valid {
		return nil, err
	} else {

		return claims, nil
	}

}
