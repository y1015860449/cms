GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o user *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build
	docker build . -t user:latest

.PHONY: run-docker
run-docker:
	docker build . -t user:latest
	docker rm -f user | echo "remove ok"
	docker run -d --name user user
	docker ps