package ttlock

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"time"

	httpClient "github.com/AVVKavvk/ttlock/utils/http"
	"github.com/go-playground/validator/v10"
)

const (
	API_BASE_V3 = "https://euapi.ttlock.com/v3"
	API_BASE    = "https://euapi.ttlock.com"
)

var (
	TTLockV3Client *TTLockV3ClientConfig
	TTLockClient   *httpClient.Client
	Validate       = validator.New()
)

type TTLockV3ClientConfig struct {
	_client *httpClient.Client
}

// ----- GET -----
func (c *TTLockV3ClientConfig) Get(url string, headers map[string]string) ([]byte, int, error) {
	return c.GetWithCtx(context.Background(), url, headers)
}

func (c *TTLockV3ClientConfig) GetWithCtx(ctx context.Context, url string, headers map[string]string) ([]byte, int, error) {

	finalUrl := API_BASE_V3 + url

	return c._client.GetWithCtx(ctx, finalUrl, headers)
}

// ----- POST -----
func (c *TTLockV3ClientConfig) Post(url string, headers map[string]string, body []byte) ([]byte, int, error) {
	return c.PostWithCtx(context.Background(), url, headers, body)
}

func (c *TTLockV3ClientConfig) PostWithCtx(ctx context.Context, url string, headers map[string]string, body []byte) ([]byte, int, error) {
	finalUrl := API_BASE_V3 + url
	return c._client.PostWithCtx(ctx, finalUrl, headers, body)
}

// ----- PUT -----
func (c *TTLockV3ClientConfig) Put(url string, headers map[string]string, body []byte) ([]byte, int, error) {
	return c.PutWithCtx(context.Background(), url, headers, body)
}

func (c *TTLockV3ClientConfig) PutWithCtx(ctx context.Context, url string, headers map[string]string, body []byte) ([]byte, int, error) {
	finalUrl := API_BASE_V3 + url
	return c._client.PutWithCtx(ctx, finalUrl, headers, body)
}

// ----- DELETE -----
func (c *TTLockV3ClientConfig) Delete(url string, headers map[string]string) ([]byte, int, error) {
	return c.DeleteWithCtx(context.Background(), url, headers)
}

func (c *TTLockV3ClientConfig) DeleteWithCtx(ctx context.Context, url string, headers map[string]string) ([]byte, int, error) {
	finalUrl := API_BASE_V3 + url
	return c._client.DeleteWithCtx(ctx, finalUrl, headers)
}

// ----- PATCH -----
func (c *TTLockV3ClientConfig) Patch(url string, headers map[string]string, body []byte) ([]byte, int, error) {
	return c.PatchWithCtx(context.Background(), url, headers, body)
}

func (c *TTLockV3ClientConfig) PatchWithCtx(ctx context.Context, url string, headers map[string]string, body []byte) ([]byte, int, error) {
	finalUrl := API_BASE_V3 + url
	return c._client.PatchWithCtx(ctx, finalUrl, headers, body)
}

func GenerateMD5HashPassword(text string) string {
	// Calculate the MD5 hash
	hash := md5.Sum([]byte(text))

	// Convert the byte array to a hex string
	// hex.EncodeToString automatically uses lowercase letters
	return hex.EncodeToString(hash[:])
}

func init() {
	TTLockV3Client = &TTLockV3ClientConfig{
		_client: httpClient.New(60 * time.Second),
	}
	TTLockClient = httpClient.New(60 * time.Second)
}
