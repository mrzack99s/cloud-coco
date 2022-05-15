package models

type Services struct {
	BaseModel
	Name     string            `json:"name,omitempty" valid:"required" gorm:"unique;not null;"`
	Versions []ServiceVersions `json:"versions,omitempty" valid:"-" gorm:"foreignKey:service_id"`
}
