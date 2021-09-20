#!/bin/bash
cd ../
docker build -f ci/Dockerfile -t product-service:latest .