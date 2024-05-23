# Golang base image
FROM golang:1.22 as go_builder
LABEL stage="nxd_builder"
WORKDIR /nxd
COPY . .
WORKDIR /nxd/cmd/nxd
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o nxd

# New stage from scratch
FROM alpine:3.20
LABEL stage="nxd_binary"
RUN apk add --no-cache tzdata
COPY --from=go_builder /nxd/cmd/nxd/nxd /cmd/nxd/nxd
WORKDIR /cmd/nxd
