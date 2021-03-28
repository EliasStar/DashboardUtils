#!/bin/sh

GOOS=linux
GOARCH=arm
GOARM=6
# CGO_ENABLED=1

rm -rf build
mkdir build

cp /usr/local/lib/libws2811.so ./build/

# go install -buildmode=shared -linkshared std
# go build -buildmode=archive -o ./build/ github.com/EliasStar/DashboardUtils/pins github.com/EliasStar/DashboardUtils/utils
# go install -buildmode=shared -linkshared ./pins/ ./utils/
# go build -linkshared -o ./build/ ./screen/ ./ledstrip/