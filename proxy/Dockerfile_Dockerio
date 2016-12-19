FROM nginx:latest

RUN apt-get update
RUN apt-get install -y php5-common php5-cli php5-fpm php5-mysql php5-dev redis-server vim
RUN apt-get install -y libapache2-mod-php5
RUN apt-get install -y php5-ldap php5-curl
RUN apt-get install -y rsync

ADD ./phpredis /phpredis
WORKDIR /phpredis
RUN phpize
RUN ./configure
RUN make && make install

WORKDIR /
ADD ./run.sh run.sh

ENTRYPOINT ["./run.sh"]
