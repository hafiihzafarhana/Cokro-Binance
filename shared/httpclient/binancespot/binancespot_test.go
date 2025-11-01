package binancespot

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)



func TestNewBinanceSpotHttpClient(t *testing.T) {
	client := NewBinanceSpotHttpClient()

	assert.NotNil(t, client)
	assert.NotNil(t, client.Client)
	assert.Equal(t, "https://api.binance.com/api/v3", client.BaseURL)

	httpClient, ok := client.Client.(*http.Client)
	assert.True(t, ok, "Client harus bertipe *http.Client")
	assert.Equal(t, 10*time.Second, httpClient.Timeout)
}

func TestHttpClient_Do_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/test-endpoint", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}))
	defer server.Close()

	client := &HttpClient{
		Client:  &http.Client{},
		BaseURL: server.URL,
	}

	ctx := context.Background()
	body, err := client.Do(ctx, http.MethodGet, "/test-endpoint", nil)

	require.NoError(t, err)
	assert.JSONEq(t, `{"status":"ok"}`, string(body))
}

func TestHttpClient_Do_BinanceError(t *testing.T) {
	// Mock server that return 400 Bad Request
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "bad request", http.StatusBadRequest)
	}))
	defer server.Close()

	client := &HttpClient{
		Client:  &http.Client{},
		BaseURL: server.URL,
	}

	ctx := context.Background()
	body, err := client.Do(ctx, http.MethodGet, "/bad", nil)

	assert.Error(t, err)
	assert.Nil(t, body)
	assert.True(t, strings.Contains(err.Error(), "bad request"))
}


func TestHttpClient_Do_RequestCreationError(t *testing.T) {
	client := &HttpClient{
		Client:  &http.Client{},
		BaseURL: string([]byte{0x7f}), // invalid URL
	}

	ctx := context.Background()
	body, err := client.Do(ctx, http.MethodGet, "/test", nil)

	assert.Error(t, err)
	assert.Nil(t, body)
	assert.True(t, strings.Contains(err.Error(), "failed to create request"))
}

type errorReader struct{}

func (e *errorReader) Read(p []byte) (int, error) {
	return 0, errors.New("read error")
}
func (e *errorReader) Close() error { return nil }

type mockHttpClient struct {
	resp *http.Response
	err  error
}

// mockHttpClient implement HTTPDoer interface
func (m mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return m.resp, m.err
}

func TestHttpClient_Do_ReadBodyError(t *testing.T) {
	// Reader that always error each read
	errReader := io.NopCloser(&errorReader{})

	mock := mockHttpClient{
		resp: &http.Response{
			StatusCode: http.StatusOK,
			Body:       errReader,
		},
		err: nil,
	}

	client := &HttpClient{
		Client:  mock,                      
		BaseURL: "http://fake-binance.test",
	}

	ctx := context.Background()
	body, err := client.Do(ctx, http.MethodGet, "/test", nil)

	assert.Error(t, err)
	assert.Nil(t, body)
	assert.Contains(t, err.Error(), "failed to read response")
}


func TestHttpClient_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"result":"ok"}`))
	}))
	defer server.Close()

	client := &HttpClient{
		Client:  &http.Client{},
		BaseURL: server.URL,
	}

	ctx := context.Background()
	body, err := client.Get(ctx, "/ping")

	require.NoError(t, err)
	assert.JSONEq(t, `{"result":"ok"}`, string(body))
}

func TestHttpClient_Do_SendRequestError(t *testing.T) {
	mock := mockHttpClient{
		resp: nil,
		err:  errors.New("network down"), // fail simulation
	}

	client := &HttpClient{
		Client:  mock,
		BaseURL: "http://fake-binance.test",
	}

	ctx := context.Background()
	body, err := client.Do(ctx, http.MethodGet, "/any", nil)

	assert.Error(t, err)
	assert.Nil(t, body)
	assert.Contains(t, err.Error(), "failed to send request")
}
