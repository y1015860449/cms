GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o gateway *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build
	docker build . -t gateway:latest

.PHONY: run-docker
run-docker:
	docker build . -t gateway:latest
	docker rm -f gateway | echo "remove ok"
	docker run -d --name gateway gateway
	docker ps