package ttlock

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	ttt "github.com/AVVKavvk/ttlock/ttlock/types"
	"github.com/AVVKavvk/ttlock/utils/logger"
)

type Token struct {
}

func (t *Token) GetAccessToken(request *ttt.AccessTokenRequestParams) (*ttt.AccessTokenResponse, error) {

	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}
	q.Add("clientId", request.ClientId)
	q.Add("clientSecret", request.ClientSecret)
	q.Add("username", request.Username)
	q.Add("password", GenerateMD5HashPassword(request.Password))

	payload := q.Encode()
	urlPath := API_BASE + "/oauth2/token"

	bytes, code, err := TTLockClient.PostWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(payload))

	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.AccessTokenResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if response.ErrCode != 0 {
		return nil, fmt.Errorf("TTLock API Error %d: %s", response.ErrCode, response.ErrMsg)
	}

	return &response, nil

}

func (t *Token) RefreshAccessToken(request *ttt.RefreshAccessTokenRequestParams) (*ttt.RefreshAccessTokenResponse, error) {

	if err := Validate.Struct(request); err != nil {
		log.Errorxf(&logger.XFields{}, "validation failed: %v", err)
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	q := url.Values{}
	q.Add("clientId", request.ClientId)
	q.Add("clientSecret", request.ClientSecret)
	q.Add("grant_type", request.GrantType)
	q.Add("refresh_token", request.RefreshToken)

	payload := q.Encode()
	urlPath := API_BASE + "/oauth2/token"

	bytes, code, err := TTLockClient.PostWithCtx(context.Background(), urlPath, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, []byte(payload))

	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("API request failed with status code: %d, body: %s", code, string(bytes))
	}

	var response ttt.RefreshAccessTokenResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if response.ErrCode != 0 {
		return nil, fmt.Errorf("TTLock API Error %d: %s", response.ErrCode, response.ErrMsg)
	}

	return &response, nil
}
