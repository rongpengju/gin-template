package dal

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"github.com/rongpengju/gin-template/pkg/logger"
)

var _ gormLogger.Interface = (*customGormLogger)(nil)

type customGormLogger struct {
	SlowThreshold time.Duration
}

func newGormLogger() *customGormLogger {
	return &customGormLogger{
		SlowThreshold: 500 * time.Millisecond, // 超过500毫秒算慢查询, 如果需要按环境定制化, 请写到配置文件中
	}
}

func (l *customGormLogger) LogMode(lev gormLogger.LogLevel) gormLogger.Interface {
	return &customGormLogger{}
}
func (l *customGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	logger.Info(ctx, msg, zap.Any("data", data))
}
func (l *customGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	logger.Error(ctx, msg, zap.Any("data", data))
}
func (l *customGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	logger.Error(ctx, msg, zap.Any("data", data))
}
func (l *customGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	// 获取运行时间
	duration := time.Since(begin).Milliseconds()

	// 获取 SQL 语句和返回条数
	sql, rows := fc()

	// Gorm 错误时记录错误日志
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Error(ctx, "SQL_ERROR", zap.String("SQL", sql), zap.Int64("Rows", rows), zap.Int64("Dru(ms)", duration))
	}

	// 慢查询日志
	if duration > l.SlowThreshold.Milliseconds() {
		logger.Warn(ctx, "SQL_SLOW", zap.String("SQL", sql), zap.Int64("Rows", rows), zap.Int64("Dru(ms)", duration))
	} else {
		logger.Debug(ctx, "SQL_DEBUG", zap.String("SQL", sql), zap.Int64("Rows", rows), zap.Int64("Dru(ms)", duration))
	}
}
