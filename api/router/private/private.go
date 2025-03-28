package private

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/rongpengju/gin-template/pkg/middleware"
)

// RegisterPrivateRouters 注册认证相关接口
func RegisterPrivateRouters(rg *gin.RouterGroup) {
	privateRouterGroup := rg.Group("/private")
	privateRouterGroup.Use(middleware.AuthJwtToken())

	v1Group := privateRouterGroup.Group("/v1")
	fmt.Println(v1Group.BasePath()) // TODO 如不需要，可以删除
	{
		// 下面书写自己的路由逻辑
	}
}
