#!/usr/bin/env bash

repository=${repository:-vishvananda/pipeline}

wget -nc https://oracle.github.io/graphpipe/models/squeezenet.pb

echo BUILDING ${repository}:squeeze

docker build -t ${repository}:squeeze \
    --build-arg http_proxy="${http_proxy}" \
    --build-arg https_proxy="${https_proxy}" \
    --build-arg ftp_proxy="${ftp_proxy}" \
    .
