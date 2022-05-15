package models

type Directories struct {
	BaseModel
	Disabled      bool               `json:"disabled" valid:"-" gorm:"default:false"`
	Name          string             `json:"name,omitempty" valid:"nameRegex, required" gorm:"unique"`
	Users         []DirectoriesUsers `json:"users,omitempty" valid:"-" gorm:"foreignKey:directory_id"`
	Subscriptions []Subscriptions    `json:"subscriptions,omitempty" valid:"-" gorm:"foreignKey:directory_id"`
}
