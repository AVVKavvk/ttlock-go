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

type EKeys struct{}

func (e *EKeys) Send(request *ttt.EKeySendRequestParams) (*ttt.EKeySendResponse, error) {

	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("receiverUsername", request.ReceiverUsername)
	q.Add("keyName", request.KeyName)
	q.Add("startDate", strconv.FormatInt(request.StartDate, 10))
	q.Add("endDate", strconv.FormatInt(request.EndDate, 10))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	if request.CreateUser != 0 {
		q.Add("createUser", strconv.Itoa(request.CreateUser))
	}
	if request.KeyRight != 0 {
		q.Add("keyRight", strconv.Itoa(request.KeyRight))
	}
	if request.RemoteEnable != 0 {
		q.Add("remoteEnable", strconv.Itoa(request.RemoteEnable))
	}
	if request.Remarks != "" {
		q.Add("remarks", request.Remarks)
	}

	payload := q.Encode()
	urlPath := "/key/send"

	bytes, code, err := TTLockV3Client.PostWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(payload))
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.EKeySendResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if response.ErrCode != 0 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}
	return &response, nil
}

func (e *EKeys) List(request *ttt.EKeyListRequestParams) (*ttt.EKeyListResponse, error) {
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
	urlPath := "/key/list"

	bytes, code, err := TTLockV3Client.PostWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(payload))
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.EKeyListResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if response.ErrCode != 0 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}
	return &response, nil
}

func (e *EKeys) GetOne(request *ttt.EKeyGetOneRequestParams) (*ttt.EKeyGetOneResponse, error) {
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
	urlPath := "/key/get?" + payload

	bytes, code, err := TTLockV3Client.GetWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.EKeyGetOneResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if response.ErrCode != 0 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}
	return &response, nil
}

func (e *EKeys) Delete(request *ttt.EKeyDeleteRequestParams) (*ttt.Err, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("keyId", strconv.Itoa(request.KeyId))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/key/delete"

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
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}
	return &response, nil
}

func (e *EKeys) Freeze(request *ttt.EKeyFreezeUnfreezeRequestParams) (*ttt.Err, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("keyId", strconv.Itoa(request.KeyId))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/key/freeze"

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
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}
	return &response, nil
}

func (e *EKeys) Unfreeze(request *ttt.EKeyFreezeUnfreezeRequestParams) (*ttt.Err, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("keyId", strconv.Itoa(request.KeyId))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/key/unfreeze"

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
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}
	return &response, nil
}

func (e *EKeys) Update(request *ttt.EKeyUpdateRequestParams) (*ttt.Err, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("keyId", strconv.Itoa(request.KeyId))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	if request.KeyName != "" {
		q.Add("keyName", request.KeyName)
	}

	if request.RemoteEnable != 0 {
		q.Add("remoteEnable", strconv.Itoa(request.RemoteEnable))
	}

	payload := q.Encode()
	urlPath := "/key/update"

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
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}
	return &response, nil
}

func (e *EKeys) ChangeValidTime(request *ttt.EKeyChangeValidTimeRequestParams) (*ttt.Err, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("keyId", strconv.Itoa(request.KeyId))
	q.Add("startDate", strconv.FormatInt(request.StartDate, 10))
	q.Add("endDate", strconv.FormatInt(request.EndDate, 10))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/key/changePeriod"

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
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}
	return &response, nil
}
