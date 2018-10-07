# godash
The [Open Dashboard Project](https://github.com/tizz98/opendash) rewritten in Golang!

[![CircleCI](https://circleci.com/gh/tizz98/godash/tree/master.svg?style=svg)](https://circleci.com/gh/tizz98/godash/tree/master)

## Getting started

```bash
go get -u github.com/tizz98/godash
cd $GOPATH/src/github.com/tizz98/godash

cp config.example.yaml config.yaml

createuser dash -W -E
createdb godash --owner dash

make migrate
```

## Commands
* `make build` creates the binary `dash`
* `make fmt` formats the Go files
* `make test` runs all tests
* `make migrate` runs migrations
* `make db-version` shows the current database version
* `make migrate-down` resets the last migration
* `make db-reset` resets the database
* `./dash dbshell` connects to a PSQL shell
* `./dash serve` starts the HTTP server on port `9090`
