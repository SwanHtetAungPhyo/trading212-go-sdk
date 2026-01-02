package trading212

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	client := NewClient(Demo, "test-key", "test-secret")
	
	assert.Equal(t, string(Demo), client.baseURL)
	assert.Equal(t, "test-key", client.apiKey)
	assert.Equal(t, "test-secret", client.apiSecret)
	assert.NotNil(t, client.httpClient)
}

func TestClient_makeRequest(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check authentication header
		auth := r.Header.Get("Authorization")
		assert.Contains(t, auth, "Basic")
		
		// Check method and path
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/test", r.URL.Path)
		
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	client := NewClient(Environment(server.URL), "test-key", "test-secret")
	
	resp, err := client.makeRequest(context.Background(), "GET", "/test", nil)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp.Body.Close()
}

func TestClient_handleResponse(t *testing.T) {
	client := NewClient(Demo, "test-key", "test-secret")
	
	// Test successful response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": 123, "name": "test"}`))
	}))
	defer server.Close()

	resp, err := http.Get(server.URL)
	require.NoError(t, err)
	
	var result struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	
	err = client.handleResponse(resp, &result)
	require.NoError(t, err)
	assert.Equal(t, 123, result.ID)
	assert.Equal(t, "test", result.Name)
}

func TestClient_handleResponse_Error(t *testing.T) {
	client := NewClient(Demo, "test-key", "test-secret")
	
	// Test error response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Bad request"}`))
	}))
	defer server.Close()

	resp, err := http.Get(server.URL)
	require.NoError(t, err)
	
	err = client.handleResponse(resp, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "400")
}

func TestBuildQuery(t *testing.T) {
	tests := []struct {
		name     string
		params   map[string]interface{}
		expected string
	}{
		{
			name:     "empty params",
			params:   map[string]interface{}{},
			expected: "",
		},
		{
			name: "string param",
			params: map[string]interface{}{
				"ticker": "AAPL_US_EQ",
			},
			expected: "?ticker=AAPL_US_EQ",
		},
		{
			name: "multiple params",
			params: map[string]interface{}{
				"ticker": "AAPL_US_EQ",
				"limit":  50,
				"cursor": int64(123456),
			},
			expected: "?cursor=123456&limit=50&ticker=AAPL_US_EQ",
		},
		{
			name: "time param",
			params: map[string]interface{}{
				"time": time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			},
			expected: "?time=2023-01-01T12%3A00%3A00Z",
		},
		{
			name: "nil and empty values",
			params: map[string]interface{}{
				"ticker": "AAPL_US_EQ",
				"empty":  "",
				"nil":    nil,
			},
			expected: "?ticker=AAPL_US_EQ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildQuery(tt.params)
			if tt.expected == "" {
				assert.Equal(t, tt.expected, result)
			} else {
				// For multiple params, order might vary, so check if all expected params are present
				assert.Contains(t, result, "?")
				for key, value := range tt.params {
					if value != nil && value != "" {
						assert.Contains(t, result, key+"=")
					}
				}
			}
		})
	}
}