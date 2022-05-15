package models

type RBACSubscriptions struct {
	SubscriptionID *uint         `json:"subscription_id" gorm:"primaryKey;autoIncrement:false;not null;"`
	Subscription   Subscriptions `json:"subscription,omitempty" valid:"-"  gorm:"references:id"`
	RoleID         *uint         `json:"role_id" gorm:"primaryKey;autoIncrement:false;not null;"`
	Role           Roles         `json:"role,omitempty" valid:"-"  gorm:"references:id"`
	UserID         *uint         `json:"user_id" gorm:"primaryKey;autoIncrement:false;not null;"`
	User           Users         `json:"user,omitempty" valid:"-"  gorm:"references:id"`
}
