package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mrzack99s/cloud-coco/src/configures"
)

func GenerateToken(user_uuid string) (*TokenDetails, error) {

	now := time.Now()
	atExp := now.Add(time.Minute * 45)
	td := &TokenDetails{}
	td.AtExpires = atExp.Format(time.RFC3339)
	td.AccessUuid = uuid.New().String()

	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["exp"] = td.AtExpires
	atClaims["issue"] = user_uuid
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(configures.Sys().COCO.Secret))
	if err != nil {
		return nil, err
	}

	td.Issue = user_uuid

	marshalTokenDetails, _ := json.Marshal(td)

	errAccess := configures.CacheInstance().Set(fmt.Sprintf("token:access/%s", td.AccessUuid), string(marshalTokenDetails), atExp.Sub(now)).Err()
	if errAccess != nil {
		return nil, errAccess
	}

	return td, nil
}

func TokenMiddleware(c *gin.Context) {
	err := TokenValid(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "not authorized")
		c.Abort()
		return
	}
	c.Next()

}
