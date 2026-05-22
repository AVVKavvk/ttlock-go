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

type Users struct {
}

func (u *Users) List(request *ttt.UserListRequestParams) (*ttt.UserListResponse, error) {

	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// url.Values for safe, proper URL encoding and integer conversion
	q := url.Values{}
	q.Add("clientId", request.ClientId)
	q.Add("clientSecret", request.ClientSecret)
	q.Add("pageNo", strconv.Itoa(request.PageNo))
	q.Add("pageSize", strconv.Itoa(request.PageSize))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	if request.StartDate != 0 {
		q.Add("startDate", strconv.FormatInt(request.StartDate, 10))
	}
	if request.EndDate != 0 {
		q.Add("endDate", strconv.FormatInt(request.EndDate, 10))
	}

	// Encode() automatically formats everything as key=value&key2=value2
	urlPath := "/user/list?" + q.Encode()

	// Make the HTTP request
	bytes, code, err := TTLockV3Client.GetWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}

	// Raise for stats code
	if code < 200 || code >= 300 {
		// Better to return an error than just nil, nil so the caller knows what failed
		return nil, fmt.Errorf("API request failed with status code: %d", code)
	}

	// Unmarshaling the response
	var response ttt.UserListResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}

func (u *Users) Register(request *ttt.UserRegisterRequestParams) (*ttt.UserRegisterResponse, error) {

	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}
	q.Add("clientId", request.ClientId)
	q.Add("clientSecret", request.ClientSecret)
	q.Add("username", request.Username)
	q.Add("password", GenerateMD5HashPassword(request.Password)) // 32-char lowercase MD5 hash
	q.Add("date", strconv.FormatInt(request.Date, 10))
	payload := q.Encode()
	urlPath := "/user/register"

	bytes, code, err := TTLockV3Client.PostWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(payload))

	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.UserRegisterResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if response.ErrCode != 0 {
		return nil, fmt.Errorf("TTLock API Error %d: %s", response.ErrCode, response.ErrMsg)
	}

	return &response, nil

}

func (u *Users) PasswordReset(request *ttt.UserPasswordResetRequestParams) (*ttt.UserPasswordResetResponse, error) {

	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}
	q.Add("clientId", request.ClientId)
	q.Add("clientSecret", request.ClientSecret)
	q.Add("username", request.Username)
	q.Add("password", GenerateMD5HashPassword(request.Password)) // 32-char lowercase MD5 hash
	q.Add("date", strconv.FormatInt(request.Date, 10))
	payload := q.Encode()
	urlPath := "/user/resetPassword"

	bytes, code, err := TTLockV3Client.PostWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(payload))

	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.UserPasswordResetResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}

func (u *Users) Delete(request *ttt.UserDeleteRequestParams) (*ttt.UserDeleteResponse, error) {

	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}
	q.Add("clientId", request.ClientId)
	q.Add("clientSecret", request.ClientSecret)
	q.Add("username", request.Username)
	q.Add("date", strconv.FormatInt(request.Date, 10))
	payload := q.Encode()
	urlPath := "/user/delete"

	bytes, code, err := TTLockV3Client.PostWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(payload))

	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.UserDeleteResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}
