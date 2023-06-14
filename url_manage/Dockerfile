FROM golang:1.18-alpine3.16

RUN mkdir /app
WORKDIR /app

COPY go.mod ./

COPY ./main.go .

RUN ls
RUN go build -o main ./

CMD sleep 15 && /app/main