# Use the official Golang image as the base image
FROM golang:1.24

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY main.go .

# Build the Go application
RUN go build -o app main.go

# Set the entry point for the container
CMD ["/app/app"]

# Expose the application port
EXPOSE 8080