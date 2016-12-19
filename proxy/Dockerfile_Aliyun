FROM registry.cn-beijing.aliyuncs.com/opendcp/proxy-env

ADD ./phpredis /phpredis
WORKDIR /phpredis
RUN phpize
RUN ./configure
RUN make && make install

WORKDIR /
ADD ./run.sh run.sh

ENTRYPOINT ["./run.sh"]
