package handler

import (
	"github.com/gin-gonic/gin"
	"go-blog/internal/model"
	"go-blog/internal/response"
	"go-blog/internal/service"
	"go-blog/pkg/errno"
	"go-blog/pkg/logger"
)

type PostHander struct {
	svc *service.PostService
}

func NewPostHander(svc *service.PostService) *PostHander {
	return &PostHander{svc: svc}
}

func (h *PostHander) Create(c *gin.Context) {
	log := logger.Ctx(c)
	log.Info("create post")
	// 直接使用model.Post{}
	post := model.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		response.JSON(c, errno.InvalidParams, nil)
		return
	}

	e := h.svc.Create(c.Request.Context(), &post)
	if e != nil {
		response.JSON(c, errno.InternalServerError, nil)
		return
	}
	response.JSON(c, errno.OK, nil)
}

func (s *PostHander) HotList(c *gin.Context) {

	list, err := s.svc.HotList(c.Request.Context())
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.JSON(c, errno.OK, list)
}
