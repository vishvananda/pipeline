#!/usr/bin/env bash
repository=${repository:-vishvananda/pipeline}

echo BUILDING ${repository}:builder

docker build -t ${repository}:builder \
    --build-arg http_proxy="${http_proxy}" \
    --build-arg https_proxy="${https_proxy}" \
    --build-arg ftp_proxy="${ftp_proxy}" \
    .
