FROM golang:1.16.2-alpine3.12

RUN mkdir /app

ADD ./src /app

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]