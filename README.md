# go-grpc-demo

## 官方文档
[https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code](https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code)

## 仓库地址
[https://github.com/grpc/grpc-go/tree/master/examples/helloworld](https://github.com/grpc/grpc-go/tree/master/examples/helloworld)

## 执行命令

```shell
go run server/server.go
```

```shell
go run client.go -name=你好
```
## 效果

![HZWgv3-EzxWlL](http://upic.renzhifan.cn/uPic/HZWgv3-EzxWlL.png)

## Regenerate gRPC code

```shell
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./proto/helloworld.proto
```






