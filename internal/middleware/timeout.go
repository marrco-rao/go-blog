package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func TimeoutWithRoute(defaultTimeout time.Duration, rules map[string]time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		timeout := defaultTimeout

		if t, ok := rules[c.FullPath()]; ok {
			timeout = t
		}
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
