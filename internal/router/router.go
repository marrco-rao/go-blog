package router

import (
	"github.com/gin-gonic/gin"
	"go-blog/internal/config"
	"go-blog/internal/db"
	"go-blog/internal/handler"
	"go-blog/internal/middleware"
	"go-blog/internal/repository"
	"go-blog/internal/service"
)

// 路由注册
func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(
		middleware.Recovery(),
		middleware.RequestID(),
		middleware.LoggerWithContext(),
		middleware.TimeoutWithRoute(
			config.Cfg.Timeout.Default,
			config.Cfg.Timeout.Routes,
		),
		middleware.Logger(),
	)

	postRepo := repository.NewPostRepo(db.DB)
	postSvc := service.NewPostService(postRepo)
	postHandler := handler.NewPostHander(postSvc)
	api := r.Group("/api")
	{
		api.GET("/posts/hot", postHandler.HotList)

		//auth := api.Group("")
		//auth.Use(middleware.JWTAuth())
		//{
		//	auth.POST("/posts", postHandler.Create)
		//}
	}
	return r
}
