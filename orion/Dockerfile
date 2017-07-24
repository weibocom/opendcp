FROM golang:latest

RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' > /etc/timezone

RUN go get github.com/davecgh/go-spew/spew && go get github.com/robfig/cron && go get github.com/stretchr/testify/assert && go get github.com/lrita/gosync

RUN mkdir -p $GOPATH/src/golang.org/x && cd $GOPATH/src/golang.org/x &&  git clone https://github.com/golang/sync.git

RUN go get -u github.com/gpmgo/gopm

ADD . $GOPATH/src/weibo.com/opendcp/orion

WORKDIR $GOPATH/src/weibo.com/opendcp/orion

RUN go build

RUN scripts/delete_src.sh

EXPOSE 8080

CMD ["./run.sh"]

