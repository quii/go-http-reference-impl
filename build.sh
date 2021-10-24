#!/bin/bash -e

echo "-- Lint --"
#golangci-lint run

echo "-- Unit tests --"
./scripts/unit-tests.sh

echo "-- Integration tests --"
./scripts/integration-tests.sh

echo "-- Acceptance tests --"
./scripts/acceptance-tests.sh
