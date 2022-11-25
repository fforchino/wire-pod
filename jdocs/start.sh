#!/bin/bash

if [[ $EUID -ne 0 ]]; then
  echo "This script must be run as root. sudo ./start.sh"
  exit 1
fi

if [[ -d ./jdocs ]]; then
   cd jdocs 
fi

#if [[ ! -f ./source.sh ]]; then
#  echo "You need to make a source.sh file. This can be done with the setup.sh script."
#  exit 0
#fi

#source source.sh

if [[ -f ./jdocs ]]; then
  ./jdocs
else
  /usr/local/go/bin/go run cmd/main.go
fi

