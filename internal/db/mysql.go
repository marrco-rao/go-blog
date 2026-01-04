package db

import (
	"fmt"
	"go-blog/internal/config"
	"go-blog/internal/model"
	"go-blog/pkg/logger"
	"go.uber.org/zap"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMySQL() (*gorm.DB, error) {
	cfg := config.Cfg.MySQL

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password,
		cfg.Host, cfg.Port, cfg.DBName,
	)
	logger.Log.Info("mysql dns", zap.String("dsn", dsn))
	var err error
	// gorm 的logger使用默认的，输出到stdout。 待改进，结合封装的zap
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDB, _ := DB.DB()
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime * time.Minute)
	// 自动建表
	if err := autoMigrate(DB); err != nil {
		return nil, err
	}
	return DB, nil
}
func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Post{},
		&model.Comment{},
	)
}

func Close() {
	if DB != nil {
		sqlDB, _ := DB.DB()
		_ = sqlDB.Close()
	}
}
