package models

type ResourcePools struct {
	BaseModel
	Name           string              `json:"name,omitempty" valid:"nameWithDigitRegex, required" gorm:"unique;not null;"`
	SubscriptionID *string             `json:"subscription_id,omitempty" valid:"-" gorm:"not null;"`
	Subscription   Subscriptions       `json:"subscription,omitempty" valid:"-"  gorm:"references:id"`
	Resources      []Resources         `json:"resources,omitempty" valid:"-" gorm:"foreignKey:resource_pool_id"`
	RBAC           []RBACResourcePools `json:"rbac,omitempty" valid:"-" gorm:"foreignKey:resource_pool_id;constraint:OnDelete:CASCADE;"`
}
