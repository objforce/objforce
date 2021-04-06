GOPATH:=$(shell go env GOPATH)

.PHONY: run
run:
	MICRO_SERVICE_NAME=com.xapis.srv.meta \
	APOLLO_NAMESPACE=application \
	APOLLO_ADDRESS=http://apollo-dev.dev.lucfish.com:8080 \
	APOLLO_APP_ID=meta-srv \
	APOLLO_CLUSTER=dev \
	REGISTRY=consul \
  REGISTRY_ADDRESS=localhost:8500 \
	go run main.go

.PHONY: build
build:
	go build -o meta-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build -t meta-srv:latest .

.PHONY: publish-image
publish-image: docker
	docker tag meta-srv:latest objforce/meta-srv:latest
	docker push objforce/meta-srv:latest

.PHONY: k8s-deploy
k8s-deploy: publish-image
	kubectl -n xapis apply -f deploy/k8s/app/meta-srv.yml