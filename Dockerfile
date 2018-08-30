ARG repository
FROM $repository:builder as builder

RUN mkdir -p /go/src/github.com/vishvananda/pipeline/cmd
ADD . /go/src/github.com/vishvananda/pipeline/cmd
RUN cd /go/src/github.com/vishvananda/pipeline/cmd && CGO_ENABLED=0 go build

FROM scratch
COPY --from=builder /go/src/github.com/vishvananda/pipeline/cmd/cmd /cmd
ENTRYPOINT ["/cmd"]
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
