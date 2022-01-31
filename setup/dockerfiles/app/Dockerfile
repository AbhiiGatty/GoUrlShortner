FROM golang:1.17-alpine AS build_base

RUN apk add --no-cache git gcc

# Set the Current Working Directory inside the container
WORKDIR /tmp/GoUrlShortner/

# We want to populate the module cache based on the go.{mod,sum} files.
COPY ./ ./

RUN go mod download
RUN go mod tidy

# Unit tests
# RUN CGO_ENABLED=0 go test -v

# Build the Go app
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o ./binary_executable/GoUrlShortner ./cmd/main.go

# Start fresh from a smaller image
FROM alpine:3.9
RUN apk add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /app/

COPY --from=build_base /tmp/GoUrlShortner/binary_executable/GoUrlShortner /app/GoUrlShortner

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary_executable program produced by `go build`
ENTRYPOINT ["/app/GoUrlShortner"]