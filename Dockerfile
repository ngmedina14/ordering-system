FROM golang:1.17.0

RUN apt update
RUN apt install -y rsyslog

WORKDIR /go/src/github.com/OrderingSystem

COPY . /go/src/github.com/OrderingSystem

RUN service rsyslog start

CMD [ "./OrderingSystem" ]