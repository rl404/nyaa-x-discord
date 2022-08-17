# Golang base image
FROM golang:1.17 as go_builder
LABEL stage="nxd_builder"
WORKDIR /nxd
COPY . .
WORKDIR /nxd/cmd/nxd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -mod vendor -o nxd

# New stage from scratch
FROM alpine:3.16
LABEL stage="nxd_binary"
RUN apk add --no-cache tzdata
COPY --from=go_builder /nxd/cmd/nxd/nxd /cmd/nxd/nxd
WORKDIR /cmd/nxd
