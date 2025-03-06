# Use the correct Go version
FROM golang:1.23

WORKDIR /app

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install a stable Air version (avoid latest)
RUN go install github.com/cosmtrek/air@v1.42.0

# Copy the rest of the project files
COPY . .

# Expose the application port
EXPOSE 8080

# Start the app using Air
CMD ["air"]
