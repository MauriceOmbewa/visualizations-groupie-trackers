# Use the official Golang image from Bitnami
FROM bitnami/golang:1.22.2
# Metadata
LABEL maintainer="oathooh@gmail.com"
LABEL version="1.0"
LABEL description="A Go web server for Groupie-trackers"
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go.mod and go.sum files
COPY go.mod ./
# Copy the source code into the container
COPY . .
# Build the Go app
RUN go build -o main .
# Expose port 8080 to the outside world
EXPOSE 8080
# Command to run the executable
CMD ["./main"]