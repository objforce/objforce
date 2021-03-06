GOPATH:=$(shell go env GOPATH)

.PHONY: run
run:
	go run main.go \
		--server_name=com.xapis.api.meta \
		--apollo_namespace=application \
		--apollo_address=http://apollo-dev.dev.lucfish.com:8080 \
		--apollo_app_id=meta-api \
		--apollo_cluster=dev \
		--registry=consul \
        --registry_address=localhost:8500

.PHONY: build
build:
	go build -o meta-api *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build -t meta-api:latest .

.PHONY: submodule
submodule:
	git submodule init && git submodule update --remote

.PHONY: publish-image
publish-image: docker
	docker tag meta-api:latest objforce/meta-api:latest
	docker push objforce/meta-api:latest

.PHONY: k8s-deploy
k8s-deploy: publish-image
	kubectl -n xapis apply -f deploy/k8s/app/meta-api.yml