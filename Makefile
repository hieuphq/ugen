.PHONY: build run install

build:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o ./bin/ugen ./ugen.go

release:
	$(foreach os, (darwin linux windows), \
		$(foreach arch, (386 amd64), \
			GOOS=${os} GOARCH=${arch} CGO_ENABLED=0 go build -v -o ./bin/ugen-${os}-${arch}/ugen; \
		)\
	)

run:
	go run ./ugen.go

install:
	go install ./ugen.go