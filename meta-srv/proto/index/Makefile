GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: go
go:
	rm -rf ./gen-go
	mkdir -p ./gen-go
	protoc --proto_path=${GOPATH}/src:. --govalidators_out=${MODIFY}:./gen-go --micro_out=${MODIFY}:./gen-go --go_out=${MODIFY}:./gen-go *.proto

