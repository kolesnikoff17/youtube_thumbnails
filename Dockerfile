FROM golang:1.18 AS builder
WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./server/cmd/main.go

FROM alpine AS server

WORKDIR /yt_thumbnails

COPY --from=builder /build/app .
COPY ./server/config/config.json .

CMD ["./app"]