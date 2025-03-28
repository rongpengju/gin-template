package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/rongpengju/gin-template/api/router"
	"github.com/rongpengju/gin-template/configs"
	"github.com/rongpengju/gin-template/dal"
	"github.com/rongpengju/gin-template/library"
	"github.com/rongpengju/gin-template/pkg/middleware"
)

func main() {
	gin.SetMode(configs.Conf.App.Env)

	// 初始化 gorm/gen，如不使用则删除
	dal.InitGormGen()

	// 初始化所有第三方库
	library.InitAllLibrary()

	// 初始化路由并启动项目
	r := router.InitRoutes()

	// 启动http服务
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", configs.Conf.App.Port),
		Handler: r,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic("应用服务启动失败" + err.Error())
		}
	}()

	// 优雅退出
	middleware.NewHook().Close(
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			if err := httpServer.Shutdown(ctx); err != nil {
				panic("应用服务关闭失败：" + err.Error())
			}
		},
	)
}
