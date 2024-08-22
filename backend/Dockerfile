# Use the official Golang image as a base image
FROM golang:1.19

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/app/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
