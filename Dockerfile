FROM golang:1.15-alpine

WORKDIR /app

RUN adduser -D -u 1000 myuser

COPY main.go .

USER myuser

EXPOSE 8080

ENTRYPOINT [ "/usr/local/go/bin/go", "run", "main.go" ]
