# Build binary.
FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-calculator

# Pack binary.
FROM ubuntu:22.04 AS build-release-stage

WORKDIR /home

COPY --from=builder /go-calculator ./go-calculator

EXPOSE 8080

CMD ["./go-calculator"]