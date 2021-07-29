#!/bin/bash -e

docker build -t=cjtest1 .
docker tag cjtest1 quii/cjtest1
docker push quii/cjtest1