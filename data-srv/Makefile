GOPATH:=$(shell go env GOPATH)

.PHONY: run
run:
	go run main.go \
		--server_name=com.xapis.api.data \
		--apollo_namespace=application \
		--apollo_address=http://apollo-dev.dev.lucfish.com:8080 \
		--apollo_app_id=data-server \
		--apollo_cluster=dev \
		--registry=consul \
        --registry_address=localhost:8500

.PHONY: build
build:
	go build -o data-server *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build -t data-server:latest .

.PHONY: submodule
submodule:
	git submodule init && git submodule update --remote

.PHONY: publish-image
publish-image: docker
	docker tag data-server:latest objforce/data-server:latest
	docker push objforce/data-server:latest

.PHONY: k8s-deploy
k8s-deploy: publish-image
	kubectl -n xapis apply -f deploy/k8s/app/data-server.yml