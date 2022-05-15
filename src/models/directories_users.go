package models

type DirectoriesUsers struct {
	DirectoryID *uint       `json:"directory_id" gorm:"primaryKey;autoIncrement:false"`
	Directory   Directories `json:"directory,omitempty" valid:"-"  gorm:"references:id"`
	UserID      *uint       `json:"user_id" gorm:"primaryKey;autoIncrement:false"`
	User        Users       `json:"users,omitempty" valid:"-"  gorm:"references:id"`
}
