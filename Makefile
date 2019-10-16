GO=go

.PHONY: all run fmt build-image

all: clean build run

build:
	@go build -o build/bin/gomsg

build-image:
	@docker build -t mcyprian/gomsg:latest .
	@docker push mcyprian/gomsg:latest

run:
	build/bin/gomsg

fmt:
	@gofmt -l -w . pkg

clean:
	rm -rf build
