package ttlock

// UserList
type UserListRequestParams struct {
	ClientId     string `json:"clientId" validate:"required"`
	ClientSecret string `json:"clientSecret" validate:"required"`
	StartDate    int64  `json:"startDate"`
	EndDate      int64  `json:"endDate"`
	PageNo       int    `json:"pageNo" validate:"required,min=1"`
	PageSize     int    `json:"pageSize" validate:"required,min=1,max=100"`
	Date         int64  `json:"date" validate:"required"`
	Username     string `json:"username"`
}

type UserListResponseListObject struct {
	Username string `json:"username"`
	Regtime  int64  `json:"regtime"`
}

type UserListResponse struct {
	List     []UserListResponseListObject `json:"list"`
	PageNo   int                          `json:"pageNo"`
	PageSize int                          `json:"pageSize"`
	Total    int                          `json:"total"`
	Pages    int                          `json:"pages"`
}

// UserRegister
type UserRegisterRequestParams struct {
	ClientId     string `json:"clientId" validate:"required"`
	ClientSecret string `json:"clientSecret" validate:"required"`
	Username     string `json:"username" validate:"required,alphanum"`
	Password     string `json:"password" validate:"required"`
	Date         int64  `json:"date" validate:"required"`
}

type UserRegisterResponse struct {
	Username string `json:"username"`
	ErrCode  int    `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
}

// User Password Reset
type UserPasswordResetRequestParams struct {
	ClientId     string `json:"clientId" validate:"required"`
	ClientSecret string `json:"clientSecret" validate:"required"`
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	Date         int64  `json:"date" validate:"required"`
}

type UserPasswordResetResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// User Delete
type UserDeleteRequestParams struct {
	ClientId     string `json:"clientId" validate:"required"`
	ClientSecret string `json:"clientSecret" validate:"required"`
	Username     string `json:"username" validate:"required"`
	Date         int64  `json:"date" validate:"required"`
}

type UserDeleteResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
