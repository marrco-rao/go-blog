package repository

import (
	"context"
	"errors"
	"go-blog/internal/model"
	"go-blog/pkg/errno"
	"go-blog/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type PostRepo struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) *PostRepo {
	return &PostRepo{db: db}
}

// 创建文章
func (r *PostRepo) Create(ctx context.Context, p *model.Post) error {
	log := logger.FromContext(ctx)
	err := r.db.WithContext(ctx).Create(p).Error
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Warn("create post timeout")
			return errno.Timeout
		}
	}
	return nil
}

// 根据ID查询
func (r *PostRepo) FindByID(ctx context.Context, id uint) (*model.Post, error) {
	var p model.Post
	err := r.db.WithContext(ctx).First(&p, id).Error
	return &p, err
}

// 根据AuthorID查询,后面补充分页
func (r *PostRepo) FindByAuthorID(ctx context.Context, userid uint) ([]model.Post, error) {
	var pList []model.Post
	err := r.db.WithContext(ctx).Where("author_id", userid).Find(&pList).Error
	return pList, err
}

// 热门文章（索引 / SQL / 慢查询入口）
func (r *PostRepo) ListHot(ctx context.Context, limit int) ([]model.Post, error) {
	start := time.Now()

	var list []model.Post
	// views > 100 魔法值待消除
	err := r.db.WithContext(ctx).
		Where("views > ?", 100).Order("views desc").Limit(limit).Find(&list).Error
	//err := r.db.WithContext(ctx).
	//	Table("post").
	//	Select("id, title, views").
	//	Where("views >= ?", 1000).
	//	Order("views DESC").
	//	Limit(limit).
	//	Find(&list).Error
	cost := time.Since(start)
	if cost > 100*time.Millisecond {
		log := logger.FromContext(ctx)
		fields := []zap.Field{
			zap.String("op", "hot_list"),
			zap.String("table", "post"),
			zap.Int("limit", limit),
			zap.Duration("cost", cost),
		}
		if err != nil {
			fields = append(fields, zap.Error(err))
			log.Error("repo hot list failed", fields...)
		} else {
			log.Warn("slow sql", fields...)
		}
	}
	return list, err
}
