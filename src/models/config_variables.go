package models

type ConfigVariables struct {
	Name  string `json:"name,omitempty" valid:"resourceNameRegex, required" gorm:"primarykey"`
	Value string `json:"value,omitempty" valid:"required"`
}
