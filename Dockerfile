# syntax=docker/dockerfile:1

# Minimal Dockerfile for a plain Go project (no logic)
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /bin/tabbybot ./cmd/tabbybot

FROM alpine:3.18
COPY --from=builder /bin/tabbybot /usr/local/bin/tabbybot
ENTRYPOINT ["tabbybot"]