# Pipeline

This repository contains example code for the following blog posts:

* [Machine Learning Model Pipelines: Part I](https://medium.com/@vishvananda/machine-learning-model-pipelines-part-i-e138b7a7c1ef)
* [Machine Learning Model Pipelines: Part II](https://medium.com/@vishvananda/machine-learning-model-pipelines-part-ii-23ebd1e6b714)

# Building

Docker images for the examples can be built using the `build.sh` script:

    ./build.sh <dir>

If you want to build the examples locally, you will have to install
graphpipe-go and its dependencies:

    go get -u github.com/oracle/graphpipe-go
    cd $GOPATH/src/github.com/oracle/graphpipe-go && make install-govendor deps

Then you should be able to build each example using `go build`:

    cd <dir> && go build

The only exceptions to the above are the two model servers and the builder,
which have a different build process:

    cd vgg && ./build.sh
    cd squeeze && ./build.sh
    cd builder && ./build.sh

The only time you would want to rebuild the builder is to pick up a new version
of `graphpipe-go`.

