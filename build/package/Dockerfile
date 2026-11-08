# Golang base image
FROM golang:1.27.3 AS go_builder
LABEL stage="nxd_builder"
WORKDIR /nxd
COPY . .
WORKDIR /nxd/cmd/nxd
RUN CGO_ENABLED=0 go build -mod=vendor -ldflags="-s -w" -o nxd

# New stage from scratch
FROM alpine:3.24.3
LABEL stage="nxd_binary"
RUN apk add --no-cache tzdata && adduser -D -u 10001 nxd
COPY --from=go_builder /nxd/cmd/nxd/nxd /usr/local/bin/nxd
WORKDIR /home/nxd
USER nxd
ENTRYPOINT ["/usr/local/bin/nxd"]
