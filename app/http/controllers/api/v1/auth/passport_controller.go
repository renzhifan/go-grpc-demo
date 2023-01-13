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

type PassportController struct {
	v1.BaseAPIController
}

func (sc *PassportController) GetMemberProfileByUid(gctx *gin.Context) {

	// 请求对象
	type MemberRequest struct {
		Uid    uint64 `json:"uid"`
		Option int32  `json:"option"`
	}
	request := MemberRequest{}

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
	c := pb.NewMemberClient(grpcconn.Connect())

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetMemberProfileByUid(ctx, &pb.GetMemberProfileByUidReq{Uid: request.Uid, Option: request.Option})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//  检查数据库并返回响应
	gctx.JSON(http.StatusOK, gin.H{
		"uid":  r.GetUid(),
		"nick": r.GetNick(),
	})
}
