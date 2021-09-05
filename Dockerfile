FROM golang:1.17

# RUN apt update
# # RUN apt install -y rsyslog

WORKDIR /go/src/github.com/OrderingSystem

COPY . .

#RUN sed -i '/imklog/s/^/#/' /etc/rsyslog.conf
# RUN service rsyslog start

CMD [ "./OrderingSystem" ]