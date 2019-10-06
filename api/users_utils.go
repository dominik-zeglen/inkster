package api

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func createPasswordResetToken(
	userID int,
	currentTime time.Time,
	secretKey string,
) (string, error) {
	claims := ActionTokenClaims{
		ID:            userID,
		AllowedAction: RESET_PASSWORD,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: currentTime.
				Add(time.Hour * 24).
				Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := fmt.Sprintf("%x", secretKey)

	return token.SignedString([]byte(key))
}
