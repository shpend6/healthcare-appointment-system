# 1. Build stage with CGO support
FROM golang:1.24-bookworm AS builder

# 2. Set working directory
WORKDIR /app

# 3. Install system dependencies for CGO + SQLite
RUN apt-get update && apt-get install -y gcc libsqlite3-dev

# 4. Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# 5. Copy source code
COPY . .

# 6. Enable CGO explicitly
ENV CGO_ENABLED=1

# 7. Build binary
RUN go build -o server ./cmd/api/

# 8. Final runtime image
FROM debian:bookworm-slim

WORKDIR /app

# 9. Copy binary
COPY --from=builder /app/server .

# 10. Expose port
EXPOSE 8080

# 11. Start the app
CMD ["./server"]
