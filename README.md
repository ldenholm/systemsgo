# systemsgo

Ongoing project featuring microservice architecture, event sourcing pattern.

## To Install

```sh
git clone https://github.com/ldenholm/systemsgo
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

Circuit breaker, redis cache, ~~dynamodb integration~~,
service discovery, kubernetes, gRPC, ES, Docker, NgInx.
