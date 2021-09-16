# systemsgo
Ongoing project featuring microservice architecture, event sourcing pattern.

##Project Structure
```
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

##To Install
```
git clone https://github.com/ldenholm/systemsgo
```

##Build Docker image and run Docker container
```
cd product-service/
docker build -t #imageName .
docker run -p #yourPort:9090 #imageName
```

##To run without Docker
```
go run main.go
```
