build:
		go build -o ./gw/gw -i ./gw/*.go
		go build -o ./logic/logic -i ./logic/*.go
		go build -o ./microservice/microservice -i ./microservice/*.go
		docker build -t gw ./gw
		docker build -t logic ./logic
		docker build -t microservice ./microservice		
		rm ./gw/gw	
		rm ./logic/logic
		rm ./microservice/microservice

run:
		docker-compose up


ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

protobuf:
		protoc -I/usr/local/include -I. -I$(GOPATH)src -I$(GOPATH)src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. proto/logic/example.proto
		protoc -I/usr/local/include -I. -I$(GOPATH)src -I$(GOPATH)src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. proto/logic/example.proto
		protoc -I/usr/local/include -I. -I$(GOPATH)src -I$(GOPATH)src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. proto/logic/example.proto
		protoc -I. --go_out=plugins=grpc:. proto/microservice/micro.proto