# Use an official Golang runtime as the base image
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Install curl package
RUN apk --no-cache add curl

# Set the entry point for the container
CMD ["./main"]

# Expose the port your application is running on
EXPOSE 8080
