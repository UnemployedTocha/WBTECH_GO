FROM golang:latest
LABEL authors="unemployedtocha"

COPY ./ ./

RUN go build -o demo_service ./cmd/app/main.go
CMD ["./demo_service"]