#!/bin/bash

GOOS=js GOARCH=wasm go build -o web/breakout.wasm
cp $(go env GOROOT)/misc/wasm/wasm_exec.js web/
