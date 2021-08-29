FROM golang:1.17.0

WORKDIR /go/src/github.com/OrderingSystem

COPY . /go/src/github.com/OrderingSystem

CMD [ "./OrderingSystem" ]