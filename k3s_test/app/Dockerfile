# Use a Golang base image
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy Go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o go-app

# Use a smaller base image for the final stage
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the Go application from the builder stage
COPY --from=builder /app/go-app .

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./go-app"]
