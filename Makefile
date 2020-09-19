GOPATH:=$(shell go env GOPATH)

.PHONY: build
build: build-meta-api
	make -C cmd/bot build
	make -C cmd/meta-srv build
	make -C cmd/data-api build
	make -C cmd/data-srv build
	make -C cmd/index-api build
	make -C cmd/index-srv build

.PHONY: build-meta-api
build-meta-api:
	go build -o meta-api cmd/meta-api/main.go


.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	make -C cmd/bot docker
	make -C cmd/meta-api docker
	make -C cmd/meta-srv docker
	make -C cmd/data-api docker
	make -C cmd/data-srv docker
	make -C cmd/index-api docker
	make -C cmd/index-srv docker


.PHONY: publish-image
publish-image: docker
	make -C cmd/bot publish-image
	make -C cmd/meta-api publish-image
	make -C cmd/meta-srv publish-image
	make -C cmd/data-api publish-image
	make -C cmd/data-srv publish-image
	make -C cmd/index-api publish-image
	make -C cmd/index-srv publish-image

.PHONY: k8s-deploy
k8s-deploy: publish-image
	make -C cmd/bot k8s-deploy
	make -C cmd/meta-api k8s-deploy
	make -C cmd/meta-srv k8s-deploy
	make -C cmd/data-api k8s-deploy
	make -C cmd/data-srv k8s-deploy
	make -C cmd/index-api k8s-deploy
	make -C cmd/index-srv k8s-deploy