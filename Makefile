.PHONY: build run install

build:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o ./bin/ugen ./cmd/ugen

run:
	go run ./ugen.go

install:
	go install ./ugen.go