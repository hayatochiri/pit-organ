FROM golang:1.11.9-alpine3.9

RUN apk add --update git gcc libc-dev &&\
    go get -u github.com/golang/dep/cmd/dep &&\
    go get -u github.com/rakyll/gotest &&\
    go get -u golang.org/x/tools/cmd/godoc &&\
    rm -rf /var/cache/apk/* &&\
    echo

ENTRYPOINT godoc --http=0.0.0.0:8080
# CMD tail -f /dev/null
