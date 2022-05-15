package services

import (
	"errors"

	"github.com/mrzack99s/cloud-coco/src/configures"
	"github.com/mrzack99s/cloud-coco/src/models"
)

func SoftDeleteServiceVersions(uuid string) (err error) {
	o := models.ServiceVersions{}
	configures.DBInstance().Where("uuid = ?", uuid).Preload("Resources").Find(&o)
	if len(o.Resources) > 0 {
		err = errors.New("cannot delete, found resource usage")
		return
	}

	err = configures.DBInstance().Where("uuid = ?", uuid).Delete(&models.ServiceVersions{}).Error
	if err != nil {
		return
	}
	return
}

func HardDeleteServiceVersions(uuid string) (err error) {
	o := models.ServiceVersions{}
	configures.DBInstance().Where("uuid = ?", uuid).Preload("Resources").Find(&o)
	if len(o.Resources) > 0 {
		err = errors.New("cannot delete, found resource usage")
		return
	}

	err = configures.DBInstance().Unscoped().Where("uuid = ?", uuid).Delete(&models.ServiceVersions{}).Error
	if err != nil {
		return
	}
	return
}

func GetByServicesVersionBySID(sid uint, page, limit int) (rCount int64, o []models.ServiceVersions, err error) {
	err = configures.DBInstance().Model(&models.ServiceVersions{}).Where("service_id = ?", sid).Count(&rCount).Error
	if err != nil {
		return
	}

	err = configures.DBInstance().Offset(page).Limit(limit).Where("service_id = ?", sid).Find(&o).Error
	if err != nil {
		return
	}

	return
}
