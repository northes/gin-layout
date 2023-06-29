VERSION=$(shell git describe --tags --always)
export CGO_ENABLED=1
ORM="ent"

.PHONY: gen
gen:
	go generate ./...

.PHONY: dev
dev:
	go build --tags $(ORM) -ldflags "-X main.Version=$(VERSION)" -o bin/dev && ./bin/dev

.PHONY: build
build:
	mkdir -p bin/ && go build --tags $(ORM) -ldflags "-X main.Version=$(VERSION)" -o ./bin/app

.PHONY: clear
clear:
	rm -rf bin && rm -rf *.log && rm -rf *.db