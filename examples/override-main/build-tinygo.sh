#!/bin/bash -ex
tinygo build -o main.wasm -target=wasi main.go
go run ../../cmd/override-main/main.go main.wasm "$@"
