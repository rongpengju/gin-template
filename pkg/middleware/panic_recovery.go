package middleware

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	subStr1 = "broken pipe"
	subStr2 = "connection reset by peer"
)

// PanicRecovery 自定义 gin recover 输出
func PanicRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var (
					brokenPipe   bool
					opError      *net.OpError
					syscallError *os.SyscallError
				)
				if errors.As(err.(error), &opError) {
					if errors.As(opError.Err, &syscallError) {
						if strings.Contains(strings.ToLower(syscallError.Error()), subStr1) || strings.Contains(strings.ToLower(syscallError.Error()), subStr2) {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					log.Println(fmt.Sprintf("HttpRequest broken pipe, path: %s, error: %s, request: %s", c.Request.URL.Path, err, string(httpRequest)))
					if err = c.Error(err.(error)); err != nil {
						log.Println("c.Error: ", err)
					}
					c.Abort()
					return
				}
				log.Println(fmt.Sprintf("HttpRequest panic, path: %s, error: %s, request: %s, stack: %s", c.Request.URL.Path, err, string(httpRequest), string(debug.Stack())))
				if err = c.AbortWithError(http.StatusInternalServerError, err.(error)); err != nil {
					log.Println("c.AbortWithError: ", err)
				}
			}
		}()
		c.Next()
	}
}
