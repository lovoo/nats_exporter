FROM alpine:latest

ENV GOPATH /go
ENV APPPATH $GOPATH/src/github.com/lovoo/nats_exporter

COPY . $APPPATH

RUN apk add --update -t build-deps go git mercurial \
    && cd $APPPATH && go get -d && go build -o /nats_exporter \
    && apk del --purge build-deps git mercurial && rm -rf $GOPATH

EXPOSE 9118

ENTRYPOINT ["/nats_exporter"]
