#!/usr/bin/env bash
#!/bin/sh


set -e
if ! [[ "$0" =~ scripts/_watch.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi


LAST=`find . | egrep ".sh|.go" | grep -v 'build/'| xargs ls -l -D '%s' | awk '{print $6}' | sort -n | tail -n 1 `
echo "LAST=$LAST"

LAST2=""

while [ true ]
do
  LAST2=`find . | egrep ".sh|.go" | grep -v 'build/'| xargs ls -l -D '%s' | awk '{print $6}' | sort -n | tail -n 1 `
  
  # echo "LAST2=$LAST2"

  if ! [ "${LAST}" -eq "${LAST2}" ]; then

  # then
    echo "LAST=$LAST"
    echo "LAST  2=$LAST2"
    rm -rf ./build/* || true
    LAST=$LAST2

    # gofmt
    test -z "$(gofmt -s -l -w  -d $(find . -type f -name '*.go' -not -path "*/vendor/*") | tee /dev/stderr)"
    # govet 
    test -z "$(go vet -all ./...  | tee /dev/stderr)"


    mkdir -p ./build

    echo "Building indexvm in ./build/indexvm"
    go build -o ./build/indexvm ./cmd/indexvm

    echo "Building index-cli in ./build/index-cli"
    go build -o ./build/index-cli ./cmd/index-cli

    find ./build
    ./build/index-cli version
  fi

done



