package authentication

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func ExtractTokenMetadata(r *http.Request) (string, error) {
	token, err := verifyToken(r)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return "", err
		}

		return accessUuid, nil
	}
	return "", err
}
