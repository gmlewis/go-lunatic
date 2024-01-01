#!/bin/bash -ex
tinygo build -o main.wasm -target=wasi main.go
