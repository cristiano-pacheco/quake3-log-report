# Stage 1: Build the Go application
FROM golang:1.22.5-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Install git to fetch dependencies
RUN apk update && apk add --no-cache git

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app and name the executable quakelogreport
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o quakelogreport ./cmd/cli

# Stage 2: Run the Go application
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/quakelogreport .

# Copy the log file to the container (optional if needed for the test run)
COPY quake3.log .

# Command to run the executable with default arguments
CMD ["./quakelogreport"]
