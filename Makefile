dep:
		dep ensure

build: dep
		go build -o dash .

fmt: dep
		go fmt ./...

test: dep
		go test -cover -race -v ./...
		go vet ./...

migrate: build
		./godash migrate up

db-version: build
		./godash migrate version

migrate-down: build
		./godash migrate down

db-reset: build
		./godash migrate reset
