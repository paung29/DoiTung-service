# -------- Stage 1: Build --------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install git (needed for dependencies)
RUN apk add --no-cache git

# Copy go mod first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy all files
COPY . .

# Build binary
RUN go build -o main ./cmd

# -------- Stage 2: Run --------
FROM alpine:latest

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

# Expose port (Fiber default)
EXPOSE 8080

# Run app
CMD ["./main"]