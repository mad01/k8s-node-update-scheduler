FROM golang:1.8.4-jessie as builder
ENV buildpath=/go/src/github.com/mad01/termination-scheduler
RUN mkdir -p $buildpath
WORKDIR $buildpath

COPY . .

RUN make build/release
RUN make test

FROM debian:8
COPY --from=builder /go/src/github.com/mad01/termination-scheduler/_release/termination-scheduler /termination-scheduler

ENTRYPOINT ["/termination-scheduler"]
