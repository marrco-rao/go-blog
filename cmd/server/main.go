package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/internal/config"
	"go-blog/internal/db"
	"go-blog/internal/router"
	"go-blog/pkg/logger"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 模拟一些隐私数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	config.InitConfig() // 读取配置
	fmt.Println("config result: ", config.Cfg)

	logger.Init()
	defer logger.Sync()

	logger.Log.Info("config & logger initialized")

	_, err := db.InitMySQL()
	defer db.Close()
	if err != nil {
		panic(err)
	}

	logger.Log.Info("mysql initialized")

	r := router.SetupRouter()

	srv := &http.Server{
		Addr:    config.Cfg.App.Addr,
		Handler: r,
	}
	go func() {
		logger.Log.Info("server starting",
			zap.String("addr", config.Cfg.App.Addr),
		)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("listen error", zap.Error(err))
		}
	}()

	// 优雅关闭——参考学习，之前没用过
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Log.Info("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if er := srv.Shutdown(ctx); err != nil {
		logger.Log.Error("Server shutdown failed", zap.Error(er))
	}
	logger.Log.Info("Server exited")
}
