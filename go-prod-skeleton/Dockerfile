# syntax=docker/dockerfile:1

FROM golang:1.23-alpine AS builder

WORKDIR /src

COPY go.* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/app ./cmd/api

FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata \
    && adduser -D -u 10001 app

USER app

COPY --from=builder /out/app /app

EXPOSE 8080

ENTRYPOINT ["/app"]
