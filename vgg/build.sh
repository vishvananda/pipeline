#!/usr/bin/env bash

repository=${repository:-vishvananda/pipeline}

echo BUILDING ${repository}:vgg

docker build -t ${repository}:vgg \
    --build-arg http_proxy="${http_proxy}" \
    --build-arg https_proxy="${https_proxy}" \
    --build-arg ftp_proxy="${ftp_proxy}" \
    .
