FROM golang:1.24-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git

COPY person/go.mod person/go.sum ./
RUN go mod download

COPY person/configs /app/configs
COPY person/. .

RUN go test ./...
RUN go build -o app .

FROM alpine:latest

WORKDIR /app
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/app /app/app
COPY --from=builder /app/configs /app/configs

EXPOSE 8080
CMD ["/app/app"]
