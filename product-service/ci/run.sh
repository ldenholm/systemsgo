#!/bin/bash

echo "Listening on localhost:9090"
docker run -p 9090:9090 product-service:latest