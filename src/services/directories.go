package services

import (
	"errors"

	"github.com/mrzack99s/cloud-coco/src/configures"
	"github.com/mrzack99s/cloud-coco/src/models"
)

func SoftDeleteDirectories(uuid string) (err error) {
	o := models.Directories{}
	configures.DBInstance().Where("uuid = ?", uuid).Preload("Subscriptions").Find(&o)
	if len(o.Subscriptions) > 0 {
		err = errors.New("cannot delete, found subscriptions")
		return
	}

	err = configures.DBInstance().Where("uuid = ?", uuid).Delete(&models.Directories{}).Error
	if err != nil {
		return
	}
	return
}

func HardDeleteDirectories(uuid string) (err error) {
	o := models.Directories{}
	configures.DBInstance().Where("uuid = ?", uuid).Preload("Subscriptions").Find(&o)
	if len(o.Subscriptions) > 0 {
		err = errors.New("cannot delete, found subscriptions")
		return
	}

	err = configures.DBInstance().Unscoped().Where("uuid = ?", uuid).Delete(&models.Directories{}).Error
	if err != nil {
		return
	}
	return
}
