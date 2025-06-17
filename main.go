package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "[hello-api] ", log.LstdFlags|log.Lmicroseconds)

	mux := http.NewServeMux()
	mux.Handle("/hello", loggingMiddleware(logger, http.HandlerFunc(helloHandler)))
	mux.Handle("/health", loggingMiddleware(logger, http.HandlerFunc(healthHandler)))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		logger.Println("INFO: Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("ERROR: Server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Println("INFO: Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("ERROR: Server forced to shutdown: %v", err)
	}

	logger.Println("INFO: Server exited")
}

type Response struct {
	Message string `json:"message"`
}

type Request struct {
	Name string `json:"name"`
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

var requestCounter uint64

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var name string

	switch r.Method {
	case http.MethodPost:
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			respondWithError(w, http.StatusUnsupportedMediaType, "Content-Type must be application/json", "INVALID_CONTENT_TYPE")
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, 1048576) // 1MB limit

		var req Request
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&req); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid JSON", "INVALID_JSON")
			return
		}

		name = req.Name
	default:
		name = r.URL.Query().Get("name")
	}

	if name == "" {
		name = "World"
	}

	message := "Hello, " + name + "!"
	resp := Response{Message: message}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("ERROR: Failed to encode response: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error", "ENCODING_ERROR")
		return
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{Status: "healthy"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("ERROR: Failed to encode health response: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error", "ENCODING_ERROR")
		return
	}
}

func respondWithError(w http.ResponseWriter, code int, message string, errorCode string) {
	errResp := ErrorResponse{
		Error: message,
		Code:  errorCode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		log.Printf("ERROR: Failed to encode error response: %v", err)
		http.Error(w, message, code)
	}
}

func loggingMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		requestID := atomic.AddUint64(&requestCounter, 1)

		wrapped := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		defer func() {
			if err := recover(); err != nil {
				logger.Printf("ERROR: Panic recovered: %v", err)
				respondWithError(w, http.StatusInternalServerError, "Internal Server Error", "PANIC_RECOVERY")
			}
		}()

		next.ServeHTTP(wrapped, r)

		duration := time.Since(start)
		logger.Printf("INFO: [%d] %s %s %d %v", requestID, r.Method, r.URL.Path, wrapped.statusCode, duration)
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
