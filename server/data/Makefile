GOPATH:=$(shell go env GOPATH)

.PHONY: run
run:
	MICRO_SERVICE_NAME=com.xapis.api.data \
	APOLLO_NAMESPACE=application \
	APOLLO_ADDRESS=http://apollo-dev.dev.lucfish.com:8080 \
	APOLLO_APP_ID=data-srv \
	APOLLO_CLUSTER=dev \
	REGISTRY=consul \
  REGISTRY_ADDRESS=localhost:8500 \
	go run main.go

.PHONY: build
build:
	go build -o data-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build -t data-srv:latest .

.PHONY: publish-image
publish-image: docker
	docker tag data-srv:latest objforce/data-srv:latest
	docker push objforce/data-srv:latest

.PHONY: k8s-deploy
k8s-deploy: publish-image
	kubectl -n xapis apply -f deploy/k8s/app/data-srv.yml