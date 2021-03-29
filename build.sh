#!/bin/sh

rm -rf build/
mkdir build/

cp /usr/local/lib/libws2811.so /usr/local/go/pkg/linux_arm_dynlink/

go build -o=build/monolith/ ./screen/ ./ledstrip/

go install -buildmode=shared -linkshared ./common/pins/ ./common/utils/
go build -linkshared -o=build/shared/ ./screen/ ./ledstrip/

cp /usr/local/go/pkg/linux_arm_dynlink/*.so build/shared/
cp /go/pkg/linux_arm_dynlink/*.so build/shared/