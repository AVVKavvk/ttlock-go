package ttlock

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	ttt "github.com/AVVKavvk/ttlock/ttlock/types"
	"github.com/AVVKavvk/ttlock/utils/logger"
)

type Lock struct {
}

func (l *Lock) Initialize(request ttt.LockInitializeRequest) (*ttt.LockInitializeResponse, error) {

	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockData", request.LockData)
	q.Add("date", strconv.FormatInt(request.Date, 10))

	if request.GroupId != 0 {
		q.Add("groupId", strconv.Itoa(request.GroupId))
	}

	if request.LockAlias != "" {
		q.Add("lockAlias", request.LockAlias)
	}

	if request.NBinitSuccess != 0 {
		q.Add("nbInitSuccess", strconv.Itoa(request.NBinitSuccess))
	}

	payload := q.Encode()
	urlPath := "/lock/initialize"

	bytes, code, err := TTLockV3Client.PostWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(payload))

	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.LockInitializeResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response.ErrCode != 0 {
		return nil, fmt.Errorf("failed to initialize lock: %s", response.ErrMsg)
	}

	return &response, nil
}

func (l *Lock) List(request *ttt.LockListRequestParams) (*ttt.LockListResponse, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("pageNo", strconv.Itoa(request.PageNo))
	q.Add("pageSize", strconv.Itoa(request.PageSize))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	if request.LockAlias != "" {
		q.Add("lockAlias", request.LockAlias)
	}

	if request.GroupId != 0 {
		q.Add("groupId", strconv.Itoa(request.GroupId))
	}

	payload := q.Encode()
	urlPath := "/lock/list?" + payload

	bytes, code, err := TTLockV3Client.GetWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.LockListResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response.ErrCode != 0 {
		return nil, fmt.Errorf("failed to list locks: %s", response.ErrMsg)
	}
	return &response, nil
}

func (l *Lock) Detail(request *ttt.LockDetailsRequestParams) (*ttt.LockDetailsResponse, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/lock/detail?" + payload

	bytes, code, err := TTLockV3Client.GetWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.LockDetailsResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response.ErrCode != 0 {
		return nil, fmt.Errorf("failed to list locks: %s", response.ErrMsg)
	}
	return &response, nil
}

func (l *Lock) Delete(request *ttt.LockDeleteRequestParams) (*ttt.Err, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/lock/delete"

	bytes, code, err := TTLockV3Client.PostWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(payload))
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.Err
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response.ErrCode != 0 {
		return nil, fmt.Errorf("failed to list locks: %s", response.ErrMsg)
	}
	return &response, nil
}

func (l *Lock) Update(request *ttt.LockUpdateRequestParams) (*ttt.Err, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("lockData", request.LockData)
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/lock/updateLockData"

	bytes, code, err := TTLockV3Client.PostWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(payload))
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.Err
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response.ErrCode != 0 {
		return nil, fmt.Errorf("failed to list locks: %s", response.ErrMsg)
	}
	return &response, nil
}

func (l *Lock) Rename(request *ttt.LockRenameRequestParams) (*ttt.Err, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("lockAlias", request.LockAlias)
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/lock/rename"

	bytes, code, err := TTLockV3Client.PostWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(payload))
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.Err
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response.ErrCode != 0 {
		return nil, fmt.Errorf("failed to list locks: %s", response.ErrMsg)
	}
	return &response, nil
}

func (l *Lock) ChangeAdminPasscode(request *ttt.LockChangeAdminPasscodeRequestParams) (*ttt.Err, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("password", request.Password)
	q.Add("changeType", strconv.Itoa(request.ChangeType))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/lock/changeAdminKeyboardPwd"

	bytes, code, err := TTLockV3Client.PostWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(payload))
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.Err
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response.ErrCode != 0 {
		return nil, fmt.Errorf("failed to list locks: %s", response.ErrMsg)
	}
	return &response, nil
}

func (l *Lock) AutoLockTime(request *ttt.LockAutoLockTimeRequestParams) (*ttt.Err, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("seconds", strconv.Itoa(request.Seconds))
	q.Add("type", strconv.Itoa(request.Type))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/lock/setAutoLockTime"

	bytes, code, err := TTLockV3Client.PostWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(payload))
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.Err
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response.ErrCode != 0 {
		return nil, fmt.Errorf("failed to list locks: %s", response.ErrMsg)
	}
	return &response, nil
}

