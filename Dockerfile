FROM golang:1.20.4 as builder

COPY go.mod /
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app

FROM migrate/migrate

ENTRYPOINT ["/app"]