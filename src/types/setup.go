package types

type SetupParams struct {
	FirstName string `json:"first_name,omitempty" valid:"nameRegex, required"`
	LastName  string `json:"last_name,omitempty" valid:"nameRegex, required"`
	Email     string `json:"email,omitempty" valid:"email, required"`
	Password  string `json:"password,omitempty" gorm:"-"`
}
