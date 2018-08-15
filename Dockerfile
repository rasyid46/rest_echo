FROM golang:1.10-alpine

RUN apk add --update tzdata bash wget curl git;

RUN mkdir -p $$GOPATH/bin && \
    go get -u github.com/golang/dep/cmd/dep && \
    dep ensure -add github.com/labstack/echo@^3.1

# directory
ADD . /go/src/rest_echo
WORKDIR /go/src/rest_echo

CMD dep ensure && go build main && ./main

EXPOSE 8000