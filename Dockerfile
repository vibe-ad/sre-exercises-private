FROM alpine:3.15 as alpine

WORKDIR /app

# Install timezone data
COPY .env .env
RUN apk add -U --no-cache tzdata strace

EXPOSE 8080

COPY bugg-go .

STOPSIGNAL SIGQUIT

ENTRYPOINT [ "/app/bugg-go" ]
