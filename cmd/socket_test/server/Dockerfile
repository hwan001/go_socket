FROM golang:1.22-alpine

COPY server /app
WORKDIR /app

RUN go build -o server server.go

CMD ["./server"]
