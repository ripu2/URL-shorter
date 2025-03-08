# Use the correct Go version
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all project files
COPY . .  

# Build the binary
RUN go build -o main ./main.go  # <- Check kar ki yahi entry file hai!

# Expose the application port
EXPOSE 8080

# Set environment variables
ENV REDIS_HOST=my_redis
ENV REDIS_PORT=6379

# Start the app using Air
CMD ["air"]
