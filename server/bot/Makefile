GOPATH:=$(shell go env GOPATH)

.PHONY: run
run:
	PROMETHEUS_ADDR= go run main.go \
		--server_name com.xapis.bot.objforce \
		-apollo_namespace application \
		-apollo_address http://apollo-dev.dev.lucfish.com:8080 \
		-apollo_app_id objforce-bot \
		-apollo_cluster=dev \
		--registry consul \
		--registry_address localhost:8500

.PHONY: build
build: proto
	go build -o objforce-bot *.go

.PHONY: proto
proto:
	make -C proto/objforce-bot go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build -t objforce-bot:latest .

.PHONY: submodule
submodule:
	git submodule init && git submodule update --remote

.PHONY: publish-image
publish-image: docker
	docker tag objforce-bot:latest ccr.ccs.tencentyun.com/mofanshow_services/objforce-bot:latest
	docker push ccr.ccs.tencentyun.com/mofanshow_services/objforce-bot:latest

.PHONY: k8s-deploy
k8s-deploy: publish-image
	kubectl -n xapis apply -f deploy/k8s/app/objforce-bot.yml