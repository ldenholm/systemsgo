#!/usr/bin/bash

# test post endpoint with dummy product
POST_HTTP_CODE=$(curl --write-out "%{http_code}\n" "localhost:9090/" -X POST -d '{"name": "test-script", "price": 5, "sku": "abc-absd-dfsdf"}' --output output.txt)
echo $HTTP_CODE 

# test get products endpoint
GET_RESPONSE=$(curl --write-out "status code: %{http_code}\n" "localhost:9090/products" --output get_response.txt)
echo $GET_RESPONSE