package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// AddTraceID 注入trace信息, 用于链路追踪/日志追踪
func AddTraceID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取trace_id, 读取不到的话生成一个trace_id给ctx
		traceID := c.Request.Header.Get("trace_id")
		if traceID == "" {
			UUID, _ := uuid.NewV4()
			traceID = UUID.String()
		}
		c.Set("trace_id", traceID)
	}
}
