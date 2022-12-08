export GO111MODULE ?= on
export GOPROXY ?= https://goproxy.cn,direct

default: server

full:
	swag init && go run ./main.go

server:
	go run ./main.go

swagger:
	swag init

.PHONY: default run server
