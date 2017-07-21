FROM registry.cn-beijing.aliyuncs.com/opendcp/golang-env:latest

ADD . $GOPATH/src/weibo.com/opendcp/jupiter

WORKDIR $GOPATH/src/weibo.com/opendcp/jupiter

RUN mkdir keys

RUN go build

# RUN scripts/delete_src.sh

EXPOSE 8080

CMD ["./run.sh"]
