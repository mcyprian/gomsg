.PHONY: all run fmt build-image test

all: clean build run

build:
	@go build -o build/bin/gomsg

build-image:
	@docker build -t mcyprian/gomsg:latest .
	@docker push mcyprian/gomsg:latest

run:
	build/bin/gomsg

test:
	@go test -v ./pkg/...

fmt:
	@gofmt -l -w . pkg

clean:
	rm -rf build
