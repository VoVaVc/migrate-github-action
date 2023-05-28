FROM golang:1.20.4 as builder

COPY entrypoint.go /entrypoint.go
RUN go build -o app

FROM migrate/migrate

ENTRYPOINT ["./app"]