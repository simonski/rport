default_target: build
.PHONY : default_target publish

usage:
	@echo "The rport Makefile"
	@echo ""
	@echo "Usage : make <command> "
	@echo ""
	@echo "commands"
	@echo ""
	@echo "  clean                 - run go clean"
	@echo "  test                  - runs go test"
	@echo "  build                 - creates native binary"
	@echo ""

clean:
	go clean

setup:
	go install honnef.co/go/tools/cmd/staticcheck@latest

build:
	staticcheck ./...
	go fmt .
	go build -o rport -ldflags "-s -w"

test:
	go test ./...

install:
	go install
