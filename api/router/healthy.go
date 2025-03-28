package router

import (
	"github.com/gin-gonic/gin"

	"github.com/rongpengju/gin-template/api/controller"
)

// 健康检查模块路由
func registerHealthyRoutes(rg *gin.RouterGroup) {
	rg.GET("/healthy", controller.HealthyCheck)
}
