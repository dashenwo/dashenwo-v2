MAKEFILE_PATH=$(shell pwd)
PB_PATH=protos/${gen}
GEN_PATH=console/${gen}/proto
export MICRO_REGISTRY=etcd
export MICRO_REGISTRY_ADDRESS=49.235.160.203:2379
export MICRO_API_NAMESPACE=com.dashenwo.srv
export MICRO_WEB_NAMESPACE=com.dashenwo.web

.PHONY: proto
proto:
	#编译标签服务protobuf
	rm -rf ${GEN_PATH}/*
	protoc --proto_path=${MAKEFILE_PATH}/${PB_PATH} \
	--micro_out=${GEN_PATH} \
	--go_out=:${GEN_PATH} \
	${MAKEFILE_PATH}/${PB_PATH}/*.proto
	protoc-go-inject-tag -input=${MAKEFILE_PATH}/${GEN_PATH}/${gen}.pb.go

.PHONY: run_web
run_web:
    #编译用户protobuf
	micro web
.PHONY: run_api
run_api:
    #编译用户protobuf
	micro api