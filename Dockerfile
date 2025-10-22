# ---------- BUILD STAGE ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

# copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# copy semua source code
COPY . .

# build binary
RUN go build -o main .

# ---------- RUN STAGE ----------
FROM alpine:latest

WORKDIR /app

# copy hasil build dari builder
COPY --from=builder /app/main .

# expose port
EXPOSE 8001

# jalankan app
CMD ["./main"]
