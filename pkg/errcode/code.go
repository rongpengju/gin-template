package errcode

import "net/http"

var codes = map[int]struct{}{}

// 此处为公共的错误码, 预留 100000 ~ 100099 间的 100 个错误码
var (
	Success            = newError(0, "Success")
	ErrServer          = newError(100000, "服务器内部错误")
	ErrParams          = newError(100001, "参数无效，请检查后重试")
	ErrNotFound        = newError(100002, "资源未找到")
	ErrPanic           = newError(100003, "系统开小差啦，请稍后重试") // 无预期的panic错误
	ErrTokenInvalid    = newError(100004, "Token无效")
	ErrForbidden       = newError(100005, "未授权") // 访问一些未授权的资源时的错误
	ErrTooManyRequests = newError(100006, "请求频率过快")
)

// 各个业务模块自定义的错误码, 从 100101 开始（10代表服务/01代表该服务下的某个模块/01代表该模块下的某个错误码序号）
// 按照不同的业务模块划分不同的号段
// Example:
//var (
//	ErrUserNotFound  = newError(100101, "用户未找到")
//)

// HttpStatusCode 公共错误码 映射 HTTP状态码
func (e *AppError) HttpStatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ErrServer.Code(), ErrPanic.Code():
		return http.StatusInternalServerError
	case ErrParams.Code():
		return http.StatusBadRequest
	case ErrNotFound.Code():
		return http.StatusNotFound
	case ErrTooManyRequests.Code():
		return http.StatusTooManyRequests
	case ErrTokenInvalid.Code():
		return http.StatusUnauthorized
	case ErrForbidden.Code():
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
