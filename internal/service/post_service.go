package service

import (
	"context"
	"errors"
	"go-blog/internal/model"
	"go-blog/internal/repository"
	"go-blog/pkg/errno"
	"go-blog/pkg/logger"
	"go.uber.org/zap"
)

type PostService struct {
	repo *repository.PostRepo
}

func NewPostService(repo *repository.PostRepo) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) Create(ctx context.Context, post *model.Post) error {
	return s.repo.Create(ctx, post)
}

func (s *PostService) HotList(ctx context.Context) ([]model.Post, *errno.Error) {
	log := logger.FromContext(ctx)

	var (
		list []model.Post
		err  error
	)

	list, err = s.repo.ListHot(ctx, 10)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Warn("hot list timeout, fail fast")
			return nil, errno.Timeout
		}
		log.Error("hot list failed", zap.Error(err))
		return nil, errno.InternalServerError
	}
	return list, errno.OK
}
