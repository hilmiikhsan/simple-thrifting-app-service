# Stage 1: Build Stage
FROM golang:1.22.8-alpine AS builder

# Mengatur direktori kerja
WORKDIR /app

# Menyalin file go.mod dan go.sum lalu mengunduh dependency
COPY go.mod go.sum ./
RUN go mod download

# Menyalin seluruh source code dan file .env
COPY . .
COPY .env .

# Build binary dengan CGO dinonaktifkan dan flag optimasi untuk mengurangi ukuran binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o thrifting-app-service

# Stage 2: Final Stage dengan image Alpine minimal
FROM alpine:3.18

# Install CA certificates (jika diperlukan)
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Menyalin binary dan file .env dari builder stage
COPY --from=builder /app/thrifting-app-service .
COPY --from=builder /app/.env .

# Mengekspos port yang diperlukan
EXPOSE 9002

# Menjalankan aplikasi
CMD ["./thrifting-app-service"]
