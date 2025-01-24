# ---- Build Stage ----
FROM golang:1.23.2-alpine AS builder

# 1. Install dependencies AS ROOT (before switching to non-root user)
RUN apk add --no-cache git

# 2. Create non-root user
RUN addgroup -S app && adduser -S -G app app
USER app

# 3. Rest of the build stage remains unchanged
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-s -w" -trimpath -o /app/main ./cmd/server

# ---- Final Stage ----
FROM alpine:3.21
RUN apk --no-cache add ca-certificates
RUN addgroup -S app && adduser -S -G app app
USER app
WORKDIR /app
COPY --from=builder --chown=app:app /app/main .
EXPOSE 9090
CMD ["./main"]