# Build stage
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Install necessary build tools
RUN apk add --no-cache make gcc musl-dev

# Copy go mod and sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN make build

# Final stage
FROM alpine:latest

# Install necessary runtime packages
RUN apk --no-cache add ca-certificates tzdata

# Set working directory
WORKDIR /app

# Create uploads directory
RUN mkdir -p /app/uploads

# Copy binary from builder
COPY --from=builder /app/bookpavilion .

# Copy any additional required files
COPY --from=builder /app/config/config.go ./config/

# Expose port
EXPOSE 8080

# Create volume for uploads
VOLUME ["/app/uploads"]

# Set environment variables
ENV GIN_MODE=release \
    PORT=8080 \
    UPLOAD_DIR=/app/uploads

# Run the application
CMD ["./bookpavilion"]
