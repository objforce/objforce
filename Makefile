GOPATH:=$(shell go env GOPATH)

.PHONY: build
build: build-meta-api build-meta-srv build-data-api build-data-srv build-index-api build-index-srv

.PHONY: build-meta-api
build-meta-api:
	go build -o build/$(ARCH)/meta-api cmd/meta-api/main.go

.PHONY: build-meta-srv
build-meta-srv:
	go build -o build/$(ARCH)/meta-srv cmd/meta-srv/main.go

.PHONY: build-data-api
build-data-api:
	go build -o build/$(ARCH)/data-api cmd/data-api/main.go

.PHONY: build-data-srv
build-data-srv:
	go build -o build/$(ARCH)/data-srv cmd/data-srv/main.go

.PHONY: build-index-api
build-index-api:
	go build -o build/$(ARCH)/index-api cmd/index-api/main.go

.PHONY: build-index-srv
build-index-srv:
	go build -o build/$(ARCH)/index-api cmd/index-srv/main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker-build
docker-build:
	docker build -t objforce:latest . -f Dockerfile.build
	docker run -v $(PWD)/build:/opt/build objforce:latest make build ARCH=linux

.PHONY: build-image
build-image: docker-build
	docker build -t objforce/meta-api:latest . -f cmd/meta-api/Dockerfile

.PHONY: publish-image
publish-image: build-image
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