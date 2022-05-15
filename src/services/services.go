package services

import (
	"errors"

	"github.com/mrzack99s/cloud-coco/src/configures"
	"github.com/mrzack99s/cloud-coco/src/models"
)

func SoftDeleteServices(uuid string) (err error) {
	o := models.Services{}
	configures.DBInstance().Where("uuid = ?", uuid).Preload("Versions").Find(&o)
	if len(o.Versions) > 0 {
		err = errors.New("cannot delete, found service versions")
		return
	}

	err = configures.DBInstance().Where("uuid = ?", uuid).Delete(&models.Services{}).Error
	if err != nil {
		return
	}
	return
}

func HardDeleteServices(uuid string) (err error) {
	o := models.Services{}
	configures.DBInstance().Where("uuid = ?", uuid).Preload("Versions").Find(&o)
	if len(o.Versions) > 0 {
		err = errors.New("cannot delete, found service versions")
		return
	}

	err = configures.DBInstance().Unscoped().Where("uuid = ?", uuid).Delete(&models.Services{}).Error
	if err != nil {
		return
	}
	return
}
