dep:
		dep ensure

build: dep
		go build .

fmt: dep
		go fmt ./...

test: build
		go test -cover -race -v ./...
