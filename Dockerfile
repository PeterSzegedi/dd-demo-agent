# Use the official Go image based on Alpine
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy application files to the container
COPY . .

# Install necessary dependencies
RUN apk add --no-cache git

# Build the Go application
RUN go build -o app .

# Expose the port the app runs on
EXPOSE 8080

# Define the command to run the application
CMD ["./app"]