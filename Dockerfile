FROM golang:1.15-alpine

WORKDIR /app

RUN adduser -D -u 1000 myuser

COPY main.go .

RUN go build -o main

USER myuser

EXPOSE 8080

CMD ["./main"]
