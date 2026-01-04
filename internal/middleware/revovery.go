package middleware

import (
	"go-blog/internal/response"
	"go-blog/pkg/errno"
	"go-blog/pkg/logger"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 1. 记录 panic 日志（最重要）
				logger.Log.Error("panic recovered",
					zap.Any("error", err),
					zap.String("path", c.Request.URL.Path),
					zap.String("method", c.Request.Method),
					zap.String("request_id", c.GetString("request_id")),
					zap.ByteString("stack", debug.Stack()),
				)

				// 2. 保证 HTTP 返回（避免连接被强制断开）
				c.AbortWithStatusJSON(
					http.StatusInternalServerError,
					response.Fail(c, errno.InternalServerError),
				)
			}
		}()

		// 3. 继续后续 middleware / handler
		c.Next()
	}
}
