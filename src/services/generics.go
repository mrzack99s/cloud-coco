package services

import (
	"github.com/mrzack99s/cloud-coco/src/configures"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Create(obj interface{}) (err error) {
	err = configures.DBInstance().Create(obj).Error
	return
}

func Update(obj interface{}) (err error) {
	err = configures.DBInstance().First(obj).Error
	if err != nil {
		return
	}

	err = configures.DBInstance().Model(obj).Updates(*obj.(*interface{})).Error
	if err != nil {
		return
	}

	return
}

func SoftDelete(uuid string, o interface{}) (err error) {
	err = configures.DBInstance().Where("uuid = ?", uuid).Delete(o).Error
	return
}

func HardDelete(uuid string, o interface{}) (err error) {
	err = configures.DBInstance().Where("uuid = ?", uuid).Unscoped().Delete(o).Error
	return
}

func GetByUUID(uuid string, o interface{}) (err error) {
	err = configures.DBInstance().Where("uuid = ?", uuid).Preload(clause.Associations).First(o).Error
	return
}

func GetByOffset(page, limit int, o interface{}) (err error) {
	err = configures.DBInstance().Offset(page).Limit(limit).Preload(clause.Associations).Find(o).Error
	return
}

func CountAll(o interface{}) (count int64, err error) {
	err = configures.DBInstance().Model(o).Count(&count).Error
	return
}

func GetWithCondition(condition *gorm.DB, o interface{}) (err error) {
	err = configures.DBInstance().Where(condition).Preload(clause.Associations).First(o).Error
	return
}
