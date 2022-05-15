package models

type Roles struct {
	BaseModel
	Name              string              `json:"name,omitempty" valid:"nameWithDigitRegex, required" gorm:"unique;not null;"`
	Permissions       []RolesPermissions  `json:"permissions,omitempty" valid:"-" gorm:"foreignKey:role_id"`
	RBACResourcePools []RBACResourcePools `json:"rbac_resource_pools,omitempty" valid:"-" gorm:"foreignKey:role_id;constraint:OnDelete:CASCADE;"`
	RBACSubscriptions []RBACSubscriptions `json:"rbac_subscriptions,omitempty" valid:"-" gorm:"foreignKey:role_id;constraint:OnDelete:CASCADE;"`
}
