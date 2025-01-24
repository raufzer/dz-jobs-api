# ---- Build Stage ----
FROM golang:1.23.2-alpine AS builder

# 1. Install build dependencies as root
RUN apk add --no-cache git \
    && go install github.com/swaggo/swag/cmd/swag@latest

# 2. Create dedicated user with permissions
RUN addgroup -S app && adduser -S -G app app \
    && mkdir -p /app \
    && chown -R app:app /app \
    && chmod -R 755 /go

# 3. Switch to app user
USER app
WORKDIR /app

# 4. Copy dependency files first (better caching)
COPY --chown=app:app go.mod go.sum ./

# 5. Download dependencies
RUN go mod download

# 6. Copy source code (including docs)
COPY --chown=app:app . .

# 7. Generate Swagger docs
RUN swag init -g cmd/server/main.go -o docs

# 8. Build binary
RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-s -w" -trimpath -o ./main ./cmd/server

# ---- Final Stage ----
FROM alpine:3.21

# 1. Install runtime dependencies
RUN apk --no-cache add ca-certificates

# 2. Create non-root user/group
RUN addgroup -S app && adduser -S -G app app

# 3. Set up final workspace
USER app
WORKDIR /app

# 4. Copy binary from builder
COPY --from=builder --chown=app:app /app/main .

# 5. Expose and run
EXPOSE 9090
CMD ["./main"]