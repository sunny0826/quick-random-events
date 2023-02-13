BINARY=qres
VERSION=dev
COMMIT=$(shell git rev-parse HEAD)
BUILD_FLAGS=-ldflags "-X 'github.com/sunny0826/quick-random-events/cmd.Version=$(VERSION)' \
					 -X 'github.com/sunny0826/quick-random-events/cmd.Commit=$(COMMIT)' \
                     -X 'github.com/sunny0826/quick-random-events/cmd.BuildDate=$(shell date -u '+%Y-%m-%dT%H:%M:%SZ')' \
                     -X 'github.com/sunny0826/quick-random-events/cmd.GoVersion=$(shell go version)' \
                     -X 'github.com/sunny0826/quick-random-events/cmd.OSArch=$(shell go env GOOS)/$(shell go env GOARCH)'"

all: build

build:
	go build $(BUILD_FLAGS) -o bin/$(BINARY) main.go

clean:
	rm bin/$(BINARY)

run: build
	bin/$(BINARY)
