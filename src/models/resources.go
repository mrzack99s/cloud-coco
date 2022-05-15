package models

type Resources struct {
	BaseModel
	Name             string          `json:"name,omitempty" valid:"resourceNameRegex, required" gorm:"unique;not null;"`
	ServiceID        *string         `json:"service_id,omitempty" valid:"-" gorm:"not null;"`
	Service          ServiceVersions `json:"service,omitempty" valid:"-"  gorm:"references:id"`
	ResourcePoolID   *uint           `json:"resource_pool_id,omitempty" valid:"-" gorm:"not null;"`
	ResourcePool     ResourcePools   `json:"resource_pool,omitempty" valid:"-"  gorm:"references:id"`
	ResourceStatusID *uint           `json:"resource_status_id,omitempty" valid:"-" gorm:"not null;"`
	ResourceStatus   ResourcesStatus `json:"resource_status,omitempty" valid:"-"  gorm:"references:id"`
}
