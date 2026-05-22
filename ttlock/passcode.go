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

type Passcode struct {
}

func (p *Passcode) Types() ([]ttt.PasscodeTypeResponse, error) {

	types := []ttt.PasscodeTypeResponse{
		{
			Type:        1,
			Name:        "One-time",
			Description: "This code only valid for once within 6 hours from the Start Time",
		},
		{
			Type:        2,
			Name:        "Permanent",
			Description: "This code must be used at least once within 24 Hours after the Start Time, or it will be invalidated",
		},
		{
			Type:        3,
			Name:        "Period",
			Description: "This code must be used at least once within 24 Hours after the Start Time, or it will be invalidated",
		},
		{
			Type:        4,
			Name:        "Delete",
			Description: "This code will delete all other codes",
		},
		{
			Type:        5,
			Name:        "Weekend Cyclic",
			Description: "This code is valid during the time period at the weekend",
		},
		{
			Type:        6,
			Name:        "Daily Cyclic",
			Description: "This code is valid during the time period everyday",
		},
		{
			Type:        7,
			Name:        "Workday Cyclic",
			Description: "This code is valid during the time period on workdays",
		},
		{
			Type:        8,
			Name:        "Monday Cyclic",
			Description: "This code is valid during the time period on Mondays",
		},
		{
			Type:        9,
			Name:        "Tuesday Cyclic",
			Description: "This code is valid during the time period on Tuesdays",
		},
		{
			Type:        10,
			Name:        "Wednesday Cyclic",
			Description: "This code is valid during the time period on Wednesdays",
		},
		{
			Type:        11,
			Name:        "Thursday Cyclic",
			Description: "This code is valid during the time period on Thursdays",
		},
		{
			Type:        12,
			Name:        "Friday Cyclic",
			Description: "This code is valid during the time period on Fridays",
		},
		{
			Type:        13,
			Name:        "Saturday Cyclic",
			Description: "This code is valid during the time period on Saturdays",
		},
		{
			Type:        14,
			Name:        "Sunday Cyclic",
			Description: "This code is valid during the time period on Sundays",
		},
	}
	return types, nil
}

func (p *Passcode) Random(request *ttt.PasscodeRandomRequestParams) (*ttt.PasscodeRandomResponse, error) {

	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}
	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("keyboardPwdType", strconv.Itoa(request.KeyboardPwdType))
	q.Add("startDate", strconv.FormatInt(request.StartDate, 10))
	q.Add("endDate", strconv.FormatInt(request.EndDate, 10))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	if request.KeyboardPwdName != "" {
		q.Add("keyboardPwdName", request.KeyboardPwdName)
	}

	payload := q.Encode()
	urlPath := "keyboardPwd/get"

	bytes, code, err := TTLockV3Client.PostWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(payload))
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.PasscodeRandomResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response.ErrCode != 0 {
		return nil, fmt.Errorf("API request failed with code: %d, body: %s", response.ErrCode, response.ErrMsg)
	}
	return &response, nil
}

func (p *Passcode) Custom(request *ttt.PasscodeCustomRequestParams) (*ttt.PasscodeCustomResponse, error) {

	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}
	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("keyboardPwd", strconv.Itoa(request.KeyboardPwd))
	q.Add("keyboardPwdType", strconv.Itoa(request.KeyboardPwdType))
	q.Add("startDate", strconv.FormatInt(request.StartDate, 10))
	q.Add("endDate", strconv.FormatInt(request.EndDate, 10))
	q.Add("addType", strconv.Itoa(request.AddType))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	if request.KeyboardPwdName != "" {
		q.Add("keyboardPwdName", request.KeyboardPwdName)
	}

	payload := q.Encode()
	urlPath := "keyboardPwd/add"

	bytes, code, err := TTLockV3Client.PostWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(payload))
	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {

		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.PasscodeCustomResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response.ErrCode != 0 {
		return nil, fmt.Errorf("API request failed with code: %d, body: %s", response.ErrCode, response.ErrMsg)
	}
	return &response, nil
}

func (p *Passcode) Delete(request *ttt.PasscodeDeleteRequestParams) (*ttt.Err, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}
	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("keyboardPwdId", strconv.Itoa(request.KeyboardPwdId))
	q.Add("deleteType", strconv.Itoa(request.DeleteType))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	payload := q.Encode()
	urlPath := "keyboardPwd/delete"

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
		return nil, fmt.Errorf("API request failed with code: %d, body: %s", response.ErrCode, response.ErrMsg)
	}
	return &response, nil
}

func (p *Passcode) Update(request *ttt.PasscodeUpdateRequestParams) (*ttt.Err, error) {
	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}
	q.Add("clientId", request.ClientId)
	q.Add("accessToken", request.AccessToken)
	q.Add("lockId", strconv.Itoa(request.LockId))
	q.Add("keyboardPwdId", strconv.Itoa(request.KeyboardPwdId))
	q.Add("date", strconv.FormatInt(request.Date, 10))

	isChanged := false
	if request.KeyboardPwdName != "" {
		q.Add("keyboardPwdName", request.KeyboardPwdName)
	}
	if request.StartDate != 0 {
		if request.EndDate != 0 {
			q.Add("startDate", strconv.FormatInt(request.StartDate, 10))
			q.Add("endDate", strconv.FormatInt(request.EndDate, 10))
			isChanged = true
		} else {
			return nil, fmt.Errorf("validation failed: endDate is required")
		}
	}
	if request.KeyboardPwdName != "" {
		q.Add("keyboardPwdName", request.KeyboardPwdName)
		isChanged = true
	}

	if isChanged {
		q.Add("changeType", strconv.Itoa(request.ChangeType))
	}

	payload := q.Encode()
	urlPath := "keyboardPwd/change"

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
		return nil, fmt.Errorf("API request failed with code: %d, body: %s", response.ErrCode, response.ErrMsg)
	}
	return &response, nil
}
