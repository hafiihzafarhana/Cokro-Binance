package binancespot

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	defaultBinanceSpotBaseURL = "https://api.binance.com/api/v3"
	defaultTimeout            = 10 * time.Second
)

type HttpClient struct {
	Client  *http.Client
	BaseURL string
}

func NewBinanceSpotHttpClient() *HttpClient {
	return &HttpClient{
		Client:  &http.Client{Timeout: defaultTimeout},
		BaseURL: defaultBinanceSpotBaseURL,
	}
}

func (c *HttpClient) Do(ctx context.Context, method, endpoint string, body io.Reader) ([]byte, error) {
	u := strings.TrimRight(c.BaseURL, "/") + "/" + strings.TrimLeft(endpoint, "/")
	fmt.Println(u)
	fmt.Println(endpoint)
	req, err := http.NewRequestWithContext(ctx, method, u, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("binance error %d: %s", resp.StatusCode, string(b))
	}

	return b, nil
}

func (c *HttpClient) Get(ctx context.Context, endpoint string) ([]byte, error) {
	return c.Do(ctx, http.MethodGet, endpoint, nil)
}
