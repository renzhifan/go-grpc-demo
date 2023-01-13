package auth

import (
	"context"
	"fmt"
	v1 "helloworld/app/http/controllers/api/v1"
	"helloworld/pkg/grpcconn"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	pb "helloworld/proto"
)

type HellowordController struct {
	v1.BaseAPIController
}

func (sc *HellowordController) SayHello(gctx *gin.Context) {

	// 请求对象
	type HelloRequest struct {
		Name string `json:"name"`
	}
	request := HelloRequest{}

	// 解析 JSON 请求
	if err := gctx.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回 422 状态码和错误信息
		gctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return
	}

	defer grpcconn.Connect().Close()
	c := pb.NewGreeterClient(grpcconn.Connect())

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: request.Name})
	if err != nil {
		fmt.Print("could not greet:", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	//  检查数据库并返回响应
	gctx.JSON(http.StatusOK, gin.H{
		"exist": r.GetMessage(),
	})
}
