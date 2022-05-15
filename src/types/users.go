package types

type UserChangePasswdParams struct {
	UUID        uint   `json:"uuid"`
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

type UserResetPasswdParams struct {
	UUID string `json:"uuid"`
}
