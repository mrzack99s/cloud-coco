package models

type ResourcesStatus struct {
	BaseModel
	Name      string      `json:"name,omitempty" valid:"required" gorm:"unique;not null;"`
	Resources []Resources `json:"resources,omitempty" valid:"-" gorm:"foreignKey:resource_status_id"`
}
