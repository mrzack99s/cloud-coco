package services

import (
	"errors"

	"github.com/mrzack99s/cloud-coco/src/configures"
	"github.com/mrzack99s/cloud-coco/src/models"
)

func SoftDeleteResourcePools(uuid string) (err error) {

	o := models.ResourcePools{}
	configures.DBInstance().Where("uuid = ?", uuid).Preload("Resources").Find(&o)
	if len(o.Resources) > 0 {
		err = errors.New("cannot delete, found resources")
		return
	}

	err = configures.DBInstance().Where("uuid = ?", uuid).Delete(&models.ResourcePools{}).Error
	if err != nil {
		return
	}
	return
}

func HardDeleteResourcePools(uuid string) (err error) {
	o := models.ResourcePools{}
	configures.DBInstance().Where("uuid = ?", uuid).Preload("Resources").Find(&o)
	if len(o.Resources) > 0 {
		err = errors.New("cannot delete, found resources")
		return
	}

	err = configures.DBInstance().Unscoped().Where("uuid = ?", uuid).Delete(&models.ResourcePools{}).Error
	if err != nil {
		return
	}
	return
}
