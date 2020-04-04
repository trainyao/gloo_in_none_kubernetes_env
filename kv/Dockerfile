FROM golang:1.13.0-stretch as builder

WORKDIR /build

# Resolve and build Go dependencies as Docker cache
COPY go.mod /build/go.mod
COPY go.sum /build/go.sum
COPY kv/go.mod /build/kv/go.mod

ENV GO111MODULE=on

COPY service.go /build/main.go
COPY client /build/client
COPY kv/ /build/kv
COPY vendor/ /build/vendor

# Build for linux
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0 
RUN go build -mod=vendor -o server
RUN cd client && go build -mod=vendor -o client

# Build the main container (Linux Runtime)
FROM alpine:latest
WORKDIR /root/

# Copy the linux amd64 binary
COPY --from=builder /build/server /bin/
COPY --from=builder /build/client /bin/

ENTRYPOINT /bin/server
