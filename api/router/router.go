package router

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/rongpengju/gin-template/api/router/manager"
	"github.com/rongpengju/gin-template/api/router/private"
	"github.com/rongpengju/gin-template/api/router/public"
	"github.com/rongpengju/gin-template/configs"
	"github.com/rongpengju/gin-template/pkg/middleware"
)

func InitRoutes() *gin.Engine {
	engine := gin.Default()

	// 禁用代理检查
	if err := engine.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	engine = registerMiddleware(engine)

	engine = registerRoutes(engine)

	return engine
}

// 注册中间件
func registerMiddleware(engine *gin.Engine) *gin.Engine {
	// 加入链路追踪
	engine.Use(middleware.AddTraceID())

	// 跨域中间件
	engine.Use(middleware.Cors())

	// 记录 请求值 和 响应值 中间件
	engine.Use(middleware.LogReqAndResp())

	// panic recover 输出
	engine.Use(middleware.PanicRecovery())

	return engine
}

// 注册路由
func registerRoutes(engine *gin.Engine) *gin.Engine {
	routerGroup := engine.Group(fmt.Sprintf("/%s", configs.Conf.App.Name))

	registerHealthyRoutes(routerGroup)

	public.RegisterPublicRouters(routerGroup)

	private.RegisterPrivateRouters(routerGroup)

	manager.RegisterManagerRouters(routerGroup)

	return engine
}
