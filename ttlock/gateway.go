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

type Gateway struct{}

func (g *Gateway) List(request *ttt.GatewayListRequestParams) (*ttt.GatewayListResponse, error) {
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
	q.Add("orderBy", strconv.Itoa(request.OrderBy))

	payload := q.Encode()
	urlPath := "/gateway/list?" + payload

	bytes, code, err := TTLockV3Client.GetWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.GatewayListResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if response.ErrCode != 0 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	return &response, nil
}

func (g *Gateway) Delete(request *ttt.GatewayDeleteRequestParams) (*ttt.Err, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("gatewayId", strconv.Itoa(request.GatewayId))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/gateway/delete"

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

func (g *Gateway) Rename(request *ttt.GatewayRenameRequestParams) (*ttt.Err, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("gatewayId", strconv.Itoa(request.GatewayId))
	q.Add("gatewayName", request.GatewayName)
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/gateway/rename"

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

func (g *Gateway) ListByLock(request *ttt.GatewayListByLockRequestParams) (*ttt.GatewayListByLockResponse, error) {
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
	urlPath := "/gateway/listByLock?" + payload

	bytes, code, err := TTLockV3Client.GetWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.GatewayListByLockResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if response.ErrCode != 0 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	return &response, nil
}

func (g *Gateway) ListLock(request *ttt.GatewayListLockRequestParams) (*ttt.GatewayListLockResponse, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}

	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("gatewayId", strconv.Itoa(request.GatewayId))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "/gateway/listLock?" + payload

	bytes, code, err := TTLockV3Client.GetWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.GatewayListLockResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if response.ErrCode != 0 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	return &response, nil
}
