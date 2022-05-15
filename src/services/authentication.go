package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mrzack99s/cloud-coco/src/authentication"
	"github.com/mrzack99s/cloud-coco/src/configures"
	"github.com/mrzack99s/cloud-coco/src/models"
	"github.com/mrzack99s/cloud-coco/src/utils"
)

func CheckCredential(params authentication.CredentialParams) (r authentication.TokenResponse, err error) {
	o := models.Users{}
	found_err := configures.DBInstance().Where("email = ?", params.Email).First(&o).Error
	if found_err != nil {
		err = errors.New("not found user")
		return
	}

	if utils.Sha512encode(params.Password) != o.PasswdChecksum {
		err = errors.New("credential is not correct")
		return
	}

	var token *authentication.TokenDetails
	token, err = authentication.GenerateToken(o.UUID)
	if err != nil {
		return
	}

	r = authentication.TokenResponse{
		AccessToken:  token.AccessToken,
		AtExpires:    token.AtExpires,
		NeedChPasswd: o.NeedChPasswd,
	}

	return

}

func GetCredential(accessToken string) (r models.Users, err error) {
	if utils.RedisFindExistingKey(fmt.Sprintf("token:access/%s", accessToken)) {
		tokenDetailStr, e := configures.CacheInstance().Get(fmt.Sprintf("token:access/%s", accessToken)).Result()
		if e != nil {
			err = e
			return
		}

		tokenDetail := authentication.TokenDetails{}
		json.Unmarshal([]byte(tokenDetailStr), &tokenDetail)

		e = configures.DBInstance().Where("uuid = ?", tokenDetail.Issue).First(&r).Error
		if e != nil {
			err = e
			return
		}

		return
	} else {
		err = errors.New("not found token")
		return
	}

}

func RevokeCredential(accessToken string) (err error) {
	if utils.RedisFindExistingKey(fmt.Sprintf("token:access/%s", accessToken)) {
		tokenDetailStr, e := configures.CacheInstance().Get(fmt.Sprintf("token:access/%s", accessToken)).Result()
		if e != nil {
			err = e
			return
		}

		tokenDetail := authentication.TokenDetails{}
		json.Unmarshal([]byte(tokenDetailStr), &tokenDetail)

		utils.RedisDeleteWithKey(fmt.Sprintf("token:access/%s", tokenDetail.AccessUuid))
		return
	} else {
		err = errors.New("not found token")
		return
	}

}
