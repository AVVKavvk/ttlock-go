package ttlock

// Send
type EKeySendRequestParams struct {
	ClientId         string `json:"clientId" validate:"required"`
	AccessToken      string `json:"accessToken" validate:"required"`
	LockId           int    `json:"lockId" validate:"required"`
	ReceiverUsername string `json:"receiverUsername" validate:"required"`
	KeyName          string `json:"keyName" validate:"required"`
	StartDate        int64  `json:"startDate" validate:"required"`
	EndDate          int64  `json:"endDate" validate:"required"`
	Remarks          string `json:"remarks"`
	RemoteEnable     int    `json:"remoteEnable" validate:"oneof=1 2"`
	KeyRight         int    `json:"keyRight" validate:"oneof=1 2"`
	CreateUser       int    `json:"createUser" validate:"oneof=1 2"`
	Date             int64  `json:"date" validate:"required"`
}
type EKeySendResponse struct {
	KeyId   int    `json:"keyId"`
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// List

type EKeyListRequestParams struct {
	Auth
	PageNo    int    `json:"pageNo" validate:"required"`
	PageSize  int    `json:"pageSize" validate:"required"`
	Date      int64  `json:"date" validate:"required"`
	LockAlias string `json:"lockAlias"`
	GroupId   int    `json:"groupId"`
}

// {
// 	"keyId": 3234293,
// 	"lockData": "xxxxxxx",
// 	"lockId": 532323,
// 	"userType": "110302",
// 	"keyStatus": "110401",
// 	"lockName": "M102_ed8e8d",
// 	"lockAlias": "Front door lock",
// 	"lockMac": "68:9E:19:00:67:81",
// 	"noKeyPwd": "345634",
// 	"electricQuantity": 85,
// 	"startDate": 1528878944000,
// 	"endDate": 1628878944000,
// 	"remarks": "Wish you a happy rent.",
// 	"keyRight": 1,
// 	"featureValue": "6C4CD1F3",
// 	"remoteEnable": 0,
// 	"passageMode": 1,
// 	"groupId": 7619,
// 	"groupName": "The 4th floor"
// }

type EKeyObject struct {
	KeyId        int    `json:"keyId"`
	LockData     string `json:"lockData"`
	LockId       int    `json:"lockId"`
	UserType     string `json:"userType"`
	KeyStatus    string `json:"keyStatus"`
	LockName     string `json:"lockName"`
	LockAlias    string `json:"lockAlias"`
	LockMac      string `json:"lockMac"`
	NoKeyPwd     string `json:"noKeyPwd"`
	ElectricQty  int    `json:"electricQuantity"`
	StartDate    int64  `json:"startDate"`
	EndDate      int64  `json:"endDate"`
	Remarks      string `json:"remarks"`
	KeyRight     int    `json:"keyRight"`
	FeatureValue string `json:"featureValue"`
	RemoteEnable int    `json:"remoteEnable"`
	PassageMode  int    `json:"passageMode"`
	GroupId      int    `json:"groupId"`
	GroupName    string `json:"groupName"`
}

type EKeyListResponse struct {
	List     []EKeyObject `json:"list"`
	ErrCode  int          `json:"errcode"`
	ErrMsg   string       `json:"errmsg"`
	PageNo   int          `json:"pageNo"`
	PageSize int          `json:"pageSize"`
	Pages    int          `json:"pages"`
	Total    int          `json:"total"`
}

// Get One

type EKeyGetOneRequestParams struct {
	Auth
	LockId int   `json:"lockId" validate:"required"`
	Date   int64 `json:"date" validate:"required"`
}

type EKeyGetOneResponse struct {
	EKeyObject
	Err
}

// Delete

type EKeyDeleteRequestParams struct {
	Auth
	KeyId int   `json:"keyId" validate:"required"`
	Date  int64 `json:"date" validate:"required"`
}

// Freeze/Unfreeze
type EKeyFreezeUnfreezeRequestParams struct {
	Auth
	KeyId int   `json:"keyId" validate:"required"`
	Date  int64 `json:"date" validate:"required"`
}

// Modify
type EKeyUpdateRequestParams struct {
	Auth
	KeyId        int    `json:"keyId" validate:"required"`
	Date         int64  `json:"date" validate:"required"`
	KeyName      string `json:"keyName"`
	RemoteEnable int    `json:"remoteEnable" validate:"oneof=1 2"`
}

// Change EKeys Valid Time
type EKeyChangeValidTimeRequestParams struct {
	Auth
	KeyId     int   `json:"keyId" validate:"required"`
	StartDate int64 `json:"startDate" validate:"required"`
	EndDate   int64 `json:"endDate" validate:"required"`
	Date      int64 `json:"date" validate:"required"`
}
