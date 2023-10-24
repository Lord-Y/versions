#!/bin/bash

usage(){
  echo "$1 argument is mantatory"
  echo "Example: $0 1.19 1.20"
}

if [ -n "$1" ]
then
  export OLD_VERSION="$1"
else
  usage OLD_VERSION
  exit 1
fi
shift

if [ -n "$1" ]
then
  export NEW_VERSION="$1"
else
  usage NEW_VERSION
  exit 1
fi
shift

OS=$(uname -s | tr -s "A-Z" "a-z")
if [[ "${OS}" == "darwin" ]]
then
  find . -type f -name go.mod -exec sed -i '' "s#^go ${OLD_VERSION}#go ${NEW_VERSION}#g" {} \;
  find . -type f -name "Dockerfile*" -exec sed -i '' "s#golang:${OLD_VERSION}#golang:${NEW_VERSION}#g" {} \;
  find . -type f -name "docker-compose*" -exec sed -i '' "s#golang:${OLD_VERSION}#golang:${NEW_VERSION}#g" {} \;
else
  find . -type f -name go.mod -exec sed -i "s#^go ${OLD_VERSION}#go ${NEW_VERSION}#g" {} \;
  find . -type f -name "Dockerfile*" -exec sed -i "s#golang:${OLD_VERSION}#golang:${NEW_VERSION}#g" {} \;
  find . -type f -name "docker-compose*" -exec sed -i "s#golang:${OLD_VERSION}#golang:${NEW_VERSION}#g" {} \;
fi

go mod tidy