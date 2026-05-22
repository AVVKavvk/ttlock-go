package httpClient

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
}

func New(timeout time.Duration) *Client {
	return &Client{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

// Core executor
func (c *Client) doRequest(ctx context.Context, method, url string, headers map[string]string, body []byte) ([]byte, int, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create request: %w", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to read response: %w", err)
	}

	return respBody, resp.StatusCode, nil
}

// ----- GET -----
func (c *Client) Get(url string, headers map[string]string) ([]byte, int, error) {
	return c.GetWithCtx(context.Background(), url, headers)
}

func (c *Client) GetWithCtx(ctx context.Context, url string, headers map[string]string) ([]byte, int, error) {
	return c.doRequest(ctx, http.MethodGet, url, headers, nil)
}

// ----- POST -----
func (c *Client) Post(url string, headers map[string]string, body []byte) ([]byte, int, error) {
	return c.PostWithCtx(context.Background(), url, headers, body)
}

func (c *Client) PostWithCtx(ctx context.Context, url string, headers map[string]string, body []byte) ([]byte, int, error) {
	return c.doRequest(ctx, http.MethodPost, url, headers, body)
}

// ----- PUT -----
func (c *Client) Put(url string, headers map[string]string, body []byte) ([]byte, int, error) {
	return c.PutWithCtx(context.Background(), url, headers, body)
}

func (c *Client) PutWithCtx(ctx context.Context, url string, headers map[string]string, body []byte) ([]byte, int, error) {
	return c.doRequest(ctx, http.MethodPut, url, headers, body)
}

// ----- DELETE -----
func (c *Client) Delete(url string, headers map[string]string) ([]byte, int, error) {
	return c.DeleteWithCtx(context.Background(), url, headers)
}

func (c *Client) DeleteWithCtx(ctx context.Context, url string, headers map[string]string) ([]byte, int, error) {
	return c.doRequest(ctx, http.MethodDelete, url, headers, nil)
}

// ----- PATCH -----
func (c *Client) Patch(url string, headers map[string]string, body []byte) ([]byte, int, error) {
	return c.PatchWithCtx(context.Background(), url, headers, body)
}

func (c *Client) PatchWithCtx(ctx context.Context, url string, headers map[string]string, body []byte) ([]byte, int, error) {
	return c.doRequest(ctx, http.MethodPatch, url, headers, body)
}
