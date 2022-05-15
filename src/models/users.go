package models

type Users struct {
	BaseModel
	FirstName         string              `json:"first_name,omitempty" valid:"nameRegex, required" gorm:"not null;"`
	LastName          string              `json:"last_name,omitempty" valid:"nameRegex, required" gorm:"not null;"`
	Email             string              `json:"email,omitempty" valid:"email, required" gorm:"unique;not null;"`
	NeedChPasswd      bool                `json:"need_chpasswd" valid:"-" gorm:"default:true"`
	TemporaryPassword string              `json:"temporary_password,omitempty" gorm:"-"`
	PasswdChecksum    string              `json:"passwd_checksum,omitempty" valid:"sha512Regex, required" gorm:"not null;"`
	Directories       []DirectoriesUsers  `json:"directories,omitempty" valid:"-" gorm:"foreignKey:user_id"`
	RBACSubscriptions []RBACSubscriptions `json:"rbac_subscriptions,omitempty" valid:"-" gorm:"foreignKey:user_id;constraint:OnDelete:CASCADE;"`
	RBACResourcePools []RBACResourcePools `json:"rbac_resource_pools,omitempty" valid:"-" gorm:"foreignKey:user_id;constraint:OnDelete:CASCADE;"`
}
