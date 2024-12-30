FROM golang:1.23.2-alpine
RUN addgroup app && adduser -S -G app app
USER app
WORKDIR /app
COPY go.mod go.sum ./
USER root
RUN chown -R app:app .
USER app
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 9090
CMD ["./main"]