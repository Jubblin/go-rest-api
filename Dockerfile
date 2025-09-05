# Build stage
# checkov:skip=CKV_DOCKER_2: Healthcheck not required for this application
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
FROM alpine:3.22.1

# Install ObjectBox
RUN curl -Ls https://raw.githubusercontent.com/objectbox/objectbox-go/main/install.sh | sh

# Create a non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy any additional files needed (like swagger docs)
COPY --from=builder /app/docs ./docs

# Change ownership of the app directory to the non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8080

# Set environment variables
ENV GIN_MODE=release

# Run the application
CMD ["./main"] 