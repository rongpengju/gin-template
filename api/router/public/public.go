package public

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// RegisterPublicRouters 注册非认证相关接口
func RegisterPublicRouters(rg *gin.RouterGroup) {
	publicRouterGroup := rg.Group("/public")

	v1Group := publicRouterGroup.Group("/v1")
	fmt.Println(v1Group.BasePath()) // TODO 如不需要，可以删除
	{
		// 下面书写自己的路由逻辑
	}
}
