package ttlock

// Initialize
type LockInitializeRequest struct {
	Auth
	LockData      string `json:"lockData" validate:"required"`
	Date          int64  `json:"date" validate:"required"`
	LockAlias     string `json:"lockAlias"`
	GroupId       int    `json:"groupId"`
	NBinitSuccess int    `json:"nbInitSuccess"`
}

type LockInitializeResponse struct {
	LockId  int    `json:"lockId"`
	KeyId   int    `json:"keyId"`
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// List

type LockListRequestParams struct {
	Auth
	PageNo    int    `json:"pageNo" validate:"required,min=1"`
	PageSize  int    `json:"pageSize" validate:"required,min=1,max=1000"`
	Date      int64  `json:"date" validate:"required"`
	LockAlias string `json:"lockAlias"`
	GroupId   int    `json:"groupId"`
}

//		{
//		"lockId": 532323,
//		"lockName":"YS1003_c18c9c",
//		"lockAlias":"Front door lock",
//		"lockMac": "C5:40:E0:9C:8C:C1",
//		"electricQuantity": 55,
//		"featureValue":"3F0421C4F5F3",
//		"hasGateway": 1,
//		"lockData": "xxxxxxxxxxxxx",
//		"groupId": 456,
//		"groupName": "The 4th floor",
//		"date": 1528878944000,
//	}
type LockListResponseObject struct {
	LockId       int    `json:"lockId"`
	LockName     string `json:"lockName"`
	LockAlias    string `json:"lockAlias"`
	LockMac      string `json:"lockMac"`
	ElectricQty  int    `json:"electricQuantity"`
	FeatureValue string `json:"featureValue"`
	HasGateway   int    `json:"hasGateway"`
	LockData     string `json:"lockData"`
	GroupId      int    `json:"groupId"`
	GroupName    string `json:"groupName"`
	Date         int64  `json:"date"`
}

// {
// "list": [
// 	{
// 		"lockId": 532323,
// 		"lockName":"YS1003_c18c9c",
// 		"lockAlias":"Front door lock",
// 		"lockMac": "C5:40:E0:9C:8C:C1",
// 		"electricQuantity": 55,
// 		"featureValue":"3F0421C4F5F3",
// 		"hasGateway": 1,
// 		"lockData": "xxxxxxxxxxxxx",
// 		"groupId": 456,
// 		"groupName": "The 4th floor",
// 		"date": 1528878944000,
// 	}
// ],
// "pageNo":1,
// "pageSize":20,
// "pages":1,
// "total":1
// }

type LockListResponse struct {
	List     []LockListResponseObject `json:"list"`
	PageNo   int                      `json:"pageNo"`
	PageSize int                      `json:"pageSize"`
	Total    int                      `json:"total"`
	Pages    int                      `json:"pages"`
	ErrCode  int                      `json:"errcode"`
	ErrMsg   string                   `json:"errmsg"`
}

// Details
type LockDetailsRequestParams struct {
	Auth
	LockId int   `json:"lockId" validate:"required"`
	Date   int64 `json:"date" validate:"required"`
}

// {
//     "date": 1749195089000,
//     "lockAlias": "S202F_55ce63",
//     "lockSound": 1,
//     "modelNum": "SN284-FINGER-T0_PV53",
//     "lockMac": "ED:AE:90:63:CE:55",
//     "privacyLock": 2,
//     "featureValue": "344CD5F7",
//     "soundVolume": 5,
//     "hasGateway": 0,
//     "autoLockTime": 5,
//     "lockName": "S202F_55ce63",
//     "resetButton": 1,
//     "firmwareRevision": "5.1.0.200705",
//     "tamperAlert": 1,
//     "noKeyPwd": "7780189",
//     "passageMode": 2,
//     "passageModeAutoUnlock": 2,
//     "timezoneRawOffset": 28800000,
//     "lockId": 3466251,
//     "electricQuantity": 100,
//     "hardwareRevision": "1.1",
// }

type LockDetailsResponse struct {
	LockId                int    `json:"lockId"`
	LockName              string `json:"lockName"`
	ModelNum              string `json:"modelNum"`
	LockSound             int    `json:"lockSound"`
	LockAlias             string `json:"lockAlias"`
	LockMac               string `json:"lockMac"`
	PrivacyLock           int    `json:"privacyLock"`
	ElectricQty           int    `json:"electricQuantity"`
	HardwareRevision      string `json:"hardwareRevision"`
	TimezoneRawOffset     int64  `json:"timezoneRawOffset"`
	PassageModeAutoUnlock int    `json:"passageModeAutoUnlock"`
	SoundVolume           int    `json:"soundVolume"`
	TamperAlert           int    `json:"tamperAlert"`
	AutoLockTime          int    `json:"autoLockTime"`
	NoKeyPwd              string `json:"noKeyPwd"`
	PassageMode           int    `json:"passageMode"`
	ResetButton           int    `json:"resetButton"`
	FeatureValue          string `json:"featureValue"`
	HasGateway            int    `json:"hasGateway"`
	LockData              string `json:"lockData"`
	GroupId               int    `json:"groupId"`
	GroupName             string `json:"groupName"`
	Date                  int64  `json:"date"`
	ErrCode               int    `json:"errcode"`
	ErrMsg                string `json:"errmsg"`
}

// Delete
type LockDeleteRequestParams struct {
	Auth
	LockId int   `json:"lockId" validate:"required"`
	Date   int64 `json:"date" validate:"required"`
}

// Update

type LockUpdateRequestParams struct {
	Auth
	LockId   int    `json:"lockId" validate:"required"`
	LockData string `json:"lockData" validate:"required"`
	Date     int64  `json:"date" validate:"required"`
}

// Rename

type LockRenameRequestParams struct {
	Auth
	LockId    int    `json:"lockId" validate:"required"`
	LockAlias string `json:"lockAlias" validate:"required"`
	Date      int64  `json:"date" validate:"required"`
}

// Change Admin Passcode

type LockChangeAdminPasscodeRequestParams struct {
	Auth
	LockId     int    `json:"lockId" validate:"required"`
	Password   string `json:"password" validate:"required"`
	ChangeType int    `json:"changeType" validate:"required,oneof=1 2"`
	Date       int64  `json:"date" validate:"required"`
}

// Auto Lock Time

type LockAutoLockTimeRequestParams struct {
	Auth
	LockId  int   `json:"lockId" validate:"required"`
	Seconds int   `json:"seconds" validate:"required"`
	Type    int   `json:"type" validate:"required,oneof=1 2"`
	Date    int64 `json:"date" validate:"required"`
}

// EKeys

type LockListEKeyRequestParams struct {
	Auth
	LockId    int    `json:"lockId" validate:"required"`
	SearchStr string `json:"searchStr"`
	KeyRight  int    `json:"keyRight" validate:"oneof=1 2"`
	OrderBy   int    `json:"orderBy" validate:"oneof=1 2"`
	PageNo    int    `json:"pageNo" validate:"required,min=1"`
	PageSize  int    `json:"pageSize" validate:"required,min=1,max=200"`
	Date      int64  `json:"date" validate:"required"`
}

// {
// 	"keyId": 3234293,
// 	"lockId": 532323,
// 	"username": "jack@google.com",
// 	"uid": 6355997,
// 	"keyName":"Ekey for Jack",
// 	"keyStatus": "110401",
// 	"startDate": 1528878944000,
// 	"endDate": 1628878944000,
// 	"keyRight": 1,
// 	"senderUsername": "alexa@google.com",
// 	"remarks": "Wish you a happy rent.",
// 	"date": 1528878944000,
// }

type LockEkeyObject struct {
	KeyId          int    `json:"keyId"`
	LockId         int    `json:"lockId"`
	Username       string `json:"username"`
	Uid            int    `json:"uid"`
	KeyName        string `json:"keyName"`
	KeyStatus      string `json:"keyStatus"`
	StartDate      int64  `json:"startDate"`
	EndDate        int64  `json:"endDate"`
	KeyRight       int    `json:"keyRight"`
	SenderUsername string `json:"senderUsername"`
	Remarks        string `json:"remarks"`
	Date           int64  `json:"date"`
}

type LockListEkeyResponse struct {
	List     []LockEkeyObject `json:"list"`
	PageNo   int              `json:"pageNo"`
	PageSize int              `json:"pageSize"`
	Total    int              `json:"total"`
	Pages    int              `json:"pages"`
	Err
}

// Lock Unlock

type LockLockUnlockRequestParams struct {
	Auth
	LockId int   `json:"lockId" validate:"required"`
	Date   int64 `json:"date" validate:"required"`
}

// Lock Query Open state

type LockQueryOpenStateRequestParams struct {
	Auth
	LockId int   `json:"lockId" validate:"required"`
	Date   int64 `json:"date" validate:"required"`
}

type LockQueryOpenStateResponse struct {
	State int `json:"state"`
	Err
}

// Lock Time

type LockTimeRequestParams struct {
	Auth
	LockId int   `json:"lockId" validate:"required"`
	Date   int64 `json:"date" validate:"required"`
}

type LockTimeResponse struct {
	Date int64 `json:"date"`
	Err
}

// Update time

type LockUpdateTimeRequestParams struct {
	Auth
	LockId int   `json:"lockId" validate:"required"`
	Date   int64 `json:"date" validate:"required"`
}

type LockUpdateTimeResponse struct {
	Date int64 `json:"date"`
	Err
}

// Query Battery

type LockQueryBatteryRequestParams struct {
	Auth
	LockId int   `json:"lockId" validate:"required"`
	Date   int64 `json:"date" validate:"required"`
}

type LockQueryBatteryResponse struct {
	ElectricQty int `json:"electricQuantity"`
	Err
}
