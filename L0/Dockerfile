FROM golang:latest
LABEL authors="unemployedtocha"

COPY ./ ./

RUN go build -o demo_service ./cmd/app/main.go && \
    go build -o producer ./cmd/producer/main.go

CMD ["./demo_service"]