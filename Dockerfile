# Build stage
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o hello-api \
    main.go

# Final stage - using alpine for health check support
FROM alpine:3.18

# Install ca-certificates for HTTPS support
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN adduser -D -u 1000 appuser

# Copy the binary from builder
COPY --from=builder /app/hello-api /hello-api

# Change ownership
RUN chown appuser:appuser /hello-api

# Expose port
EXPOSE 8080

# Switch to non-root user
USER appuser

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the binary
ENTRYPOINT ["/hello-api"]