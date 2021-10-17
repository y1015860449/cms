GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o advert *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build
	docker build . -t advert:latest

.PHONY: run-docker
run-docker:
	docker build . -t advert:latest
	docker rm -f advert | echo "remove ok"
	docker run -d --name advert advert
	docker ps