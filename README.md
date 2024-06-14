# Buggy Go program

## Build
```
./build.sh
```
Commit all output from dist into sre-exercise public repository


# Usage:


```
docker build -t bugg-go:local .
docker run --rm -p 8080:8080 --name bugg-go -it bugg-go:local
curl 'http://localhost:8080/hello'
```

Find why this does not work


## Resolution

Manual exploration:
docker exec -it bugg-go /bin/sh
Check what's in current folder, notices the .env
Check what's in /etc/hosts

Like a boss: Edit Dockerfile
ENTRYPOINT [ "strace", "-e", "trace=open,openat,close,read,write,connect,accept", "/app/bugg-go" ]
```
openat(AT_FDCWD, ".env", O_RDONLY|O_CLOEXEC) = 3
read(3, "LANGUAGE=en_US:en\nSHELL=/bin/bas"..., 4096) = 34
read(3, "", 4062)                       = 0
close(3)                                = 0

openat(AT_FDCWD, "/etc/hosts", O_WRONLY|O_CREAT|O_APPEND|O_CLOEXEC, 0600) = 3
write(3, "127.0.1.2 httpbin.org", 21)   = 21
```
