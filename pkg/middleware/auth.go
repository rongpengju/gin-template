package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rongpengju/gin-template/pkg/errcode"
)

// AuthJwtToken 验证 JWT Token
func AuthJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			authorization   = c.GetHeader("Authorization")
			tokenInvalidErr = errcode.ErrTokenInvalid
			ginH            = gin.H{
				"code":     tokenInvalidErr.Code(),
				"msg":      tokenInvalidErr.Msg(),
				"trace_id": c.GetString("trace_id"),
			}
		)

		if authorization == "" {
			c.JSON(http.StatusForbidden, ginH)
			c.Abort()
			return
		}

		// 解析 token
		claims, err := ParseJwtToken(authorization)
		if err != nil {
			c.JSON(http.StatusForbidden, ginH)
			c.Abort()
			return
		}

		// 判断 uuid 是否为空
		if claims.Uuid == "" {
			c.JSON(http.StatusForbidden, ginH)
			c.Abort()
			return
		}

		c.Set("uuid", claims.Uuid)
		c.Next()
	}
}
