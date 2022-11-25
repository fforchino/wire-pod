#!/bin/bash
UNAME=$(uname -a)
echo "Building jdocs..."
if [[ "${UNAME}" == *"aarch64"* ]]; then
   ARCH=arm64
elif [[ "${UNAME}" == *"armv7l"* ]]; then
   ARCH=arm
elif [[ "${UNAME}" == *"x86_64"* ]]; then
   ARCH=amd64
fi

CGO_ENABLED=1 GOOS=linux GOARCH=${ARCH} /usr/local/go/bin/go build \
-ldflags "-w -s -extldflags "-static"" \
-trimpath \
-o jdocs cmd/main.go
echo "Built jdocs!"
