package ttlock

type GatewayListRequestParams struct {
	Auth
	PageNo   int   `json:"pageNo" validate:"required,min=1"`
	PageSize int   `json:"pageSize" validate:"required,min=1,max=200"`
	OrderBy  int   `json:"orderBy" validate:"required,oneof=0 1 2"`
	Date     int64 `json:"date" validate:"required"`
}

// {
// 	"gatewayId": 78979,
// 	"gatewayMac": "C5:40:E0:9C:8C:C1",
// 	"gatewayVersion":1,
// 	"networkName": "1-101-WIFI",
// 	"lockNum": 2,
// 	"isOnline": 1
// }

type GatewayObject struct {
	GatewayId      int    `json:"gatewayId"`
	GatewayMac     string `json:"gatewayMac"`
	GatewayVersion int    `json:"gatewayVersion"`
	NetworkName    string `json:"networkName"`
	LockNum        int    `json:"lockNum"`
	IsOnline       int    `json:"isOnline"`
}

type GatewayListResponse struct {
	List     []GatewayObject `json:"list"`
	PageNo   int             `json:"pageNo"`
	PageSize int             `json:"pageSize"`
	Total    int             `json:"total"`
	Pages    int             `json:"pages"`
	Err
}

// Delete

type GatewayDeleteRequestParams struct {
	Auth
	GatewayId int   `json:"gatewayId" validate:"required"`
	Date      int64 `json:"date" validate:"required"`
}

// Rename

type GatewayRenameRequestParams struct {
	Auth
	GatewayId   int    `json:"gatewayId" validate:"required"`
	GatewayName string `json:"gatewayName" validate:"required"`
	Date        int64  `json:"date" validate:"required"`
}

// List by lock

type GatewayListByLockRequestParams struct {
	Auth
	LockId int   `json:"lockId" validate:"required"`
	Date   int64 `json:"date" validate:"required"`
}

// {
// "gatewayId": 78979,
// "gatewayMac": "C5:40:E0:9C:8C:C1",
// "gatewayName": "Gateway for 1-101",
// "rssi": -65,
// "rssiUpdateDate": 1626674053000
// }

type GatewayObjectByLock struct {
	GatewayId      int    `json:"gatewayId"`
	GatewayMac     string `json:"gatewayMac"`
	GatewayName    string `json:"gatewayName"`
	Rssi           int    `json:"rssi"`
	RssiUpdateDate int64  `json:"rssiUpdateDate"`
}

type GatewayListByLockResponse struct {
	List []GatewayObjectByLock `json:"list"`
	Err
}

// Lock List

type GatewayListLockRequestParams struct {
	Auth
	GatewayId int   `json:"gatewayId" validate:"required"`
	Date      int64 `json:"date" validate:"required"`
}

// {
// "lockId": 532323,
// "lockName":"YS1003_c18c9c",
// "lockAlias":"Front door lock",
// "lockMac": "C5:40:E0:9C:8C:C1",
// "rssi":-65,
// "updateDate": 1626674053000
// }

type GatewayLockObject struct {
	LockId     int    `json:"lockId"`
	LockName   string `json:"lockName"`
	LockAlias  string `json:"lockAlias"`
	LockMac    string `json:"lockMac"`
	Rssi       int    `json:"rssi"`
	UpdateDate int64  `json:"updateDate"`
}

type GatewayListLockResponse struct {
	List []GatewayLockObject `json:"list"`
	Err
}
