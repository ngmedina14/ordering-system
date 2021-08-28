FROM golang:1.17.0-alpine3.14

WORKDIR /go/src/github.com/ngmedina14/OrderingSystem

COPY . /go/src/github.com/ngmedina14/OrderingSystem

CMD [ "./OrderingSystem" ]