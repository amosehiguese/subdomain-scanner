.PHONY: test
test:
  go test -v ./...
.PHONY: build
build:
	go build -o bin/frontend
.PHONY: run
run: build
	./bin/frontend
