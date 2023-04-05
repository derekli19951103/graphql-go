# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the server files into the container at /app
COPY . /app

# Build the server executable
RUN go build -o server server.go

# Expose port 8080 for the server
EXPOSE 8080

# Run the server when the container starts
CMD ["./server"]