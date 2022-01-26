#!/bin/bash

CONTAINER_NAME="arkansas"
cd ../docker/ && docker-compose up -d && docker exec -it ${CONTAINER_NAME} bash
