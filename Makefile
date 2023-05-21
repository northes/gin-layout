VERSION=$(shell git describe --tags --always)
export CGO_ENABLED=1

.PHONY: gen
gen:
	go generate ./...

.PHONY: dev
dev:
	go build -ldflags "-X main.Version=$(VERSION)" -o bin/dev && ./bin/dev

.PHONY: build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...