FROM golang:alpine

RUN apk add --update bash openssh

RUN mkdir /app
ADD main.go /app/
ADD tls/ /app/tls

WORKDIR /app
VOLUME /app/ssh/

RUN go build -o main .

CMD ["./main"]