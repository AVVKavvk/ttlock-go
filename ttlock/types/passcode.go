package ttlock

type PasscodeType int

const (
	PasscodeOneTime         PasscodeType = 1
	PasscodePermanent       PasscodeType = 2
	PasscodePeriod          PasscodeType = 3
	PasscodeDelete          PasscodeType = 4
	PasscodeWeekendCyclic   PasscodeType = 5
	PasscodeDailyCyclic     PasscodeType = 6
	PasscodeWorkdayCyclic   PasscodeType = 7
	PasscodeMondayCyclic    PasscodeType = 8
	PasscodeTuesdayCyclic   PasscodeType = 9
	PasscodeWednesdayCyclic PasscodeType = 10
	PasscodeThursdayCyclic  PasscodeType = 11
	PasscodeFridayCyclic    PasscodeType = 12
	PasscodeSaturdayCyclic  PasscodeType = 13
	PasscodeSundayCyclic    PasscodeType = 14
)

var PasscodeTypeMap = map[PasscodeType]string{
	PasscodeOneTime:         "One-time",
	PasscodePermanent:       "Permanent",
	PasscodePeriod:          "Period",
	PasscodeDelete:          "Delete",
	PasscodeWeekendCyclic:   "Weekend Cyclic",
	PasscodeDailyCyclic:     "Daily Cyclic",
	PasscodeWorkdayCyclic:   "Workday Cyclic",
	PasscodeMondayCyclic:    "Monday Cyclic",
	PasscodeTuesdayCyclic:   "Tuesday Cyclic",
	PasscodeWednesdayCyclic: "Wednesday Cyclic",
	PasscodeThursdayCyclic:  "Thursday Cyclic",
	PasscodeFridayCyclic:    "Friday Cyclic",
	PasscodeSaturdayCyclic:  "Saturday Cyclic",
	PasscodeSundayCyclic:    "Sunday Cyclic",
}

type PasscodeTypeResponse struct {
	Type        int    `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// {
// "keyboardPwdId": 3234293,
// "lockId": 532323,
// "keyboardPwd": "034234",
// "keyboardPwdName":"Passcode for Jack",
// "keyboardPwdType": "3",
// "startDate": 1528878944000,
// "endDate": 1628878944000,
// "sendDate": 1528878944000,
// "senderUsername": "alexa@google.com"
// }

type PasscodeObject struct {
	KeyboardPwdId   int    `json:"keyboardPwdId"`
	LockId          int    `json:"lockId"`
	KeyboardPwd     string `json:"keyboardPwd"`
	KeyboardPwdName string `json:"keyboardPwdName"`
	KeyboardPwdType string `json:"keyboardPwdType"`
	StartDate       int64  `json:"startDate"`
	EndDate         int64  `json:"endDate"`
	SendDate        int64  `json:"sendDate"`
	SenderUsername  string `json:"senderUsername"`
}

type PasscodeListRequestParams struct {
	Auth
	LockId    int    `json:"lockId" validate:"required"`
	Date      int64  `json:"date" validate:"required"`
	OrderBy   int    `json:"orderBy" validate:"required,oneof=1 2"`
	PageNo    int    `json:"pageNo" validate:"required,min=1"`
	PageSize  int    `json:"pageSize" validate:"required,min=1,max=200"`
	SearchStr string `json:"searchStr"`
}

type PasscodeListResponse struct {
	List     []PasscodeObject `json:"list"`
	PageNo   int              `json:"pageNo"`
	PageSize int              `json:"pageSize"`
	Total    int              `json:"total"`
	Pages    int              `json:"pages"`
	Err
}

// Random Passcode

type PasscodeRandomRequestParams struct {
	Auth
	LockId          int    `json:"lockId" validate:"required"`
	KeyboardPwdType int    `json:"keyboardPwdType" validate:"required"`
	StartDate       int64  `json:"startDate" validate:"required"`
	EndDate         int64  `json:"endDate" validate:"required"`
	Date            int64  `json:"date" validate:"required"`
	KeyboardPwdName string `json:"keyboardPwdName"`
}

type PasscodeRandomResponse struct {
	KeyboardPwd   string `json:"keyboardPwd"`
	KeyboardPwdId int    `json:"keyboardPwdId"`
	Err
}

// Custom Passcode

type PasscodeCustomRequestParams struct {
	Auth
	LockId          int    `json:"lockId" validate:"required"`
	KeyboardPwd     int    `json:"keyboardPwd" validate:"required"`
	KeyboardPwdType int    `json:"keyboardPwdType" validate:"required,oneof=2 3"`
	StartDate       int64  `json:"startDate" validate:"required"`
	EndDate         int64  `json:"endDate" validate:"required"`
	AddType         int    `json:"addType" validate:"required,oneof=1 2"`
	Date            int64  `json:"date" validate:"required"`
	KeyboardPwdName string `json:"keyboardPwdName"`
}

type PasscodeCustomResponse struct {
	KeyboardPwdId int `json:"keyboardPwdId"`
	Err
}

// Delete Passcode

type PasscodeDeleteRequestParams struct {
	Auth
	LockId        int   `json:"lockId" validate:"required"`
	KeyboardPwdId int   `json:"keyboardPwdId" validate:"required"`
	DeleteType    int   `json:"deleteType" validate:"required,oneof=1 2"`
	Date          int64 `json:"date" validate:"required"`
}

// Update Passcode

type PasscodeUpdateRequestParams struct {
	Auth
	LockId          int    `json:"lockId" validate:"required"`
	KeyboardPwdId   int    `json:"keyboardPwdId" validate:"required"`
	KeyboardPwdName string `json:"keyboardPwdName"`
	NewKeyboardPwd  string `json:"newKeyboardPwd"`
	StartDate       int64  `json:"startDate"`
	EndDate         int64  `json:"endDate"`
	ChangeType      int    `json:"changeType" validate:"required,oneof=1 2"`
	Date            int64  `json:"date" validate:"required"`
}
