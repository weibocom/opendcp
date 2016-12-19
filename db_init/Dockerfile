FROM library/mysql

RUN apt-get update
RUN apt-get install -y netcat

ADD . /

CMD ["./run.sh"]
