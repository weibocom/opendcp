FROM golang:latest

RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' > /etc/timezone

ADD . $GOPATH/src/weibo.com/opendcp/jupiter

WORKDIR $GOPATH/src/weibo.com/opendcp/jupiter

RUN mkdir keys

RUN go build

RUN scripts/delete_src.sh

EXPOSE 8080

CMD ["./run.sh"]
