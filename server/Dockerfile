FROM golang:1.22.1-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o go-http-server ./cmd/go-http-server/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/go-http-server .

EXPOSE 8080

CMD ["./go-http-server"]
