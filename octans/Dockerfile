FROM python:2.7

RUN mkdir -p /data/octans

COPY ./requirements.txt /data/octans/

RUN pip install -r /data/octans/requirements.txt

COPY . /data/octans

WORKDIR /data/octans

CMD ["./run.sh"]
