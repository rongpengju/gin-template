package app

import (
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/rongpengju/gin-template/library/wechat_work"
	"github.com/rongpengju/gin-template/pkg/errcode"
	"github.com/rongpengju/gin-template/pkg/logger"
)

type response struct {
	ctx               *gin.Context
	Code              int         `json:"code"`
	Msg               string      `json:"msg"`
	RequestId         string      `json:"request_id"`
	Data              interface{} `json:"data,omitempty"`
	Pagination        *Pagination `json:"pagination,omitempty"`        // 分页信息
	robotNotification bool        `json:"robotNotification,omitempty"` // 机器人通知
}

func NewResponse(c *gin.Context) *response {
	return &response{ctx: c}
}

// Success 返回链路成功的响应
func (r *response) Success(data interface{}) {
	if data == nil {
		r.Data = make(map[string]interface{})
	} else {
		r.Data = data
	}
	r.Code = errcode.Success.Code()
	r.Msg = errcode.Success.Msg()
	r.RequestId = r.ctx.GetString("trace_id")
	r.ctx.JSON(errcode.Success.HttpStatusCode(), r)
}

// Error 返回链路失败的响应
func (r *response) Error(err error) {
	newErr := errcode.ErrServer.Clone()
	if !errors.As(err, &newErr) {
		newErr.WithCause(err)
	}
	r.Code = newErr.Code()
	r.Msg = newErr.Msg()
	r.RequestId = r.ctx.GetString("trace_id")
	if r.robotNotification {
		go func() {
			if robotErr := wechat_work.RobotNotification(err.Error()); robotErr != nil {
				logger.Error(r.ctx, "企业微信机器人通知错误", zap.Error(robotErr))
			}
		}()
	}
	r.ctx.JSON(newErr.HttpStatusCode(), r)
}

// SetPagination 设置Response的分页信息
func (r *response) SetPagination(pagination *Pagination) *response {
	r.Pagination = pagination
	return r
}

// SetRobotNotification 设置机器人通知
func (r *response) SetRobotNotification() *response {
	r.robotNotification = true
	return r
}
