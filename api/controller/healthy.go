package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/rongpengju/gin-template/configs"
	"github.com/rongpengju/gin-template/pkg/app"
)

// HealthyCheck k8s 健康服务检查
func HealthyCheck(c *gin.Context) {
	app.NewResponse(c).Success(
		gin.H{
			"health_check_message": fmt.Sprintf("%s is healthy", configs.Conf.App.Name),
		},
	)
}
