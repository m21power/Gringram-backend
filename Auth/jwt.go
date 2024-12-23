package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/m21power/GrinGram/types"
)

var secret = []byte("mesay")

// Create a new token object, specifying signing method and the claims
// you would like it to contain.
func GenerateToken(username string, role string, userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     username,
		"role":    role,
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secret)

	return tokenString, err
}
func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	})
	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// check exp and here
		exp := claims["exp"]
		if float64(time.Now().Unix()) > exp.(float64) {
			return false, fmt.Errorf("the token is expired")
		}
	}
	return true, nil

}
func GetTokenValues(tokenString string) (*types.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	var role string
	var username string
	var userID int
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// check exp and here
		exp := claims["exp"]
		if float64(time.Now().Unix()) > exp.(float64) {
			return nil, fmt.Errorf("the token is expired")
		}
		role = claims["role"].(string)
		username = claims["sub"].(string)
		userID = int(claims["user_id"].(float64))

	}
	return &types.Token{UserID: userID, Role: role, Username: username}, nil
}
func GetTokens(r *http.Request) (string, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
