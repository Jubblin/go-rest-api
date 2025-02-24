# Build stage
FROM golang:1.24-alpine AS builder

# Install ObjectBox dependencies
RUN apk add --no-cache curl build-base git bash

# Install ObjectBox
RUN curl -Ls https://raw.githubusercontent.com/objectbox/objectbox-go/main/install.sh | sh

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Generate ObjectBox files
RUN go generate ./...

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Final stage
FROM alpine:3.21.3

# Install ObjectBox
RUN curl -Ls https://raw.githubusercontent.com/objectbox/objectbox-go/main/install.sh | sh

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy any additional files needed (like swagger docs)
COPY --from=builder /app/docs ./docs

# Expose port
EXPOSE 8080

# Set environment variables
ENV GIN_MODE=release

# Run the application
CMD ["./main"] 