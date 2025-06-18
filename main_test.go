package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		url            string
		body           string
		contentType    string
		expectedStatus int
		expectedBody   Response
	}{
		{
			name:           "GET request returns hello world",
			method:         http.MethodGet,
			url:            "/hello",
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Message: "Hello, World!"},
		},
		{
			name:           "GET request with name parameter",
			method:         http.MethodGet,
			url:            "/hello?name=Justin",
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Message: "Hello, Justin!"},
		},
		{
			name:           "GET request with empty name parameter",
			method:         http.MethodGet,
			url:            "/hello?name=",
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Message: "Hello, World!"},
		},
		{
			name:           "GET request with special characters in name",
			method:         http.MethodGet,
			url:            "/hello?name=John%20Doe",
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Message: "Hello, John Doe!"},
		},
		{
			name:           "POST request with JSON returns hello world",
			method:         http.MethodPost,
			url:            "/hello",
			body:           `{"name":""}`,
			contentType:    "application/json",
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Message: "Hello, World!"},
		},
		{
			name:           "PUT request returns hello world",
			method:         http.MethodPut,
			url:            "/hello",
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Message: "Hello, World!"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			if tt.body != "" {
				req = httptest.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			} else {
				req = httptest.NewRequest(tt.method, tt.url, nil)
			}
			if tt.contentType != "" {
				req.Header.Set("Content-Type", tt.contentType)
			}
			rec := httptest.NewRecorder()

			helloHandler(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			contentType := rec.Header().Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("expected Content-Type 'application/json', got '%s'", contentType)
			}

			var resp Response
			if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if resp.Message != tt.expectedBody.Message {
				t.Errorf("expected message '%s', got '%s'", tt.expectedBody.Message, resp.Message)
			}
		})
	}
}

func TestServerConfiguration(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if server.Addr != ":8080" {
		t.Errorf("expected server address ':8080', got '%s'", server.Addr)
	}

	if server.Handler == nil {
		t.Error("expected server handler to be set")
	}
}

func BenchmarkHelloHandler(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		helloHandler(rec, req)
	}
}

func TestJSONEncoding(t *testing.T) {
	resp := Response{Message: "Test Message"}

	data, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("failed to marshal response: %v", err)
	}

	var decoded Response
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if decoded.Message != resp.Message {
		t.Errorf("expected message '%s', got '%s'", resp.Message, decoded.Message)
	}
}

func TestPOSTHandler(t *testing.T) {
	tests := []struct {
		name           string
		body           interface{}
		contentType    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "POST with valid JSON",
			body:           Request{Name: "Alice"},
			contentType:    "application/json",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Hello, Alice!"}`,
		},
		{
			name:           "POST with empty name",
			body:           Request{Name: ""},
			contentType:    "application/json",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Hello, World!"}`,
		},
		{
			name:           "POST without Content-Type",
			body:           Request{Name: "Bob"},
			contentType:    "",
			expectedStatus: http.StatusUnsupportedMediaType,
			expectedBody:   `{"error":"Content-Type must be application/json","code":"INVALID_CONTENT_TYPE"}`,
		},
		{
			name:           "POST with invalid JSON",
			body:           `{"invalid json"`,
			contentType:    "application/json",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid JSON","code":"INVALID_JSON"}`,
		},
		{
			name:           "POST with extra fields",
			body:           map[string]string{"name": "Charlie", "extra": "field"},
			contentType:    "application/json",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid JSON","code":"INVALID_JSON"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			var err error

			switch v := tt.body.(type) {
			case string:
				body = []byte(v)
			default:
				body, err = json.Marshal(v)
				if err != nil {
					t.Fatalf("failed to marshal request body: %v", err)
				}
			}

			req := httptest.NewRequest(http.MethodPost, "/hello", bytes.NewReader(body))
			if tt.contentType != "" {
				req.Header.Set("Content-Type", tt.contentType)
			}
			rec := httptest.NewRecorder()

			helloHandler(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			respBody := strings.TrimSpace(rec.Body.String())
			expectedBody := strings.TrimSpace(tt.expectedBody)

			if tt.expectedStatus == http.StatusOK {
				var resp Response
				if err := json.Unmarshal([]byte(respBody), &resp); err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}

				var expectedResp Response
				if err := json.Unmarshal([]byte(expectedBody), &expectedResp); err != nil {
					t.Fatalf("failed to unmarshal expected response: %v", err)
				}

				if resp.Message != expectedResp.Message {
					t.Errorf("expected message '%s', got '%s'", expectedResp.Message, resp.Message)
				}
			} else {
				if respBody != expectedBody {
					t.Errorf("expected body '%s', got '%s'", expectedBody, respBody)
				}
			}
		})
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	contentType := rec.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type 'application/json', got '%s'", contentType)
	}

	var resp HealthResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Status != "healthy" {
		t.Errorf("expected status 'healthy', got '%s'", resp.Status)
	}
}

func TestPingHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedStatus int
		expectedPong   string
	}{
		{
			name:           "GET ping returns pong",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
			expectedPong:   "pong",
		},
		{
			name:           "POST ping returns pong",
			method:         http.MethodPost,
			expectedStatus: http.StatusOK,
			expectedPong:   "pong",
		},
		{
			name:           "PUT ping returns pong",
			method:         http.MethodPut,
			expectedStatus: http.StatusOK,
			expectedPong:   "pong",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/ping", nil)
			rec := httptest.NewRecorder()

			pingHandler(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			contentType := rec.Header().Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("expected Content-Type 'application/json', got '%s'", contentType)
			}

			var resp PingResponse
			if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if resp.Pong != tt.expectedPong {
				t.Errorf("expected pong '%s', got '%s'", tt.expectedPong, resp.Pong)
			}
		})
	}
}

func TestInfoHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		url            string
		headers        map[string]string
		expectedStatus int
		validate       func(t *testing.T, resp InfoResponse)
	}{
		{
			name:           "GET info with basic request",
			method:         http.MethodGet,
			url:            "/info",
			expectedStatus: http.StatusOK,
			validate: func(t *testing.T, resp InfoResponse) {
				if resp.Method != http.MethodGet {
					t.Errorf("expected method %s, got %s", http.MethodGet, resp.Method)
				}
				if resp.URL != "/info" {
					t.Errorf("expected URL /info, got %s", resp.URL)
				}
			},
		},
		{
			name:           "GET info with query parameters",
			method:         http.MethodGet,
			url:            "/info?foo=bar&test=123",
			expectedStatus: http.StatusOK,
			validate: func(t *testing.T, resp InfoResponse) {
				if resp.QueryParams["foo"] != "bar" {
					t.Errorf("expected query param foo=bar, got foo=%s", resp.QueryParams["foo"])
				}
				if resp.QueryParams["test"] != "123" {
					t.Errorf("expected query param test=123, got test=%s", resp.QueryParams["test"])
				}
			},
		},
		{
			name:   "POST info with custom headers",
			method: http.MethodPost,
			url:    "/info",
			headers: map[string]string{
				"X-Custom-Header": "custom-value",
				"User-Agent":      "test-agent",
			},
			expectedStatus: http.StatusOK,
			validate: func(t *testing.T, resp InfoResponse) {
				if resp.Method != http.MethodPost {
					t.Errorf("expected method %s, got %s", http.MethodPost, resp.Method)
				}
				if resp.Headers["X-Custom-Header"] != "custom-value" {
					t.Errorf("expected header X-Custom-Header=custom-value, got %s", resp.Headers["X-Custom-Header"])
				}
				if resp.UserAgent != "test-agent" {
					t.Errorf("expected user agent test-agent, got %s", resp.UserAgent)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.url, nil)

			// Add custom headers if provided
			for key, value := range tt.headers {
				req.Header.Set(key, value)
			}

			// Set a default host if not provided
			if req.Host == "" {
				req.Host = "example.com"
			}

			// Set RemoteAddr for testing
			req.RemoteAddr = "192.168.1.1:12345"

			rec := httptest.NewRecorder()

			infoHandler(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			contentType := rec.Header().Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("expected Content-Type 'application/json', got '%s'", contentType)
			}

			var resp InfoResponse
			if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			// Basic validations that apply to all tests
			if resp.Host != req.Host {
				t.Errorf("expected host %s, got %s", req.Host, resp.Host)
			}
			if resp.RemoteAddr != req.RemoteAddr {
				t.Errorf("expected remote addr %s, got %s", req.RemoteAddr, resp.RemoteAddr)
			}

			// Run test-specific validations
			if tt.validate != nil {
				tt.validate(t, resp)
			}
		})
	}
}

func TestLoggingMiddleware(t *testing.T) {
	var logBuffer bytes.Buffer
	logger := log.New(&logBuffer, "[test] ", log.LstdFlags)

	handler := loggingMiddleware(logger, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	logOutput := logBuffer.String()
	if !strings.Contains(logOutput, "INFO:") {
		t.Error("expected log output to contain 'INFO:'")
	}
	if !strings.Contains(logOutput, "GET /test 200") {
		t.Error("expected log output to contain request details")
	}
}

func TestPanicRecovery(t *testing.T) {
	var logBuffer bytes.Buffer
	logger := log.New(&logBuffer, "[test] ", log.LstdFlags)

	handler := loggingMiddleware(logger, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	}))

	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	logOutput := logBuffer.String()
	if !strings.Contains(logOutput, "ERROR: Panic recovered") {
		t.Error("expected panic to be recovered and logged")
	}
}

func BenchmarkPingHandler(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		pingHandler(rec, req)
	}
}

func BenchmarkInfoHandler(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/info?test=value", nil)
	req.Header.Set("User-Agent", "benchmark-test")
	req.Host = "benchmark.test"
	req.RemoteAddr = "127.0.0.1:8080"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		infoHandler(rec, req)
	}
}
