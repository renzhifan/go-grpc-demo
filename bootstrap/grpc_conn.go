package bootstrap

import (
	"flag"
	"helloworld/pkg/grpcconn"
)

func SetupGrpcConnect() {

	// 建立 Grpc 连接
	addr := flag.String("addr", "localhost:50051", "the address to connect to")
	grpcconn.Connect(*addr)
}
