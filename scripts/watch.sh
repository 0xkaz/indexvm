#!/usr/bin/env bash
#!/bin/sh


set -e
if ! [[ "$0" =~ scripts/watch.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi


LAST=`find . | egrep ".sh|.go" | grep -v 'build/'| xargs ls -l -D '%s' | awk '{print $6}' | sort -n | tail -n 1 `

LAST2=""
echo "LAST=$LAST"
echo "LAST2=$LAST2"

while [ true ]
do
  

  # echo "LAST2=$LAST2"

  if ! [ "${LAST}" -eq "${LAST2}" ]; then

  # then
    echo "LAST=$LAST"
    echo "LAST2=$LAST2"
    rm -rf ./build/* || true
    LAST2=$LAST

    # gofmt
    test -z "$(gofmt -s -l -w  -d $(find . -type f -name '*.go' -not -path "*/vendor/*") | tee /dev/stderr)"
    # govet 
    test -z "$(go vet -all ./...  | tee /dev/stderr)"


    mkdir -p ./build

    echo "Building indexvm in ./build/indexvm"
    go build -o ./build/indexvm ./cmd/indexvm

    # echo "Building weavedbvm in ./build/weavedbvm"
    # go build -o ./build/weavedbvm ./cmd/weavedbvm

    # echo "Building weavedb-cli in ./build/weavedb-cli"
    # go build -o ./build/weavedb-cli ./cmd/weavedb-cli

    echo "Building index-cli in ./build/index-cli"
    go build -o ./build/index-cli ./cmd/index-cli

    find ./build
    # ./build/token-cli version
    # ./build/weavedb-cli version
  fi

  LAST=`find . | egrep ".sh|.go" | grep -v 'build/'| xargs ls -l -D '%s' | awk '{print $6}' | sort -n | tail -n 1 `

done


