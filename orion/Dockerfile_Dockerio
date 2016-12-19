FROM golang:latest

RUN go get -u github.com/gpmgo/gopm

ADD . $GOPATH/src/weibo.com/opendcp/orion

WORKDIR $GOPATH/src/weibo.com/opendcp/orion

RUN go build

RUN scripts/delete_src.sh

EXPOSE 8080

CMD ["./run.sh"]

