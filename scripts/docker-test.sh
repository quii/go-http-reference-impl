#!/bin/sh -e

# this is failing in gh action world, not sure why yet
#trap 'catch' ERR
#catch() {
#  docker-compose down
#}

docker-compose build
docker-compose run --rm unit-tests
docker-compose run --rm integration-tests
docker-compose run --rm acceptance-tests
docker-compose down