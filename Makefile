BINARY=qres

all: build

build:
	go build -o bin/$(BINARY) main.go

clean:
	rm bin/$(BINARY)

run: build
	bin/$(BINARY)
