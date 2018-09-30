# godash
The [Open Dashboard Project](https://github.com/tizz98/opendash) rewritten in Golang!

## Getting started

```bash
go get -u github.com/tizz98/godash
cd $GOPATH/src/github.com/tizz98/godash

cp config.example.yaml config.yaml

createuser dash -W -E
createdb godash --owner dash

make migrate
```
