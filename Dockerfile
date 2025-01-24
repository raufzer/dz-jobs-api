# ---- Build Stage ----
FROM golang:1.23.2-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Create non-root user
RUN addgroup -S app && adduser -S -G app app
USER app

# Set working directory
WORKDIR /app

# Copy dependency files first (better caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all source files
COPY . .

# Build binary from cmd/server/main.go
RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-s -w" -trimpath -o /app/main ./cmd/server

# ---- Final Stage ----
FROM alpine:3.21

# Install CA certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -S app && adduser -S -G app app
USER app

WORKDIR /app

# Copy compiled binary from builder
COPY --from=builder --chown=app:app /app/main .

EXPOSE 9090
CMD ["./main"]