# Stage 1: Build the Go application
FROM golang:1.22.6-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o url-shortener .

# Stage 2: Run the Go application in a lightweight image
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go binary from the builder stage
COPY --from=builder /app/url-shortener .

# Expose the port that the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./url-shortener"]
