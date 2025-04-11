# --------- Builder stage ---------
    FROM golang:1.23-alpine AS builder

    WORKDIR /app
    
    # Install dependencies
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy source code
    COPY . .
    
    # Build the Go server
    RUN go build -o /url-shortener ./cmd/server
    
    # --------- Final minimal image ---------
    FROM alpine:latest
    
    WORKDIR /app
    
    COPY --from=builder /url-shortener .
    
    # Optional: for pretty Zerolog output (if using color)
    ENV TERM=xterm-256color
    
    EXPOSE 8080 9090
    ENTRYPOINT ["./url-shortener"]
    