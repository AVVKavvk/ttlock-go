package ttlock

type Err struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type Auth struct {
	ClientId    string `json:"clientId" validate:"required"`
	AccessToken string `json:"accessToken" validate:"required"`
}
