# ---- Build Stage ----
FROM golang:1.23.2-alpine AS builder

# 1. Install dependencies as root
RUN apk add --no-cache git

# 2. Create non-root user and workspace with proper permissions
RUN addgroup -S app && adduser -S -G app app \
    && mkdir -p /app \
    && chown -R app:app /app  # Fixes permission issues

# 3. Switch to non-root user
USER app
WORKDIR /app

# 4. Copy dependency files (with ownership)
COPY --chown=app:app go.mod go.sum ./

# 5. Download dependencies
RUN go mod download

# 6. Copy source code (with ownership)
COPY --chown=app:app . .

# 7. Build binary
RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-s -w" -trimpath -o ./main ./cmd/server

# ---- Final Stage ----
FROM alpine:3.21
RUN apk --no-cache add ca-certificates
RUN addgroup -S app && adduser -S -G app app
USER app
WORKDIR /app
COPY --from=builder --chown=app:app /app/main .
EXPOSE 9090
CMD ["./main"]