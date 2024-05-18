# Build stage
FROM golang:1.18.3-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go

# Run stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY db/migration ./db/migration

LABEL maintainer="Roman Bozhko <rm.bozhko@gmail.com>"

EXPOSE 8080
CMD ["/app/main"]