func (l *Lock) ListEKeys(request *ttt.LockListEKeyRequestParams) (*ttt.LockListEkeyResponse, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("pageNo", strconv.Itoa(request.PageNo))
	q.Add("pageSize", strconv.Itoa(request.PageSize))
	q.Add("orderBy", strconv.Itoa(request.OrderBy))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	if request.SearchStr != "" {
		q.Add("searchStr", request.SearchStr)
	}

	if request.KeyRight != 0 {
		q.Add("keyRight", strconv.Itoa(request.KeyRight))
	}

	payload := q.Encode()
	urlPath := "/lock/listKey?" + payload

	bytes, code, err := TTLockV3Client.GetWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.LockListEkeyResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response.ErrCode != 0 {
		return nil, fmt.Errorf("failed to list locks: %s", response.ErrMsg)
	}
	return &response, nil

}

func (l *Lock) ListPasscodes(request *ttt.PasscodeListRequestParams) (*ttt.PasscodeListResponse, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("pageNo", strconv.Itoa(request.PageNo))
	q.Add("pageSize", strconv.Itoa(request.PageSize))
	q.Add("orderBy", strconv.Itoa(request.OrderBy))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	if request.SearchStr != "" {
		q.Add("searchStr", request.SearchStr)
	}

	payload := q.Encode()
	urlPath := "/lock/listKeyboardPwd?" + payload

	bytes, code, err := TTLockV3Client.GetWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.PasscodeListResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response.ErrCode != 0 {
		return nil, fmt.Errorf("failed to list locks: %s", response.ErrMsg)
	}
	return &response, nil
}

func (l *Lock) Lock(request *ttt.LockLockUnlockRequestParams) (*ttt.Err, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/lock/lock"

	bytes, code, err := TTLockV3Client.PostWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(payload))
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.Err
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if response.ErrCode != 0 {
		return nil, fmt.Errorf("failed to lock lock: %s", response.ErrMsg)
	}
	return &response, nil
}

func (l *Lock) Unlock(request *ttt.LockLockUnlockRequestParams) (*ttt.Err, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/lock/unlock"

	bytes, code, err := TTLockV3Client.PostWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(payload))
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.Err
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if response.ErrCode != 0 {
		return nil, fmt.Errorf("failed to unlock lock: %s", response.ErrMsg)
	}
	return &response, nil
}
func (l *Lock) QueryOpenState(request *ttt.LockQueryOpenStateRequestParams) (*ttt.LockQueryOpenStateResponse, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/lock/queryOpenState?" + payload

	bytes, code, err := TTLockV3Client.GetWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.LockQueryOpenStateResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if response.ErrCode != 0 {
		return nil, fmt.Errorf("failed to unlock lock: %s", response.ErrMsg)
	}
	return &response, nil
}

func (l *Lock) LockTime(request *ttt.LockTimeRequestParams) (*ttt.LockTimeResponse, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/lock/queryDate?" + payload

	bytes, code, err := TTLockV3Client.GetWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.LockTimeResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if response.ErrCode != 0 {
		return nil, fmt.Errorf("failed to unlock lock: %s", response.ErrMsg)
	}
	return &response, nil
}

func (l *Lock) UpdateTime(request *ttt.LockUpdateTimeRequestParams) (*ttt.LockUpdateTimeResponse, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/lock/updateDate"

	bytes, code, err := TTLockV3Client.PostWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(payload))
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.LockUpdateTimeResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if response.ErrCode != 0 {
		return nil, fmt.Errorf("failed to unlock lock: %s", response.ErrMsg)
	}
	return &response, nil
}

func (l *Lock) BatteryStatus(request *ttt.LockQueryBatteryRequestParams) (*ttt.LockQueryBatteryResponse, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/lock/queryElectricQuantity?" + payload

	bytes, code, err := TTLockV3Client.GetWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.LockQueryBatteryResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if response.ErrCode != 0 {
		return nil, fmt.Errorf("failed to unlock lock: %s", response.ErrMsg)
	}
	return &response, nil
}
