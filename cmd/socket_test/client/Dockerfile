FROM golang:1.22-alpine

COPY client /app
WORKDIR /app

RUN go build -o client client.go

CMD ["./client"]
