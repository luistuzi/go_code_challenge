FROM golang:1.27 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/api .

RUN chmod +x /app/api

EXPOSE 8080

CMD ["/app/api"]