# Stage 1: Build
FROM golang:1.23.2 as builder

# Set environment variables for Go
ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

# Set working directory
WORKDIR /build

# Cache dependencies by copying go.mod and go.sum separately
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN go build -o main ./cmd/server

# Stage 2: Final image
FROM alpine

# # Install certificates for HTTPS
# RUN apk --no-cache add ca-certificates && update-ca-certificates

# Set working directory
WORKDIR /root/

COPY ./config ./config

# Copy the binary from the builder stage
COPY --from=builder /build/main .

# Expose the application port (adjust based on your app)
EXPOSE 8002

# Command to run the application
CMD ["./main"]
