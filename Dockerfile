# ---- Build Stage ----
FROM golang:1.23.2-alpine AS builder

# 1. Install build dependencies as root
RUN apk add --no-cache git \
    && go install github.com/swaggo/swag/cmd/swag@latest

# 2. Create app user and set up directories
RUN addgroup -S app && adduser -S -G app app \
    && mkdir -p /app \
    && chown -R app:app /app \
    && mkdir -p /go-mod-cache \
    && chown -R app:app /go-mod-cache

# 3. Set Go module environment variables
ENV GOMODCACHE=/go-mod-cache \
    GOPATH=/go

# 4. Switch to app user
USER app
WORKDIR /app

# 5. Copy dependency files
COPY --chown=app:app go.mod go.sum ./

# 6. Download dependencies
RUN go mod download

# 7. Copy source code
COPY --chown=app:app . .

# 8. Generate Swagger docs
RUN swag init -g cmd/server/main.go -o docs

# 9. Build binary
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