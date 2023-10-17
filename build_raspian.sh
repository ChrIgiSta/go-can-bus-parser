#!/bin/bash

# enable c compiler
#export CGO_ENABLED=1
# set compiler to mingw
#export CC=arm-linux-gnueabi-gcc 

export GOOS=linux
export GOARCH=arm
export GOARM=7

go mod tidy

go build -o opel_bc.bin .
