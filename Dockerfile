# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /quadratic-solver

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /quadratic-solver /app/quadratic-solver
ENTRYPOINT ["/app/quadratic-solver"]