package authentication

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/mrzack99s/cloud-coco/src/configures"
)

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := r.Header.Get("x-access-token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configures.Sys().COCO.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		return err
	}
	if err := token.Claims.Valid(); err != nil {
		return err
	}
	return nil
}
