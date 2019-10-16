GO=go

.PHONY: all run fmt

all: build run

build:
	@go build -o build/bin/gomsg

run:
	build/bin/gomsg

fmt:
	@gofmt -l -w . pkg

clean:
	rm -rf build
