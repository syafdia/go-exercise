# Build binary.
FROM golang:1.17 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-calculator

# Pack binary.
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=builder /go-calculator /go-calculator

EXPOSE 8080

USER nonroot:nonroot

EXPOSE 8080

CMD ["/go-calculator"]