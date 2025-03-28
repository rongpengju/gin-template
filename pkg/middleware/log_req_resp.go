package middleware

import (
	"bytes"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/rongpengju/gin-template/pkg/logger"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 包装 gin.ResponseWriter，通过这种方式拦截写响应
// 让 gin 写响应的时候先写到 bodyLogWriter 再写gin.ResponseWriter
// 这样利用中间件里输出访问日志时就能拿到响应了
// 参考地址：https://stackoverflow.com/questions/38501325/how-to-log-response-body-in-gin
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// LogReqAndResp 记录请求内容和响应内容
func LogReqAndResp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			reqBody     = make([]byte, 0)
			contentType = c.GetHeader("Content-Type")
		)
		// multipart/form-data 不打印 请求体内容
		if !strings.Contains(contentType, "multipart/form-data") {
			reqBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewReader(reqBody))
		}
		startTime := time.Now()
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		logReqAndResp(c, time.Since(startTime), reqBody, nil)
		defer func() {
			var responseLogging string
			if c.Writer.Size() > 100*1024 { // 响应大于100KB 不记录
				responseLogging = "响应数据大于100kb"
			} else {
				responseLogging = blw.body.String()
			}
			logReqAndResp(c, time.Since(startTime), reqBody, responseLogging)
		}()
		c.Next()

		return
	}
}

func logReqAndResp(c *gin.Context, dur time.Duration, body []byte, responseData interface{}) {
	if responseData == nil {
		req := c.Request
		logger.Info(c, "RequestLog",
			zap.String("request_ip", c.ClientIP()),
			zap.String("request_method", req.Method),
			zap.String("request_path", req.URL.Path),
			zap.String("request_query", req.URL.RawQuery),
			zap.String("request_body", string(body)),
			zap.String("request_token", req.Header.Get("Authorization")),
		)
	} else {
		logger.Info(c, "ResponseLog",
			zap.Any("response_data", responseData),
			zap.Int64("response_time(ms)", int64(dur/time.Millisecond)),
		)
	}
}
