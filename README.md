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

## 测试
启动服务
```shell
go run server/server.go
```
客户端发送请求

```shell
curl --location --request POST 'http://localhost:3000/v1/auth/SayHello' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "北京你好"
}'
```
返回数据
```json
{"exist":"Hello 北京你好"}
```

```shell
curl --location --request POST 'http://localhost:3000/v1/auth/GetMemberProfileByUid' \
--header 'Content-Type: application/json' \
--data-raw '{
    "uid": 11333,
    "option":11
}'
```
返回数据
```json
{"nick":"hhh","uid":11333}
```

# 如何调试gRPC

## grpcurl-命令行工具

类似curl，可以直接用来发送请求调用远程的grpc服务，官方地址请看[grpcurl](https://github.com/fullstorydev/grpcurl)

### 安装

#### Homebrew (macOS)

On macOS, grpcurl is available via Homebrew:

```shell
brew install grpcurl
```
#### Docker

For platforms that support Docker, you can download an image that lets you run grpcurl:

```shell
# Download image
docker pull fullstorydev/grpcurl:latest
# Run the tool
docker run fullstorydev/grpcurl api.grpc.me:443 list
```

### 使用
>注意grpcl是基于反射，需要在启动服务前添加这样一行代码

```shell
reflection.Register(s)

```

#### 1.查询服务列表

```shell
☁  ~  grpcurl -plaintext 127.0.0.1:50051 list

grpc.reflection.v1alpha.ServerReflection
proto.Greeter
proto.Member
```
#### 2.查询服务提供的方法
```shell
☁  ~  grpcurl -plaintext 127.0.0.1:50051 list proto.Greeter

proto.Greeter.SayHello
```

#### 3.查看更详细的描述

```shell
☁  ~  grpcurl -plaintext 127.0.0.1:50051 describe proto.Greeter
proto.Greeter is a service:
service Greeter {
  rpc SayHello ( .proto.HelloRequest ) returns ( .proto.HelloReply );
}
```
#### 4.其他的方法参考文档

## grpcui-界面工具

[官方地址grpcui](https://github.com/fullstorydev/grpcui)

### 安装

```go
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
```

### 使用

运行web界面，指定grpc的地址
```shell
grpcui -plaintext localhost:50051
```
>注意 这里的50051  是你的grpc服务指定的端口号（另外 确定你在grpc服务启动前已添加这行代码即可 reflection.Register(s)）

直接把所有服务和方法都列出来了，真的是相当人性化

![11-16-57-la1yot-tTInAG](https://raw.githubusercontent.com/renzhifan/upic_img/master/uPic/2023/01/18/11-16-57-la1yot-tTInAG.png)