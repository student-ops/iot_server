FROM golang:1.18-alpine3.16 AS builder

RUN mkdir /app
WORKDIR /app

COPY ./src/ .

RUN go mod download
RUN go build -o main ./*.go

FROM alpine:3.16

COPY --from=builder /app/main .
CMD ["/main"]