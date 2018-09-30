dep:
		dep ensure

build: dep
		go build .

fmt: dep
		go fmt ./...

test: build
		go test -cover -race -v ./...

migrate: build
		./godash migrate up

db-version: build
		./godash migrate version

migrate-down: build
		./godash migrate down

db-reset: build
		./godash migrate reset
