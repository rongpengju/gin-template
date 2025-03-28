package logger

import (
	"context"
	"path"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	log(ctx, zapcore.DebugLevel, msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	log(ctx, zapcore.InfoLevel, msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	log(ctx, zapcore.WarnLevel, msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	log(ctx, zapcore.ErrorLevel, msg, fields...)
}

func log(ctx context.Context, lvl zapcore.Level, msg string, fields ...zap.Field) {
	// 日志行信息中增加追踪参数
	var ctxTraceId = ctx.Value("trace_id")
	if ctxTraceId != nil {
		fields = append(fields, zap.Any("trace_id", ctxTraceId.(string)))
	}

	// 增加日志调用者信息, 方便查日志时定位程序位置
	funcName, file, line := getLoggerCallerInfo()
	fields = append(fields, zap.Any("caller_func", funcName))
	fields = append(fields, zap.Any("caller_file", file))
	fields = append(fields, zap.Any("caller_line", line))

	ce := _logger.Check(lvl, msg)
	ce.Write(fields...)
}

// getLoggerCallerInfo 日志调用者信息
// 返回值：方法名, 文件名, 行号
func getLoggerCallerInfo() (funcName, file string, line int) {
	pc, file, line, ok := runtime.Caller(3) // 回溯拿调用日志方法的业务函数的信息
	if !ok {
		return
	}
	file = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	return
}
