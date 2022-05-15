package models

import (
	"gorm.io/gorm"
)

type Subscriptions struct {
	BaseModel
	Name          string              `json:"name,omitempty" valid:"nameWithDigitRegex" gorm:"unique;"`
	Disabled      bool                `json:"disabled" valid:"-" gorm:"default:false"`
	DirectoryID   *uint               `json:"directory_id,omitempty" valid:"-" gorm:"not null;"`
	Directory     Directories         `json:"directory,omitempty" valid:"-"  gorm:"references:id"`
	ResourcePools []ResourcePools     `json:"resource_pools,omitempty" valid:"-" gorm:"foreignKey:subscription_id"`
	RBAC          []RBACSubscriptions `json:"rbac,omitempty" valid:"-" gorm:"foreignKey:subscription_id;constraint:OnDelete:CASCADE;"`
}

func (o *Subscriptions) AfterCreate(tx *gorm.DB) (err error) {
	if o.Name == "" {
		tx.Model(o).Update("Name", o.UUID)
	}
	return
}
