FROM docker.io/library/golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server .

FROM docker.io/library/alpine:3.22.0
WORKDIR /app
COPY --from=builder /app/server /app/
EXPOSE 8080
CMD ["/app/server"]