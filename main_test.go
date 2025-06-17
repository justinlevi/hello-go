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
