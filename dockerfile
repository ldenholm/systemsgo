# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY /src/go.mod /

RUN go mod download

COPY /src/*.go ./

RUN go build -o /product-api

CMD [ "/product-api" ]