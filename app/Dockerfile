FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o hello .

FROM alpine
COPY --from=builder /app/hello /hello
EXPOSE 8080
CMD ["/hello"]

