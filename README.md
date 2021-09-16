# systemsgo
Ongoing project featuring microservice architecture, event sourcing pattern, NATS streaming

##Project Structure
|   docker-compose.yml
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
