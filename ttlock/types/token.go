package ttlock

type AccessTokenRequestParams struct {
	ClientId     string `json:"clientId" validate:"required"`
	ClientSecret string `json:"clientSecret" validate:"required"`
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	UID          int    `json:"uid"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
}

type RefreshAccessTokenRequestParams struct {
	ClientId     string `json:"clientId" validate:"required"`
	ClientSecret string `json:"clientSecret" validate:"required"`
	GrantType    string `json:"grant_type" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
}
