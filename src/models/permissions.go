package models

type Permissions struct {
	BaseModel
	Name  string             `json:"name,omitempty" valid:"required" gorm:"unique;not null;"`
	Roles []RolesPermissions `json:"permissions,omitempty" valid:"-" gorm:"foreignKey:permission_id"`
}
