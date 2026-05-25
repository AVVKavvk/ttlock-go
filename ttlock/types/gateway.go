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

// List Device

type GatewayListDeviceRequestParams struct {
	Auth
	GatewayId int   `json:"gatewayId" validate:"required"`
	Date      int64 `json:"date" validate:"required"`
}

// {
// "deviceType": 0,
// "deviceId": 532323,
// "deviceName":"YS1003_c18c9c",
// "deviceAlias":"entrance lock",
// "deviceMac": "C5:40:E0:9C:8C:C1",
// "rssi":-65,
// "updateDate": 1626674053000
// }

type GatewayDeviceObject struct {
	DeviceType  int    `json:"deviceType"`
	DeviceId    int    `json:"deviceId"`
	DeviceName  string `json:"deviceName"`
	DeviceAlias string `json:"deviceAlias"`
	DeviceMac   string `json:"deviceMac"`
	Rssi        int    `json:"rssi"`
	UpdateDate  int64  `json:"updateDate"`
}

type GatewayListDeviceResponse struct {
	List []GatewayDeviceObject `json:"list"`
	Err
}

// Details

type GatewayDetailsRequestParams struct {
	Auth
	GatewayId int   `json:"gatewayId" validate:"required"`
	Date      int64 `json:"date" validate:"required"`
}

// {
// "gatewayMac": "DC:A4:53:85:39:74",
// "lockNum": 1,
// "gatewayName": "G2_743985",
// "networkName": "ttlock",
// "isOnline": 0,
// "gatewayVersion": 2,
// "gatewayId": 347
// }

type GatewayDetailsResponse struct {
	GatewayMac     string `json:"gatewayMac"`
	LockNum        int    `json:"lockNum"`
	GatewayName    string `json:"gatewayName"`
	NetworkName    string `json:"networkName"`
	IsOnline       int    `json:"isOnline"`
	GatewayVersion int    `json:"gatewayVersion"`
	GatewayId      int    `json:"gatewayId"`
	Err
}

// Upload Detail

type GatewayUploadDetailRequestParams struct {
	Auth
	GatewayId        int    `json:"gatewayId" validate:"required"`
	ModelNum         string `json:"modelNum" validate:"required"`
	HardwareRevision string `json:"hardwareRevision" validate:"required"`
	FirmwareRevision string `json:"firmwareRevision" validate:"required"`
	NetworkName      string `json:"networkName" validate:"required"`
	Date             int64  `json:"date" validate:"required"`
}

// Check Upgrade

type GatewayCheckUpgradeRequestParams struct {
	Auth
	GatewayId int   `json:"gatewayId" validate:"required"`
	Date      int64 `json:"date" validate:"required"`
}

// {
// "needUpgrade": 1,
// "firmwareInfo":
// {
// 	"modelNum":"SN227",
// 	"hardwareRevision":"1.1.2",
// 	"firmwareRevision":"1.1.20.1027"
// },
// "version":"1.1.21.1027"
// }

type GatewayCheckUpgradeResponse struct {
	NeedUpgrade  int `json:"needUpgrade"`
	FirmwareInfo struct {
		ModelNum         string `json:"modelNum"`
		HardwareRevision string `json:"hardwareRevision"`
		FirmwareRevision string `json:"firmwareRevision"`
	} `json:"firmwareInfo"`
	Version string `json:"version"`
	Err
}

// Set Upgrade Mode

type GatewaySetUpgradeModeRequestParams struct {
	Auth
	GatewayId int   `json:"gatewayId" validate:"required"`
	Date      int64 `json:"date" validate:"required"`
}
