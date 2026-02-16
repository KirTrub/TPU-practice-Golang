
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server_app .


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server_app .

COPY --from=builder /app/db ./db

EXPOSE 8080

ENTRYPOINT [ "./server_app" ]