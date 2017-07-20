FROM registry.cn-beijing.aliyuncs.com/opendcp/golang-env:latest

RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' > /etc/timezone

ADD . $GOPATH/src/weibo.com/opendcp/orion

WORKDIR $GOPATH/src/weibo.com/opendcp/orion

RUN go build

RUN scripts/delete_src.sh

EXPOSE 8080

CMD ["./run.sh"]

