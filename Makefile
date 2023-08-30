TAG ?= 1.0.0

.PHONY: build
build:
	go build -o bin/subd
.PHONY: run
run: build
	./bin/subd
.PHONY: docker-build
docker-build:
	docker build -t amosehiguese/subd:$(TAG) .