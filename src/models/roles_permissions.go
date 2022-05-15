package models

type RolesPermissions struct {
	RoleID       *uint       `json:"role_id" gorm:"primaryKey;autoIncrement:false;not null;"`
	Role         Roles       `json:"role,omitempty" valid:"-"  gorm:"references:id"`
	PermissionID *uint       `json:"permission_id" gorm:"primaryKey;autoIncrement:false;not null;"`
	Permission   Permissions `json:"permission,omitempty" valid:"-"  gorm:"references:id"`
}
