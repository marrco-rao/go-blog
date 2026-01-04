package logger

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// loggerKey 防冲突
type loggerKeyType struct{}

var loggerKey = loggerKeyType{}

// WithContext 把 logger 放入 context
func WithContext(ctx context.Context, log *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, log)
}

// 从 Gin Context 中提取日志对象
func FromContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return Log
	}
	if log, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return log
	}
	return Log
}

// Ctx 从 gin.Context 中获取 logger
func Ctx(c *gin.Context) *zap.Logger {
	if c == nil {
		return Log
	}
	return FromContext(c.Request.Context())
}
