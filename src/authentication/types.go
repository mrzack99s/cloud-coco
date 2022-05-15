package authentication

type CredentialParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CredentialResponse struct {
	Email        string `json:"email"`
	UUID         string `json:"uuid"`
	NeedChPasswd bool   `json:"need_chpasswd"`
	Expires      string `json:"exp"`
}

type AccessDetails struct {
	AccessUuid string
	Email      string
}

type TokenDetails struct {
	Issue       string `json:"Issue"`
	AccessToken string `json:"access_token"`
	AccessUuid  string `json:"access_uuid"`
	AtExpires   string `json:"access_token_expires"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	AtExpires    string `json:"access_token_expires"`
	NeedChPasswd bool   `json:"need_chpasswd"`
}
