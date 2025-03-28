package manager

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/rongpengju/gin-template/pkg/middleware"
)

// RegisterManagerRouters 注册后台管理相关接口
func RegisterManagerRouters(rg *gin.RouterGroup) {
	managerRouterGroup := rg.Group("/manager")
	managerRouterGroup.Use(middleware.AuthUserId())

	v1Group := managerRouterGroup.Group("/v1")
	fmt.Println(v1Group.BasePath()) // TODO 如不需要，可以删除
	{
		// 下面书写自己的路由逻辑
	}
}
