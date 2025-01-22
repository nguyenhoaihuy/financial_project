# Stage 1: Build the Go binary
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files into the container
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Copy the .env file into the build stage


# Build the binary
RUN go build -o ./bin/company_info_collector ./cmd/company_info_collector/company_info_collector.go
RUN go build -o ./bin/financial_statement_collector ./cmd/financial_statement_collector/financial_statement_collector.go

# Stage 2: Use Alpine with upgraded glibc
FROM alpine:latest

# Install glibc
RUN apk add --no-cache libc6-compat
# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/bin/financial_statement_collector .
COPY --from=builder /app/bin/company_info_collector .
COPY --from=builder /app/bin/.env .
