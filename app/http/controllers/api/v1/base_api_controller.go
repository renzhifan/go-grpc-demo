package v1

import "flag"

var (
	Addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

// BaseAPIController 基础控制器
type BaseAPIController struct {
}
