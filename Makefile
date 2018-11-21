GOPATH:=$(shell go env GOPATH)


proto:
	protoc --proto_path=${GOPATH}/src:. --go_out=${GOPATH}/src nulls.proto