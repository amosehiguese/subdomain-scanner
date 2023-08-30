TAG ?= 1.0.0
IMAGENAME ?= amosehiguese/subd

.PHONY: build
build:
	go build -o bin/subd
.PHONY: run
run: build
	./bin/subd
.PHONY: docker-build
docker-build:
	docker build -t $(IMAGENAME):$(TAG) .