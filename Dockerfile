# Use an official lightweight Go image
FROM golang:1.23.1-alpine

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the current directory into the container
COPY . .

# Set the working directory to the directory with your main.go file
WORKDIR /app/cmd/smart-plan-2

# Build the application binary
RUN go build -o /app/smart-plan-2

# Expose port 8080 for Cloud Run
EXPOSE 8080

# Command to run the executable
CMD ["/app/smart-plan-2"]
