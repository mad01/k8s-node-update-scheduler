FROM golang:1.8.4-jessie as builder
ENV buildpath=/go/src/github.com/mad01/k8s-node-update-scheduler
RUN mkdir -p $buildpath
WORKDIR $buildpath

COPY . .

RUN make build/release
RUN make test

FROM debian:8
COPY --from=builder /go/src/github.com/mad01/k8s-node-update-scheduler/_release/k8s-node-update-scheduler /k8s-node-update-scheduler

ENTRYPOINT ["/k8s-node-update-scheduler"]
