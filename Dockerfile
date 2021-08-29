FROM golang:1.17.0

ENV GO111MODULE=on
ENV GOOS="linux"

WORKDIR /go/src/github.com/OrderingSystem

COPY . .

RUN rm -rf /go/src/github.com/OrderingSystem/vendor
RUN go mod download

CMD [ "./OrderingSystem" ]
