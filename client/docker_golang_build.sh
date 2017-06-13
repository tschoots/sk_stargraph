#!/bin/bash

_GO_VERSION=1.8
_TARGET_OS=linux
_TARGET_ARCH=amd64
_IMAGE='maiastra/sk_stargraph_client'
_TAG='0.1'


printf "build executable"
docker run --rm -v $(pwd):/usr/src/myapp -w /usr/src/myapp -e GOOS=$_TARGET_OS -e GOARCH=$_TARGET_ARCH golang:$_GO_VERSION bash -c ./build.sh
printf "\r"
printf "building image %s:%s\n" $_IMAGE $_TAG
docker build -t $_IMAGE:$_TAG .
echo "image build you can start with :"
echo "docker run -d -p 8888:8080 -e STARGRAPH_SERVER_IP=<ip of stargraph server> $_IMAGE:$_TAG"
