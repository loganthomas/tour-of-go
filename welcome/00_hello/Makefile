.DEFAULT_GOAL := build

# .PHONY keeps make from getting confused if a directory or file
# in your projects has the same name as one of the listed targets
.PHONY: clean fmt vet build
clean:
		go clean
fmt: clean
		go fmt ./...

vet: fmt
		go vet ./...

build: vet
		go build -o bin/
