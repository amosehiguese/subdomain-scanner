.PHONY: build
build:
	go build -o bin/subd
.PHONY: run
run: build
	./bin/subd