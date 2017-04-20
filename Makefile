GO ?= go

# command to build and run on the local OS.
GO_BUILD = go build

# publish this many batches of 10 messages by default
BATCH_COUNT ?= 1000

# command to compiling the distributable. Specify GOOS and GOARCH for
# the target OS.
GO_DIST = GOOS=linux GOARCH=amd64 go build

.PHONY: dist

all: clean build

dist:
	mkdir -p dist
	$(GO_DIST) -o dist/queue-foo-publish cmd/queue-foo-publish/main.go
	$(GO_DIST) -o dist/queue-foo-receive cmd/queue-foo-receive/main.go

build:
	mkdir -p build
	$(GO_BUILD) -o build/queue-foo-publish cmd/queue-foo-publish/main.go
	$(GO_BUILD) -o build/queue-foo-receive cmd/queue-foo-receive/main.go

run-publish:
	BATCH_COUNT=$(BATCH_COUNT) go run cmd/queue-foo-publish/main.go

run-receive:
	go run cmd/queue-foo-receive/main.go

test:
	$(GO) test

docker:
	docker build -t queue-foo-receive .

docker-run:
	docker run --env-file ./conf/docker.env queue-foo-receive

clean:
	rm -rf build dist
