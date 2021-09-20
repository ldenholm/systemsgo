# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

#COPY /src/go.mod /
COPY . .

RUN go mod download

RUN go build -o /product-api

CMD [ "/product-api" ]