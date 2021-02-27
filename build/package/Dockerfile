# Golang base image
FROM golang:1.15 as go_builder
LABEL stage="nxd_builder"
WORKDIR /go/src/github.com/rl404/nyaa-x-discord
COPY . .
WORKDIR /go/src/github.com/rl404/nyaa-x-discord/cmd/nxd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -mod vendor -o nxd

# New stage from scratch
FROM alpine:3.13
RUN apk add --no-cache tzdata
COPY --from=go_builder /go/src/github.com/rl404/nyaa-x-discord/cmd/nxd/nxd nxd