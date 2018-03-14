FROM golang:alpine AS builder
RUN apk add --no-cache --update git
ENV GOPATH=/go

RUN go get -u github.com/golang/dep/cmd/dep
COPY . $GOPATH/src/github.com/jonaklab/snk
WORKDIR $GOPATH/src/github.com/jonaklab/snk

ARG NETRC
RUN echo "${NETRC}" > /root/.netrc \
  && ./build.sh

FROM alpine:latest
RUN apk add --no-cache --update ca-certificates openssl
COPY --from=builder /go/bin/snk /usr/local/bin/snk

ENTRYPOINT ["snk"]
