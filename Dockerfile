# # Build stage
# FROM golang:1.17 AS build

# WORKDIR /app

# COPY . .

# # Install Go module dependencies
# RUN go mod download

# # Build the Go app
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/main.go

# # Run stage
# FROM alpine:3.14 AS run

# WORKDIR /app

# # Install MySQL client
# RUN apk add --no-cache mysql-client

# # Copy the Go app from the build stage
# COPY --from=build /app/app .

# # # Copy the script to initialize the database
# # COPY init_db.sh .

# # Copy the .env file to the build directory
# COPY ../.env ./

# # # Make the script executable
# # RUN chmod +x init_db.sh

# # Expose the port
# EXPOSE 9000

# # Start the app
# CMD ["./app"]

# Build stage
FROM golang:1.17 AS build

WORKDIR /app

COPY . .

# Install Go module dependencies
RUN go mod download

# Copy the .env file to the build directory
COPY .env .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/main.go

# Run stage
FROM alpine:3.14 AS run

WORKDIR /app

# Install MySQL client
RUN apk add --no-cache mysql-client

# Copy the Go app from the build stage
COPY --from=build /app/app .

# Copy the .env file to the run directory
COPY --from=build /app/.env .

# Expose the port
EXPOSE 9000

# Start the app
CMD ["./app"]
