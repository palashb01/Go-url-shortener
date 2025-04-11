# --------- Builder stage ---------
    FROM golang:1.23-alpine AS builder

    WORKDIR /app
    
    COPY go.mod go.sum ./
    RUN go mod download
    
    COPY . .
    RUN go build -o /url-shortener ./cmd/server
    
    # --------- Final minimal image ---------
    FROM alpine:latest
    
    WORKDIR /app
    
    # ✅ Copy the binary
    COPY --from=builder /url-shortener .
    
    # ✅ Copy the frontend folder explicitly
    COPY --from=builder /app/frontend ./frontend
    
    # Optional: for pretty logs
    ENV TERM=xterm-256color
    
    EXPOSE 8080 9090
    ENTRYPOINT ["./url-shortener"]
    