package trading212

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Environment represents the Trading 212 API environment
type Environment string

const (
	// Demo environment for paper trading
	Demo Environment = "https://demo.trading212.com"
	// Live environment for real money trading
	Live Environment = "https://live.trading212.com"
)

// Client represents the Trading 212 API client
type Client struct {
	baseURL    string
	apiKey     string
	apiSecret  string
	httpClient *http.Client
}

// NewClient creates a new Trading 212 API client
func NewClient(env Environment, apiKey, apiSecret string) *Client {
	return &Client{
		baseURL:    string(env),
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

// SetHTTPClient allows setting a custom HTTP client
func (c *Client) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}

// makeRequest performs an HTTP request with authentication
func (c *Client) makeRequest(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set authentication header
	credentials := base64.StdEncoding.EncodeToString([]byte(c.apiKey + ":" + c.apiSecret))
	req.Header.Set("Authorization", "Basic "+credentials)
	
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return resp, nil
}

// handleResponse processes the HTTP response and unmarshals JSON
func (c *Client) handleResponse(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// buildQuery builds URL query parameters
func buildQuery(params map[string]interface{}) string {
	if len(params) == 0 {
		return ""
	}

	values := url.Values{}
	for key, value := range params {
		if value != nil {
			switch v := value.(type) {
			case string:
				if v != "" {
					values.Add(key, v)
				}
			case int:
				values.Add(key, strconv.Itoa(v))
			case int64:
				values.Add(key, strconv.FormatInt(v, 10))
			case float64:
				values.Add(key, strconv.FormatFloat(v, 'f', -1, 64))
			case bool:
				values.Add(key, strconv.FormatBool(v))
			case time.Time:
				values.Add(key, v.Format(time.RFC3339))
			}
		}
	}

	if len(values) > 0 {
		return "?" + values.Encode()
	}
	return ""
}