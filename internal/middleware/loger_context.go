package middleware

import (
	"github.com/gin-gonic/gin"
	"go-blog/pkg/logger"
	"go.uber.org/zap"
)

func LoggerWithContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid, _ := c.Get(RequestIDKey)
		log := logger.Log.With(
			zap.String(RequestIDKey, rid.(string)),
			zap.String("path", c.FullPath()),
		)

		// ⭐ 把 logger 注入到 request.Context
		ctx := logger.WithContext(c.Request.Context(), log)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
