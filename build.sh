#!/bin/bash
mkdir -p dist
CGO_ENABLED=0 go build -ldflags "-s -w" -o primes_no_symbols_from_go_build -o ./bugg-go ./main.go
cp bugg-go dist
cp .env dist
cp Dockerfile dist
