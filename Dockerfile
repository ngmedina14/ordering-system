FROM golang:1.17.0

ENV GO111MODULE=on
ENV GOOS="linux"

WORKDIR /go/src/github.com/ngmedina14/OrderingSystem

COPY . .

RUN go mod download

RUN go build -o OrderingSystem .

CMD [ "./OrderingSystem" ]