# Build stage
FROM golang:1.22 AS builder
WORKDIR /app

COPY . .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build \
      -ldflags="-w -s" \
      -o ./inu-backyard ./cmd/http_server/main.go

# Runner stage
FROM scratch AS runner
WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/inu-backyard /

EXPOSE 3001
CMD ["/inu-backyard"]
