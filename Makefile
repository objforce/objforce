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
	go build -o build/meta-api cmd/meta-api/main.go

.PHONY: build-meta-srv
build-meta-srv:
	go build -o build/meta-srv cmd/meta-srv/main.go

.PHONY: build-data-api
build-data-api:
	go build -o build/data-api cmd/data-api/main.go

.PHONY: build-data-srv
build-data-srv:
	go build -o build/data-srv cmd/data-srv/main.go

.PHONY: build-index-api
build-index-api:
	go build -o build/index-api cmd/index-api/main.go

.PHONY: build-index-srv
build-index-srv:
	go build -o build/index-api cmd/index-srv/main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build -t data-api:latest . -f cmd/data-api/Dockerfile
	docker build -t data-srv:latest . -f cmd/data-srv/Dockerfile

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