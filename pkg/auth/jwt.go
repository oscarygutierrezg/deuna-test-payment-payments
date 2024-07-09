package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

const JwtAuthorizationHeader = "Authorization"
const JwtMessageUnauthorized = "Unauthorized"

const jwtEnvKey = "JWT_SECRET"

var signMethod = jwt.SigningMethodHS512
var defaultDuration = time.Hour * 72

func NewJwtToken(payload jwt.MapClaims, duration *time.Duration) (token string, err error) {
	if duration != nil {
		return signJwtToken(payload, *duration)
	}
	return signJwtToken(payload, defaultDuration)
}

func IsJwtTokenValid(jwtToken string) (bool, error) {
	var jwtSecret = os.Getenv(jwtEnvKey)
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, err
}

func GetJwtTokenMapClaims(jwtToken string) (jwt.MapClaims, error) {
	var jwtSecret = os.Getenv(jwtEnvKey)
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("jwt.MapClaims parse error")
	}

	return claims, nil
}

// private functions
func signJwtToken(payload jwt.MapClaims, duration time.Duration) (string, error) {

	// Set duration token
	payload["exp"] = time.Now().Add(duration).Unix()

	// Create the token
	token := jwt.NewWithClaims(signMethod, payload)

	// Get secret
	var jwtSecret = os.Getenv(jwtEnvKey)

	// Sign the token with our secret
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
