.PHONY: build

LINTER_VERSION=v1.31.0

default: build

get-linter:
	command -v golangci-lint || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b ${GOPATH}/bin $(LINTER_VERSION)

## lint: download/install golangci-lint and analyse the source code with the configuration in .golangci.yml
lint: get-linter
	golangci-lint run --timeout=5m

unit-tests:
	./scripts/unit-tests.sh

## test-race: run tests with race detection
race-condition-tests:
	go test -v -race ./...

integration-tests:
	./scripts/integration-tests.sh

acceptance-tests:
	./scripts/acceptance-tests.sh

docker-tests:
	docker-compose build
	docker-compose run --rm unit-tests
	docker-compose run --rm integration-tests
	docker-compose run --rm acceptance-tests
	docker-compose down

build: lint unit-tests race-condition-tests integration-tests acceptance-tests

tidy:
	go mod tidy -v

down:
	docker-compose down