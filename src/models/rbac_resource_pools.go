package models

type RBACResourcePools struct {
	ResourcePoolID *uint         `json:"resource_pool_id" gorm:"primaryKey;autoIncrement:false;not null;"`
	ResourcePool   ResourcePools `json:"resource_pool,omitempty" valid:"-"  gorm:"references:id"`
	RoleID         *uint         `json:"role_id" gorm:"primaryKey;autoIncrement:false;not null;"`
	Role           Roles         `json:"role,omitempty" valid:"-"  gorm:"references:id"`
	UserID         *uint         `json:"user_id" gorm:"primaryKey;autoIncrement:false;not null;"`
	User           Users         `json:"user,omitempty" valid:"-"  gorm:"references:id"`
}
