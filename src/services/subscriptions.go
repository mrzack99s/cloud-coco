package services

import (
	"errors"

	"github.com/mrzack99s/cloud-coco/src/configures"
	"github.com/mrzack99s/cloud-coco/src/models"
)

func SoftDeleteSubscriptions(uuid string) (err error) {

	o := models.Subscriptions{}
	configures.DBInstance().Where("uuid = ?", uuid).Preload("ResourcePools").Find(&o)
	if len(o.ResourcePools) > 0 {
		err = errors.New("cannot delete, found resource pools")
		return
	}

	err = configures.DBInstance().Where("uuid = ?", uuid).Delete(&models.Subscriptions{}).Error
	if err != nil {
		return
	}
	return
}

func HardDeleteSubscriptions(uuid string) (err error) {
	o := models.Subscriptions{}
	configures.DBInstance().Where("uuid = ?", uuid).Preload("ResourcePools").Find(&o)
	if len(o.ResourcePools) > 0 {
		err = errors.New("cannot delete, found resource pools")
		return
	}

	err = configures.DBInstance().Unscoped().Where("uuid = ?", uuid).Delete(&models.Subscriptions{}).Error
	if err != nil {
		return
	}
	return
}
