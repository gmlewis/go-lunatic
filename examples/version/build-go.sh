#!/bin/bash -ex
GOOS=wasip1 GOARCH=wasm go build -o main.wasm .
