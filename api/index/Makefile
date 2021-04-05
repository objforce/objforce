GOPATH:=$(shell go env GOPATH)

.PHONY: run
run:
	go run main.go \
		--server_name=com.xapis.api.index \
		--apollo_namespace=application \
		--apollo_address=http://apollo-dev.dev.lucfish.com:8080 \
		--apollo_app_id=index-api \
		--apollo_cluster=dev \
		--registry=consul \
        --registry_address=localhost:8500

.PHONY: docker
docker:
	docker build -t index-api:latest .

.PHONY: submodule
submodule:
	git submodule init && git submodule update --remote

.PHONY: publish-image
publish-image: docker
	docker tag index-api:latest objforce/index-api:latest
	docker push objforce/index-api:latest

.PHONY: k8s-deploy
k8s-deploy: publish-image
	kubectl -n xapis apply -f deploy/k8s/app/index-api.yml