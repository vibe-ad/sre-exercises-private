# Buggy Go program

`docker build -t buddy-go:local .`

`docker run --rm --volume /usr/share/zoneinfo:/usr/share/zoneinfo:ro -p 8080:8080 -it buddy-go:local `

`ab -n 30000 -c 30 'http://localhost:8080/hello'`
