# systemsgo

Ongoing project featuring microservice architecture, event sourcing pattern.

## To Install

```sh
git clone https://github.com/ldenholm/systemsgo
```

## Run local NATS

```sh
docker run -p 4222:4222 -ti nats:latest
```

## Build Docker image and run Docker container

```sh
cd product-service/
docker build -t #imageName .
docker run -p #yourPort:9090 #imageName
```

## To run without Docker

```sh
go run main.go
```

## Project Structure

```txt

|   docker-compose.yml
|   readme.md
|
\---product-service
    |   dockerfile
    |   go.mod
    |   go.sum
    |   main.go
    |
    +---data
    |       products.go
    |
    \---handlers
            default.go
            goodbye.go
            product.go
```

## Areas to explore

circuit breaker, redis cache, ~~dynamodb integration~~,
service discovery, transports protobuf/gRPC, ci (docker compose + ecs)

### CI Things

Multi-step process

![image](https://user-images.githubusercontent.com/47731607/134104063-1208a7c7-2b65-4210-ae95-64c6dccf8e36.png)



_Example build pipeline for a gRPC service_

![image](https://user-images.githubusercontent.com/47731607/134103895-49543677-2391-4947-a1f1-2b24b3f12c12.png)



Glide: Vendor Package Management for Golang: https://github.com/Masterminds/glide
