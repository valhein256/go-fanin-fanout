SERVICE := go
STAGE := stg
PLATFORM := local
WORKSPACE := /opt/go
APP := fanin.go

.PHONY: build run devenv update clean

build: clean
	@echo "Start building..."
	@DOCKER_BUILDKIT=1 docker build --pull . \
		--platform ${PLATFORM} \
		--build-arg TAG=${STAGE} \
		--target release --output bin/

run:
	@docker run \
		-v ${PWD}:${WORKSPACE} \
		--rm -it golang:1-buster \
		go run ${WORKSPACE}/${APP}

devenv:
	@docker run \
		-v ${PWD}:/opt/go \
		--rm -it golang:1-buster /bin/bash

update:
	@go mod tidy

clean:
	@rm -rf bin
