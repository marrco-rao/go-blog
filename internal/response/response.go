package response

import (
	"github.com/gin-gonic/gin"
	"go-blog/pkg/errno"
)

// 统一返回结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(errno.OK.HTTP, Response{
		Code:    errno.OK.Code,
		Message: errno.OK.Message,
		Data:    data,
	})
}

// 失败（核心）
func Fail(c *gin.Context, err *errno.Error) Response {
	return Response{
		Code:    err.Code,
		Message: err.Message,
	}
}

// 统一输出（可选）
func JSON(c *gin.Context, err *errno.Error, data interface{}) {
	if err == errno.OK {
		Success(c, data)
		return
	}

	c.JSON(err.HTTP, Response{
		Code:    err.Code,
		Message: err.Message,
	})
}
