package services

import (
	"errors"

	"github.com/mrzack99s/cloud-coco/src/configures"
	"github.com/mrzack99s/cloud-coco/src/models"
	"github.com/mrzack99s/cloud-coco/src/types"
	"github.com/mrzack99s/cloud-coco/src/utils"
)

func CreateAdminUser(o *types.SetupParams) (r models.Users, err error) {
	found_err := configures.DBInstance().Where("email = ?", o.Email).First(&models.Users{}).Error
	if found_err == nil {
		err = errors.New("email is duplicate")
		return
	}

	r.PasswdChecksum = utils.Sha512encode(o.Password)
	r.Email = o.Email
	r.FirstName = o.FirstName
	r.LastName = o.LastName
	r.NeedChPasswd = false

	err = configures.DBInstance().Create(&r).Error
	if err != nil {
		return
	}

	return
}

func CreateUser(o *models.Users) (err error) {
	found_err := configures.DBInstance().Where("email = ?", o.Email).First(&models.Users{}).Error
	if found_err == nil {
		err = errors.New("email is duplicate")
		return
	}

	newPasswd := utils.PasswordGenerator(12)
	o.PasswdChecksum = utils.Sha512encode(newPasswd)
	o.NeedChPasswd = true
	o.TemporaryPassword = newPasswd
	err = configures.DBInstance().Create(o).Error
	if err != nil {
		return
	}

	return

}

func ResetUserPassword(uuid string) (r models.Users, err error) {
	found_err := configures.DBInstance().Where("uuid = ?", uuid).First(&r).Error
	if found_err != nil {
		err = errors.New("not found user")
		return
	}

	newPasswd := utils.PasswordGenerator(12)
	r.PasswdChecksum = utils.Sha512encode(newPasswd)
	r.TemporaryPassword = newPasswd
	r.NeedChPasswd = true
	err = configures.DBInstance().Model(&r).Updates(r).Error
	if err != nil {
		return
	}

	return

}

func ChangeUserPassword(params types.UserChangePasswdParams) (err error) {

	o := models.Users{}
	found_err := configures.DBInstance().Where("uuid = ?", params.UUID).First(&o).Error
	if found_err != nil {
		err = errors.New("not found user")
		return
	}

	if utils.Sha512encode(params.OldPassword) != o.PasswdChecksum {
		err = errors.New("current password is not correct")
		return
	}

	newPasswdhash := utils.Sha512encode(params.NewPassword)
	if newPasswdhash == o.PasswdChecksum {
		err = errors.New("your new password is current password")
		return
	} else {
		o.PasswdChecksum = newPasswdhash
		o.NeedChPasswd = false
		err = configures.DBInstance().Model(&o).Updates(o).Error
		if err != nil {
			return
		}
	}

	return

}
