###################
# BUILD BASE IMAGE
###################
FROM golang:1.22.3-bullseye AS builder_base

WORKDIR /app

# Copy go.mod for dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy your project code
COPY . .

###################
# BUILD FOR DEVELOPMENT
###################
FROM builder_base AS builder_dev

# Install air for live reloading
RUN go install github.com/cosmtrek/air@v1.51.0

# Start the application using Air
CMD ["air"]

###################
# BUILD FOR PRODUCTION
###################
FROM builder_base AS builder_prod

# Build the project
RUN go build -o /app/main ./cmd/main.go

###################
# PRODUCTION
###################
FROM golang:1.22.3-bullseye

WORKDIR /app

# Copy the built binary from builder_prod stage
COPY --from=builder_prod /app/main ./
COPY --from=builder_prod /app/go.mod ./

# Start the application
CMD ["./main"]
