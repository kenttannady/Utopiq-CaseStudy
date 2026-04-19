# Stage 1: Build
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o backend main.go

# Stage 2: Runtime (distroless - paling minimal)
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /app/backend /backend
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/backend"]