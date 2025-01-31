
FROM golang:1.22-alpine

# Set up environment and install necessary packages
RUN apk add --no-cache git netcat-openbsd gcc musl-dev

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

RUN ls 

# Build the Go application
WORKDIR /app/cmd
RUN go build -o /app/main .

CMD ["/app/main"]
