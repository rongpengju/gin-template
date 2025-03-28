package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rongpengju/gin-template/pkg/errcode"
)

// AuthGuid 验证C端用户的 guid
func AuthGuid() gin.HandlerFunc {
	return func(c *gin.Context) {
		guid := c.GetHeader("X-XTJ-UID")
		if guid == "" {
			// 返回403状态码
			tokenInvalidErr := errcode.ErrTokenInvalid
			c.JSON(http.StatusForbidden, gin.H{
				"code": tokenInvalidErr.Code(),
				"msg":  tokenInvalidErr.Msg(),
			})
			c.Abort()
			return
		}
		c.Set("guid", guid)
		c.Next()
	}
}

// AuthUserId 验证管理端用户的 id
func AuthUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		managerId := c.GetHeader("X-User-Id")
		if managerId == "" {
			// 返回403状态码
			tokenInvalidErr := errcode.ErrTokenInvalid
			c.JSON(http.StatusForbidden, gin.H{
				"code": tokenInvalidErr.Code(),
				"msg":  tokenInvalidErr.Msg(),
			})
			c.Abort()
			return
		}
		c.Set("mid", managerId)
		c.Next()
	}
}
