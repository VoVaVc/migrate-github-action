FROM golang:1.20.4 as builder

COPY entrypoint.go /entrypoint.go
COPY go.mod go.sum /
RUN CGO_ENABLED=0 GOOS=linux go build -o /app

FROM migrate/migrate

ENTRYPOINT ["/app"]