package models

type ServiceVersions struct {
	BaseModel
	ServiceID *uint       `json:"service_id,omitempty" valid:"-" gorm:"not null;"`
	Service   Services    `json:"service,omitempty" valid:"-"  gorm:"references:id"`
	Version   string      `json:"version" valid:"-" gorm:"not null;"`
	Image     string      `json:"image" valid:"-" gorm:"not null;"`
	Resources []Resources `json:"resources,omitempty" valid:"-" gorm:"foreignKey:service_id"`
}
