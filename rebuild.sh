#!/bin/bash

GOOS=js GOARCH=wasm go build -o generated/breakout.wasm src/*.go
cp $(go env GOROOT)/misc/wasm/wasm_exec.js generated/
