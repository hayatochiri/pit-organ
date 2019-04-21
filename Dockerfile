FROM golang:1.11.9-alpine3.9

RUN apk add --update git gcc libc-dev &&\
    go get -u github.com/golang/dep/cmd/dep &&\
    go get -u github.com/rakyll/gotest &&\
    rm -rf /var/cache/apk/* &&\
    echo

ENTRYPOINT tail -f /dev/null
# CMD tail -f /dev/null
