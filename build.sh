#!/bin/sh

rm -rf build/
mkdir build/

go build -o=build/ ./screen/ ./ledstrip/