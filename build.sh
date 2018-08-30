#!/usr/bin/env bash

if [ "$1" == "" ]; then
    echo "SPECIFY A BINARY TO BUILD" && exit 1
fi
repository=${repository:-vishvananda/pipeline}

echo BUILDING ${repository}:${1}

docker build -t ${repository}:${1} \
    --build-arg repository=${repository} \
    --build-arg http_proxy="${http_proxy}" \
    --build-arg https_proxy="${https_proxy}" \
    --build-arg ftp_proxy="${ftp_proxy}" \
    --file Dockerfile \
    ${1}
