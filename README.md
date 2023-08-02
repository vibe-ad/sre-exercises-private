# Buggy Go program

`docker build -t buddy-go:local .`

`docker run -p 8080:8080 -it buddy-go:local `

`ab -n 200 -c 30 'http://localhost:8080/test'`
