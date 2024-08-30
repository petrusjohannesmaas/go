# Use the official Go image as the builder
FROM golang:1.23 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Use a lightweight image for the final container
FROM alpine:latest

# Install necessary libraries to run the app
RUN apk --no-cache add ca-certificates

# Set environment variables if needed, e.g., database URL
ENV DB_URL=postgres://postgres:password@postgres:5432/pgtest?sslmode=disable

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the Go app listens on
EXPOSE 8080

# Run the Go app
CMD ["./main"]
