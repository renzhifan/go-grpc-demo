// Package routes 注册路由
package routes

import (
	"helloworld/app/http/controllers/api/v1/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.HellowordController)
			authGroup.POST("/SayHello", suc.SayHello)

			puc := new(auth.PassportController)
			authGroup.POST("/GetMemberProfileByUid", puc.GetMemberProfileByUid)
		}
	}
}
