FROM golang:1.24.0

WORKDIR /app

# Copy Go files
COPY . .

# Build the Go app
RUN go build -o backend_go .

# Expose port 8080 for the backend
EXPOSE 8080

# Run the Go backend
CMD ["./backend_go"]
