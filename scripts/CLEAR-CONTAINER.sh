#!/bin/bash

cd ../docker/ && docker-compose down && docker rmi `docker images -q`