# Use the official Golang image as a build environment
FROM golang:1.21-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Install build tools
RUN apk add --no-cache git curl

# Install reflex
RUN go install github.com/cespare/reflex@latest

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run reflex
CMD ["reflex", "-r", "\\.go$", "-s", "--", "sh", "-c", "go mod tidy && go build -o main . && ./main"]
