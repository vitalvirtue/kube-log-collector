# Build stage
FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o kube-log-collector

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /cmd/kube-log-collector .

ENTRYPOINT ["./kube-log-collector"]
