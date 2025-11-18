FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod first
COPY go.mod ./

# Copy source code before tidy (important for dependency detection)
COPY . .

# Download dependencies (creates go.sum)
RUN go mod tidy

# Build binary
RUN go build -o hello .

# Final lightweight runtime image
FROM alpine

COPY --from=builder /app/hello /hello

EXPOSE 8080

CMD ["/hello"]

