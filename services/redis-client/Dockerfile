FROM golang:1.16.2-alpine3.12

RUN mkdir /app

ADD ./src /app

WORKDIR /app

RUN go build -o redisClient .

#CMD ["/app/main"]
#
#
#FROM golang:1.16.2-alpine3.12
#
#RUN mkdir /app
#
#ADD ./src /app
#
#WORKDIR /app
#
#RUN go build -o redisClient github.com/cdugga/scaling_with_go/redisclient; \

CMD ["/app/redisClient"]
