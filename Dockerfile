# Builder phase.
FROM golang:1.10 AS builder
MAINTAINER dosmith@redhat.com
ENV build_date 2018-04-12
ENV GOPATH /usr/
RUN mkdir -p /usr/src/
WORKDIR /usr/src/
RUN git clone https://github.com/dougbtv/k8s-dummy-device-plugin.git
WORKDIR /usr/src/k8s-dummy-device-plugin
# RUN go build dummy.go
RUN CGO_ENABLED=0 go build -a dummy.go

# Copy phase
FROM alpine:latest
# If you need to debug, add bash.
# RUN apk add --no-cache bash
COPY --from=builder /usr/src/k8s-dummy-device-plugin/dummy /dummy
COPY --from=builder /usr/src/k8s-dummy-device-plugin/dummyResources.json /dummyResources.json
ENTRYPOINT ["/dummy"]
