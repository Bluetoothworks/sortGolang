# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go server source code into the container
COPY . .

# Initialize Go modules
RUN go mod init mygoapp

# Build the Go application inside the container
RUN go build -o go-server

# Expose the port that the Go server will run on
EXPOSE 8000

# Command to run the Go server when the container starts
CMD ["./go-server"]
